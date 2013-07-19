package server

import (
	"fmt"
	"net/http"

	"code.google.com/p/go.net/websocket"
	"github.com/PokemonUniverse/nonamelib/log"
	pnet "github.com/PokemonUniverse/nonamelib/network"

	cmap "github.com/PokemonUniverse/nonamelib/container/concurrentmap"

	"gameserver/data"
	"gameserver/data/netmsg"
	"gameserver/game"
)

var gameserver *GameServer

type GameServer struct {
	port int // Port to listen on

	// Keep own list of connected playernames paired with internal UID
	// So we don't have extra locking on the main player list on login
	// Key: Idplayer (int) | Value: UID (uint64)
	connectedPlayers *cmap.ConcurrentMap
}

func init() {
	gameserver = newGameServer()
}

func newGameServer() *GameServer {
	return &GameServer{port: 6161,
		connectedPlayers: cmap.New()}
}

// Start listening for new connections
func Listen(_port int) {
	gameserver.port = _port

	log.Info("GameServer", "Listen", "Listening for connections on port: %d", _port)

	http.Handle("/puserver", websocket.Handler(clientConnection))
	err := http.ListenAndServe(fmt.Sprintf(":%d", _port), nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

// Entrance method for new websocket connections 
func clientConnection(clientsock *websocket.Conn) {
	packet := pnet.NewPacket()
	buffer := make([]uint8, pnet.PACKET_MAXSIZE)
	recv, err := clientsock.Read(buffer)
	if err == nil {
		copy(packet.Buffer[0:recv], buffer[0:recv])
		packet.GetHeader()
		parseFirstMessage(clientsock, packet)
	} else {
		if err.Error() != "EOF" {
			log.Warning("GameServer", "clientConnection", "Client connection error: %v", err.Error())
		}
	}
}

// Check login credentials
// Returns nothing. If incomming packet is wrong connection is closed, otherwise the login result will be send to the client
func parseFirstMessage(_conn *websocket.Conn, _packet *pnet.Packet) {
	// Read packet header
	header, err := _packet.ReadUint8()
	if err != nil || (err == nil && header != pnet.HEADER_LOGIN) {
		_conn.Close()
		return
	}

	// Parse packet, we can use the same packet for sending the return status
	firstMessage := &netmsg.LoginMessage{}
	if err := firstMessage.ReadPacket(_packet); err != nil {
		_conn.Close()
		return
	}

	// Create wrapper for websocket connection
	connection := NewConnectionWrapper(_conn)

	// TODO: Add pointer to Hood ORM
	player, err := data.PlayerHelper.GetPlayerUsingCredentials(nil, firstMessage.Username, firstMessage.Password)
	if err != nil {
		firstMessage.Status = netmsg.LOGINSTATUS_WRONGACCOUNT
		log.Debug("GameServer", "parseFirstMessage", "Invalid credentials for '%s'. Error: %s", firstMessage.Username, err.Error())
	} else {
		var playerLoaded bool = false
		// Check if player is already logged in
		if uid, found := gameserver.connectedPlayers.Get(player.GetPlayerId()); found {
			if _, f := game.GetPlayerByUID(uid.(uint64)); f {
				playerLoaded = true
			} else {
				// Remove player from connected player list
				gameserver.connectedPlayers.Remove(player.GetPlayerId())
			}
		}

		if playerLoaded {
			firstMessage.Status = netmsg.LOGINSTATUS_ALREADYLOGGEDIN
			log.Debug("GameServer", "parseFirstMessage", "Player '%s' is already logged in", firstMessage.Username)
		} else {
			// Load rest of player data into player object
			if !player.LoadCharacterData() {
				firstMessage.Status = netmsg.LOGINSTATUS_FAILPROFILELOAD
				log.Debug("GameServer", "parseFirstMessage", "Failed to load profile for '%s'", firstMessage.Username)
			} else {
				firstMessage.Status = netmsg.LOGINSTATUS_READY
				log.Debug("GameServer", "parseFirstMessage", "Login OK - Player: %s", player.GetName())

				// Assign connection to player, so Netmessages are handled through player object
				connection.AssignToPlayer(player)
				connection.txChan <- firstMessage

				// Add Player object to game
				game.AddCreature(player)

				return
			}
		}
	}

	connection.txChan <- firstMessage
	connection.Close()
}
