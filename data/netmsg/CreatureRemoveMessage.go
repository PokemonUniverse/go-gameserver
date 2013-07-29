package netmsg

import (
	pnet "github.com/PokemonUniverse/nonamelib/network"
	
	"gameserver/interfaces"
)

type CreatureRemoveMessage struct {
	Creature interfaces.ICreature
}

func NewCreatureRemoveMessage(_creature interfaces.ICreature) *CreatureRemoveMessage {
	return &CreatureRemoveMessage { Creature: _creature }
}

// GetHeader returns the header value of this message
func (m *CreatureRemoveMessage) GetHeader() uint8 {
	return pnet.HEADER_REMOVECREATURE
}

// WritePacket write the needed object data to a Packet and returns it
func (m *CreatureRemoveMessage) WritePacket() pnet.IPacket {
	packet := pnet.NewPacketExt(m.GetHeader())
	packet.AddUint64(m.Creature.GetUID())
	
	return packet
}