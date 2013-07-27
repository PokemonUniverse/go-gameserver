package world

import ()

type TilePoint struct {
	tilePointLayers map[int]*TilePointLayer

	x     int
	y     int
	index int64
}

func NewTilePoint(_x, _y int) *TilePoint {
	return &TilePoint{tilePointLayers: make(map[int]*TilePointLayer),
		x:     _x,
		y:     _y,
		index: GenerateKey(_x, _y)}
}

func (t *TilePoint) GetX() int {
	return t.x
}

func (t *TilePoint) GetY() int {
	return t.y
}

func (t *TilePoint) GetIndex() int64 {
	return t.index
}

func (t *TilePoint) GetTilePointLayers() map[int]*TilePointLayer {
	return t.tilePointLayers
}

func (t *TilePoint) HasLayer(_layer int) bool {
	_, found := t.tilePointLayers[_layer]
	return found
}

// Gets a TilePointLayer from layer index
// Returns:	TilePointLayer object
//			Bool, true if the TilePointLayer was found, otherwise false
func (t *TilePoint) GetTilePointLayer(_layer int) (*TilePointLayer, bool) {
	tpl, found := t.tilePointLayers[_layer]
	return tpl, found
}

// Gets a TilePointLayer from layer index. If the TilePointLayer doesn't exists it will be created.
// Returns:	TilePointLayer object
//			Bool, true if TilePointLayer is new, otherwise false
func (t *TilePoint) getOrAddTilePointLayer(_layer int) (*TilePointLayer, bool) {
	tpl, found := t.GetTilePointLayer(_layer)
	if !found {
		tpl = NewTilePointLayer(_layer)
		t.tilePointLayers[_layer] = tpl
	}

	return tpl, !found
}
