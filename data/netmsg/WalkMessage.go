package netmsg

import (
	pnet "github.com/PokemonUniverse/nonamelib/network"
	pos "github.com/PokemonUniverse/nonamelib/position"

	"gameserver/interfaces"
)

type WalkMessage struct {
	// Receive
	Creature  interfaces.ICreature
	Direction uint16
	SendMap   bool

	// Send
	From pos.Position
	To   pos.Position
	Teleport bool
}

func NewWalkMessage(_creature interfaces.ICreature) *WalkMessage {
	message := &WalkMessage{}
	message.Creature = _creature

	return message
}

// GetHeader returns the header value of this message
func (m *WalkMessage) GetHeader() uint8 {
	return pnet.HEADER_WALK
}

// ReadPacket reads all data from a packet and puts it in the object
func (m *WalkMessage) ReadPacket(_packet pnet.IPacket) error {
	direction, err := _packet.ReadUint16()
	if err != nil {
		return err
	}
	m.Direction = direction

	sendMap, err := _packet.ReadUint16()
	if err != nil {
		return err
	}

	if sendMap == 1 {
		m.SendMap = true
	}

	return nil
}

func (m *WalkMessage) AddPositions(_from pos.Position, _to pos.Position) {
	m.From = _from
	m.To = _to
}

// WritePacket write the needed object data to a Packet and returns it
func (m *WalkMessage) WritePacket() pnet.IPacket {
	packet := pnet.NewPacketExt(m.GetHeader())
	packet.AddUint64(m.Creature.GetUID())
	packet.AddUint16(uint16(m.From.X))
	packet.AddUint16(uint16(m.From.Y))
	packet.AddUint16(uint16(m.To.X))
	packet.AddUint16(uint16(m.To.Y))

	return packet
}
