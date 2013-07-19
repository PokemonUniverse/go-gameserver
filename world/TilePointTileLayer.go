package world

import ()

type TilePointTileLayer struct {
	layer  int
	tileId string
}

func NewTilePointTileLayer(_layer int, _tileId string) *TilePointTileLayer {
	return &TilePointTileLayer{layer: _layer,
		tileId: _tileId}
}

func (t *TilePointTileLayer) GetLayer() int {
	return t.layer
}

func (t *TilePointTileLayer) GetTileId() string {
	return t.tileId
}
