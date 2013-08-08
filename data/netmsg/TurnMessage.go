package netmsg

import (
	pnet "github.com/PokemonUniverse/nonamelib/network"

	"gameserver/interfaces"
)

type TurnMessage struct {
	Creature  interfaces.ICreature
	Direction uint16
}

func NewTurnMessage(_creature interfaces.ICreature) *TurnMessage {
	return &TurnMessage{Creature: _creature}
}

// GetHeader returns the header value of this message
func (m *TurnMessage) GetHeader() uint8 {
	return pnet.HEADER_TURN
}

func (m *TurnMessage) AddDirection(_dir uint16) {
	m.Direction = _dir
}

func (m *TurnMessage) ReadPacket(_packet pnet.IPacket) error {
	direction, err := _packet.ReadUint16()
	if err != nil {
		return err
	}
	m.Direction = direction

	return nil
}

// WritePacket write the needed object data to a Packet and returns it
func (m *TurnMessage) WritePacket() pnet.IPacket {
	packet := pnet.NewPacketExt(m.GetHeader())
	packet.AddUint64(m.Creature.GetUID())
	packet.AddUint16(m.Direction)

	return packet
}
