package netmsg

import (
	pnet "github.com/PokemonUniverse/nonamelib/network"
	
	"gameserver/interfaces"
)

type CreatureAddMessage struct {
	Creature interfaces.ICreature
}

func NewCreatureAddMessage(_creature interfaces.ICreature) *CreatureAddMessage {
	return &CreatureAddMessage { Creature: _creature }
}

// GetHeader returns the header value of this message
func (m *CreatureAddMessage) GetHeader() uint8 {
	return pnet.HEADER_ADDCREATURE
}

// WritePacket write the needed object data to a Packet and returns it
func (m *CreatureAddMessage) WritePacket() pnet.IPacket {
	packet := pnet.NewPacketExt(m.GetHeader())
	packet.AddUint64(m.Creature.GetUID())
	packet.AddString(m.Creature.GetName())
	packet.AddUint16(uint16(m.Creature.GetPosition().X))
	packet.AddUint16(uint16(m.Creature.GetPosition().Y))
	packet.AddUint16(m.Creature.GetDirection())
	
	// Outfit
	/*packet.AddUint8(uint8(m.Creature.GetOutfit().GetOutfitStyle(pul.OUTFIT_UPPER)))
	packet.AddUint32(uint32(m.Creature.GetOutfit().GetOutfitColour(pul.OUTFIT_UPPER)))
	packet.AddUint8(uint8(m.Creature.GetOutfit().GetOutfitStyle(pul.OUTFIT_NEK)))
	packet.AddUint32(uint32(m.Creature.GetOutfit().GetOutfitColour(pul.OUTFIT_NEK)))
	packet.AddUint8(uint8(m.Creature.GetOutfit().GetOutfitStyle(pul.OUTFIT_HEAD)))
	packet.AddUint32(uint32(m.Creature.GetOutfit().GetOutfitColour(pul.OUTFIT_HEAD)))
	packet.AddUint8(uint8(m.Creature.GetOutfit().GetOutfitStyle(pul.OUTFIT_UPPER)))
	packet.AddUint32(uint32(m.Creature.GetOutfit().GetOutfitColour(pul.OUTFIT_UPPER)))
	packet.AddUint8(uint8(m.Creature.GetOutfit().GetOutfitStyle(pul.OUTFIT_FEET)))
	packet.AddUint32(uint32(m.Creature.GetOutfit().GetOutfitColour(pul.OUTFIT_FEET)))
	packet.AddUint8(uint8(m.Creature.GetOutfit().GetOutfitStyle(pul.OUTFIT_LOWER)))
	packet.AddUint32(uint32(m.Creature.GetOutfit().GetOutfitColour(pul.OUTFIT_LOWER)))*/
	
	return packet
}