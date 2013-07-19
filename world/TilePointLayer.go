package world

import (
	"gameserver/interfaces"
)

type TilePointLayer struct {
	tilePointTileLayers map[int]*TilePointTileLayer

	// List of creatures whom are active on this TilePointLayer
	creatures map[uint64]interfaces.ICreature

	blocking interfaces.TileBlocking
	layer    int
}

func NewTilePointLayer(_layer int) *TilePointLayer {
	return &TilePointLayer{tilePointTileLayers: make(map[int]*TilePointTileLayer),
		creatures: make(map[uint64]interfaces.ICreature),
		layer:     _layer,
		blocking:  interfaces.TILEBLOCK_BLOCK}
}

func (t *TilePointLayer) SetBlocking(_blocking interfaces.TileBlocking) {
	t.blocking = _blocking
}

// Gets a TilePointTileLayer using layer index.
// Returns:	TilePointTileLayer object
//			Bool, true if the TilePointTileLayer was found, otherwise false
func (t *TilePointLayer) GetTilePointTileLayer(_layer int) (*TilePointTileLayer, bool) {
	tptl, found := t.tilePointTileLayers[_layer]
	return tptl, found
}

// Gets a TilePointTileLayer using layer index. If the TilePointTileLayer doesn't exists it will be created.
// Returns:	TilePointTileLayer object
//			Bool, true if TilePointTileLayer is new, otherwise false
func (t *TilePointLayer) getOrAddTilePointTileLayer(_layer int, _tileId string) (*TilePointTileLayer, bool) {
	tptl, found := t.GetTilePointTileLayer(_layer)
	if !found {
		tptl = NewTilePointTileLayer(_layer, _tileId)
		t.tilePointTileLayers[_layer] = tptl
	}

	return tptl, !found
}
