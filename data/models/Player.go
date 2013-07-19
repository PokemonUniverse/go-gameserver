package models

import (
	pnet "github.com/PokemonUniverse/nonamelib/network"

	"gameserver/data/netmsg"
	"gameserver/game"
	"gameserver/interfaces"
)

type Player struct {
	Name string

	rxChan <-chan pnet.INetMessageReader
	txChan chan<- pnet.INetMessageWriter
}

// Start ICreature
func (p *Player) GetCreatureType() interfaces.CreatureType {
	return interfaces.CREATURE_TYPE_PLAYER
}

func (p *Player) GetUID() uint64 {
	return 0
}

func (p *Player) GetName() string {
	return ""
}

func (p *Player) LoadCharacterData() bool {
	return true
}

func (p *Player) Walk() {
}

// End ICreature

func (p *Player) GetPlayerId() int {
	return 0
}

func (p *Player) DoSomething() {
	game.AddCreature(nil)
}

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
		case pnet.HEADER_WALK:
			p.netHeaderWalk(netmessage.(*netmsg.WalkMessage))
		}
	}
}

func (p *Player) netHeaderWalk(_netmessage *netmsg.WalkMessage) {
	game.OnPlayerMove(p, _netmessage.Direction)
}
