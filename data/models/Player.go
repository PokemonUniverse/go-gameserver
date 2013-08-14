package models

import (
	"github.com/eaigner/hood"

	"github.com/PokemonUniverse/nonamelib/log"
	pnet "github.com/PokemonUniverse/nonamelib/network"
	"github.com/PokemonUniverse/nonamelib/position"

	"gameserver/data/entities"
	"gameserver/data/netmsg"
	"gameserver/game"
	"gameserver/interfaces"
	"gameserver/world"
)

type Player struct {
	Creature // base
	PlayerEntity *entities.Player
	
	dbConn *hood.Hood
	rxChan <-chan pnet.INetMessageReader
	txChan chan<- pnet.INetMessageWriter

	pokemon map[int64]*PlayerPokemon
}

func NewPlayer(_entity *entities.Player, _dbConn *hood.Hood) *Player {
	p := &Player{}
	p.Creature.init()
	p.PlayerEntity = _entity
	p.Creature.Name = _entity.Name
	p.dbConn = _dbConn

	p.pokemon = make(map[int64]*PlayerPokemon)

	return p
}

// Start ICreature
func (p *Player) GetCreatureType() interfaces.CreatureType {
	return interfaces.CREATURE_TYPE_PLAYER
}

func (p *Player) LoadCharacterData() bool {
	if p.loadPlayerPokemon() == false {
		return false
	}

	return true
}

func (p *Player) Walk(_from position.Position, _to position.Position, _teleported bool, _direction uint16) {
	if _to == position.ZP {
		// Position is reset, send position to client so it's synced
		p.SetPosition(_from)
		p.netSendCreatureMove(p, _from, _from, false)
	} else {
		// Call base function
		p.Creature.Walk(_from, _to, _teleported, _direction)

		// Send new tiles to client
		p.netSendMapData(_direction)
	}
}

func (p *Player) OnCreatureMove(_creature interfaces.ICreature, _from position.Position, _to position.Position, _teleport bool) {
	// Call base function
	p.Creature.OnCreatureMove(_creature, _from, _to, _teleport)

	// Do nothing is creature is me
	if p.GetUID() == _creature.GetUID() {
		return
	}

	canSeeFromPosition := p.CanSeePosition(_from)
	canSeeToPosition := p.CanSeePosition(_to)

	if canSeeFromPosition && canSeeToPosition {
		// We can see both from and to positions, meaning the creature is moving inside the viewport
		p.netSendCreatureMove(_creature, _from, _to, _teleport)
	} else if canSeeFromPosition && !canSeeToPosition {
		// We can see the from position but not the to positioin, meaning the creature is leaving our viewport
		// Send CreatureMove message to client before removing the creature from the visible creatures list
		p.netSendCreatureMove(_creature, _from, _to, _teleport)

		p.RemoveVisibleCreature(_creature)
		_creature.RemoveVisibleCreature(p)
	} else if !canSeeFromPosition && canSeeToPosition {
		// We can't see the from position but we can see the to position, meaning the creature is entering our viewport
		// First add the creature to our visible creature list before sending the CreatureMove message to the client 
		p.AddVisibleCreature(_creature)
		_creature.AddVisibleCreature(p)

		p.netSendCreatureMove(_creature, _from, _to, _teleport)
	} else {
		// Somehow it's impossible to see the creature, may this ever happen somehow
		p.RemoveVisibleCreature(_creature)
		_creature.RemoveVisibleCreature(p)
	}
}

func (p *Player) OnCreatureTurn(_creature interfaces.ICreature) {
	p.Creature.OnCreatureTurn(_creature)

	// Do nothing is creature is me
	if p.GetUID() == _creature.GetUID() {
		return
	}

	p.netSendCreatureTurn(_creature)
}

func (p *Player) AddVisibleCreature(_creature interfaces.ICreature) bool {
	ret := p.Creature.AddVisibleCreature(_creature)

	if ret {
		p.netSendCreatureAdd(_creature)
	}

	return ret
}

func (p *Player) RemoveVisibleCreature(_creature interfaces.ICreature) bool {
	ret := p.Creature.RemoveVisibleCreature(_creature)

	if ret {
		p.netSendCreatureRemove(_creature)
	}

	return ret
}

// Player data loading
func (p *Player) loadPlayerPokemon() bool {
	var result []*entities.PlayerPokemon
	if err := p.dbConn.Where("PlayerId", "=", p.GetPlayerId()).Find(result); err != nil {
		log.Error("Player", "laodPlayerPokemon", "Failed to load pokemon for player (%d). %s", p.GetPlayerId(), err.Error())
		return false
	}
	
	p.pokemon = make(map[int64]*PlayerPokemon)
	for _, playerPokemonEntity := range(result) {
		playerPokemon := NewPlayerPokemon(p.dbConn, playerPokemonEntity)
		p.pokemon[playerPokemon.GetPlayerPokemonId()] = playerPokemon
	}
	
	// TODO: Maybe add extra check if pokemonList size is zero. But only if we assign a pokemon to the player when creating a character
	
	return true
}

