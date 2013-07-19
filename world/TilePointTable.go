package world

import ()

type TilePointTable struct {
	mapName    string
	tilePoints map[int64]*TilePoint
}

func NewTilePointTable(_mapName string) *TilePointTable {
	return &TilePointTable{mapName: _mapName,
		tilePoints: make(map[int64]*TilePoint)}
}

// Gets a TilePoint from X, Y coordinates.
// Returns:	TilePoint object
//			Bool, true if the tilepoint was found, otherwise false
func (t *TilePointTable) GetTilePoint(_x, _y int) (*TilePoint, bool) {
	index := GenerateKey(_x, _y)
	tp, found := t.tilePoints[index]

	return tp, found
}

// Gets a TilePoint from X, Y coordinates. If the TilePoint doesn't exists it will be created.
// Returns:	TilePoint object
//			Bool, true if TilePoint is new, otherwise false
func (t *TilePointTable) getOrAddTilePoint(_x, _y int) (*TilePoint, bool) {
	tilepoint, found := t.GetTilePoint(_x, _y)
	if !found {
		tilepoint = NewTilePoint(_x, _y)
		t.tilePoints[tilepoint.index] = tilepoint
	}

	return tilepoint, !found
}
