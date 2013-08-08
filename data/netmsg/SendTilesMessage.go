package netmsg

import (
	pnet "github.com/PokemonUniverse/nonamelib/network"

	"gameserver/world"
)

type SendTilesMessage struct {
	tiles map[int64]*world.TilePoint
}

func NewSendTilesMessage() *SendTilesMessage {
	return &SendTilesMessage{tiles: make(map[int64]*world.TilePoint)}
}

// GetHeader returns the header value of this message
func (m *SendTilesMessage) GetHeader() uint8 {
	return pnet.HEADER_TILES
}

func (m *SendTilesMessage) AddTile(_tile *world.TilePoint) {
	m.tiles[_tile.GetIndex()] = _tile
}

// WritePacket write the needed object data to a Packet and returns it
func (m *SendTilesMessage) WritePacket() pnet.IPacket {
	totalTiles := uint16(len(m.tiles))
	packet := pnet.NewPacketExt(m.GetHeader())
	packet.AddUint16(totalTiles)

	if totalTiles == 0 {
		return packet
	}

	for _, tile := range m.tiles {
		if tile == nil {
			continue
		}

		packet.AddUint16(uint16(tile.GetX()))
		packet.AddUint16(uint16(tile.GetY()))

		// Loop TilePointLayers
		tilePointLayers := tile.GetTilePointLayers()
		packet.AddUint16(uint16(len(tilePointLayers)))
		for level, tilePointLayer := range tilePointLayers {
			if tilePointLayer == nil {
				continue
			}

			packet.AddUint8(uint8(level))
			packet.AddUint16(uint16(tilePointLayer.GetBlocking()))

			// Loop TilePointTileLayers
			tilePointTileLayers := tilePointLayer.GetTilePointTileLayers()
			packet.AddUint16(uint16(len(tilePointTileLayers)))
			for layer, tilePointTileLayer := range tilePointTileLayers {
				if tilePointTileLayer == nil {
					continue
				}

				packet.AddUint8(uint8(layer))
				packet.AddString(tilePointTileLayer.GetTileId())
			}
		}
	}

	m.tiles = make(map[int64]*world.TilePoint)

	return packet
}
