package server

import (
	"code.google.com/p/go.net/websocket"
	"github.com/PokemonUniverse/nonamelib/log"
	pnet "github.com/PokemonUniverse/nonamelib/network"

	"gameserver/data/netmsg"
	"gameserver/interfaces"
)

type ConnectionWrapper struct {
	socket *websocket.Conn

	txChan chan pnet.INetMessageWriter
	rxChan chan pnet.INetMessageReader

	player interfaces.IPlayer
}

func NewConnectionWrapper(_socket *websocket.Conn) *ConnectionWrapper {
	connectionWrapper := &ConnectionWrapper{socket: _socket,
		txChan: make(chan pnet.INetMessageWriter),
		rxChan: make(chan pnet.INetMessageReader)}

	go connectionWrapper.ReceivePoller()
	go connectionWrapper.SendPoller()

	return connectionWrapper
}

func (cw *ConnectionWrapper) Close() {
	// Close channels
	close(cw.txChan)
	close(cw.rxChan)

	// Close socket
	cw.socket.Close()

	cw.player = nil
}

func (cw *ConnectionWrapper) AssignToPlayer(_player interfaces.IPlayer) {
	if _player == nil {
		panic("ConnectionWrapper - Player interface can not be nil")
	}

	cw.player = _player
	_player.SetNetworkChans(cw.rxChan, cw.txChan)
}

// Read data from socket stream
func (cw *ConnectionWrapper) ReceivePoller() {
	for {
		packet := pnet.NewPacket()
		var buffer []uint8
		err := websocket.Message.Receive(cw.socket, &buffer)
		if err == nil {
			copy(packet.Buffer[0:len(buffer)], buffer[0:len(buffer)])
			packet.GetHeader()

			cw.processPacket(packet)
		} else {
			println(err.Error())
			break
		}
	}
}

// Write data to socket stream
func (cw *ConnectionWrapper) SendPoller() {
	for {
		// Read messages from transmit channel
		netmessage := <-cw.txChan

		if netmessage == nil {
			log.Debug("ConnectionWrapper", "SendPoller", "Netmessage == nil, breaking loop")
			break
		}

		// Convert netmessage to packet
		packet := netmessage.WritePacket()
		packet.SetHeader()

		// Create byte buffer
		buffer := packet.GetBuffer()
		data := buffer[0:packet.GetMsgSize()]

		// Send bytes off to the internetz
		websocket.Message.Send(cw.socket, data)
	}
}

func (cw *ConnectionWrapper) processPacket(_packet pnet.IPacket) {
	header, err := _packet.ReadUint8()
	if err != nil {
		return
	}

	var netmessage pnet.INetMessageReader
	if header == pnet.HEADER_WALK {
		netmessage = netmsg.NewWalkMessage(nil)
		netmessage.ReadPacket(_packet)
	} else {
		log.Warning("ConnectionWrapper", "processPacket", "Received packet with unknown header: %d", header)
	}

	// Push netmessage on rxChan
	if netmessage != nil {
		cw.rxChan <- netmessage
	}
}