// Player specific methods

func (p *Player) GetPlayerId() int64 {
	return int64(p.PlayerEntity.PlayerId)
}

func (p *Player) GetMoney() int32 {
	return p.PlayerEntity.Money
}

// Networking - Receive

func (p *Player) SetNetworkChans(_rx <-chan pnet.INetMessageReader, _tx chan<- pnet.INetMessageWriter) {
	p.rxChan = _rx
	p.txChan = _tx

	go p.netReceiveMessages()
}

func (p *Player) netReceiveMessages() {
	for {
		netmessage := <-p.rxChan
		if netmessage == nil {
			break
		}

		switch netmessage.GetHeader() {
		case pnet.HEADER_LOGIN:
			p.netSendPlayerData()
		case pnet.HEADER_WALK:
			p.netHeaderWalk(netmessage.(*netmsg.WalkMessage))
		case pnet.HEADER_TURN:
			p.netHeaderTurn(netmessage.(*netmsg.TurnMessage))
		default:
			log.Warning("Player", "netReceiveMessages", "No handler for messages with header %d", netmessage.GetHeader())
		}
	}
}

func (p *Player) netHeaderWalk(_netmessage *netmsg.WalkMessage) {
	game.OnPlayerMove(p, _netmessage.Direction)
}

func (p *Player) netHeaderTurn(_netmessage *netmsg.TurnMessage) {
	game.OnCreatureTurn(p, _netmessage.Direction)
}

// Networking - Send

func (p *Player) netSendPlayerData() {
	playerData := &netmsg.SendPlayerData{}
	playerData.UID = p.GetUID()
	playerData.Name = p.GetName()
	playerData.Position = p.GetPosition()
	playerData.Direction = p.GetDirection()
	playerData.Money = p.GetMoney()

	p.txChan <- playerData

	// TODO: Send PkMn
	// ----

	// TODO: Send items
	// ----

	// TODO: Send friends list to client
	// ---

	// TODO: Send quests
	// ----

	// Send map
	p.netSendMapData(interfaces.DIR_NULL)

	// Ready to roll
	readyMessage := &netmsg.LoginMessage{}
	readyMessage.Status = netmsg.LOGINSTATUS_READY
	p.txChan <- readyMessage
}

func (p *Player) netSendCreatureMove(_creature interfaces.ICreature, _from position.Position, _to position.Position, _teleport bool) {
	msg := netmsg.NewWalkMessage(_creature)
	msg.From = _from
	msg.To = _to
	msg.Teleport = _teleport

	p.txChan <- msg
}

func (p *Player) netSendCreatureTurn(_creature interfaces.ICreature) {
	msg := netmsg.NewTurnMessage(_creature)
	msg.AddDirection(_creature.GetDirection())

	p.txChan <- msg
}

func (p *Player) netSendMapData(_direction uint16) {
	xMin := 1
	xMax := interfaces.CLIENT_VIEWPORT.X
	yMin := 1
	yMax := interfaces.CLIENT_VIEWPORT.Y

	if _direction != interfaces.DIR_NULL {
		switch _direction {
		case interfaces.DIR_NORTH:
			yMax = 1
		case interfaces.DIR_EAST:
			xMin = interfaces.CLIENT_VIEWPORT.X
		case interfaces.DIR_SOUTH:
			yMin = interfaces.CLIENT_VIEWPORT.Y
		case interfaces.DIR_WEST:
			xMax = 1
		}
	}

	// Top-left coordinates
	positionX := (p.GetPosition().X - interfaces.CLIENT_VIEWPORT_CENTER.X)
	positionY := (p.GetPosition().Y - interfaces.CLIENT_VIEWPORT_CENTER.Y)
	positionZ := p.GetPosition().Z
	mapId := p.GetPosition().MapId

	tilesMessage := netmsg.NewSendTilesMessage()
	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			if tp, found := world.GetTilePoint(mapId, positionX+x, positionY+y); found {
				if tp.HasLayer(positionZ) || tp.HasLayer(positionZ+1) || tp.HasLayer(positionZ-1) {
					tilesMessage.AddTile(tp)
				}
			}
		}

		if _direction == interfaces.DIR_NULL {
			p.txChan <- tilesMessage
		}
	}

	if _direction != interfaces.DIR_NULL {
		p.txChan <- tilesMessage
	}
}

func (p *Player) netSendCreatureAdd(_creature interfaces.ICreature) {
	msg := netmsg.NewCreatureAddMessage(_creature)
	p.txChan <- msg
}

func (p *Player) netSendCreatureRemove(_creature interfaces.ICreature) {
	msg := netmsg.NewCreatureRemoveMessage(_creature)
	p.txChan <- msg
}
