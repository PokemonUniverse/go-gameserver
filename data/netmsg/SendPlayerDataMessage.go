package netmsg

import (
	pnet "github.com/PokemonUniverse/nonamelib/network"
	"github.com/PokemonUniverse/nonamelib/position"
)

type SendPlayerData struct {
	UID       uint64
	Position  position.Position
	Direction uint16
	Money     int32
	Name      string
	//Outfit		pul.IOutfit
}

// GetHeader returns the header value of this message
func (m *SendPlayerData) GetHeader() uint8 {
	return pnet.HEADER_IDENTITY
}

// WritePacket write the needed object data to a Packet and returns it
func (m *SendPlayerData) WritePacket() pnet.IPacket {
	packet := pnet.NewPacketExt(m.GetHeader())
	packet.AddUint64(uint64(m.UID))
	packet.AddString(m.Name)
	packet.AddUint16(uint16(m.Position.X))
	packet.AddUint16(uint16(m.Position.Y))
	packet.AddUint16(m.Direction)
	packet.AddUint32(uint32(m.Money))

	/*packet.AddUint8(uint8(m.Outfit.GetOutfitStyle(pul.OUTFIT_UPPER)))
	packet.AddUint32(uint32(m.Outfit.GetOutfitColour(pul.OUTFIT_UPPER)))
	packet.AddUint8(uint8(m.Outfit.GetOutfitStyle(pul.OUTFIT_NEK)))
	packet.AddUint32(uint32(m.Outfit.GetOutfitColour(pul.OUTFIT_NEK)))
	packet.AddUint8(uint8(m.Outfit.GetOutfitStyle(pul.OUTFIT_HEAD)))
	packet.AddUint32(uint32(m.Outfit.GetOutfitColour(pul.OUTFIT_HEAD)))
	packet.AddUint8(uint8(m.Outfit.GetOutfitStyle(pul.OUTFIT_FEET)))
	packet.AddUint32(uint32(m.Outfit.GetOutfitColour(pul.OUTFIT_FEET)))
	packet.AddUint8(uint8(m.Outfit.GetOutfitStyle(pul.OUTFIT_LOWER)))
	packet.AddUint32(uint32(m.Outfit.GetOutfitColour(pul.OUTFIT_LOWER)))*/

	return packet
}
