package world

import (
	"fmt"
	"time"

	"github.com/PokemonUniverse/nonamelib/log"
	"github.com/eaigner/hood"

	"gameserver/data/entities"
	"gameserver/interfaces"
)

/*
Map
- MapId
- Name

TilePoint
- TilePointId
- MapId
- X
- Y
- Index

TilePointLayer
- TilePointLayerId
- TilePointId
- Layer (Level)
- Blocking

TilePointTileLayer
- TilePointTileLayerId
- TilePointLayerId
- Layer
- TileId

*/

type TilePointRow struct {
	X        int
	Y        int
	Level    int
	Blocking int
	Layer    int
	TileId   string
}

var (
	loaded   bool
	worldmap map[int]*TilePointTable

	processExitChan      chan bool
	numOfProcessRoutines int
)

func init() {
	loaded = false
	worldmap = make(map[int]*TilePointTable)
	processExitChan = make(chan bool)
	numOfProcessRoutines = 0
}

// Generate a int64 key from X/Y-Coordinates
// Mr_Dark: I know there are some bits left, but that's because in the past we also had a Z-coordinate
func GenerateKey(_x, _y int) int64 {
	var x64 int64
	if _x < 0 {
		x64 = (int64(1) << 50) | ((^(int64(_x) - 1)) << 34)
	} else {
		x64 = (int64(_x) << 34)
	}

	var y64 int64
	if _y < 0 {
		y64 = (int64(1) << 33) | ((^(int64(_y) - 1)) << 17)
	} else {
		y64 = (int64(_y) << 17)
	}

	var index int64 = int64(x64 | y64)

	return index
}

func LoadWorldmap(_hood *hood.Hood) {
	if loaded {
		panic("Worldmap has already been loaded!")
	}

	// Load maps
	var maps []entities.Map
	if err := _hood.Find(&maps); err != nil {
		panic(err)
	}

	if len(maps) > 0 {
		for _, mapEntity := range maps {
			mapId := int(mapEntity.MapId)
			worldmap[mapId] = NewTilePointTable(mapEntity.Name)
		}

		numOfProcessRoutines = len(maps)
		go internalLoadWorldmap(_hood)

		waitForLoadComplete()
	} else {
		log.Warning("World", "LoadWorldmap", "No maps found in the database")
	}

	loaded = true
}

func GetTilePoint(_mapid, _x, _y int) (*TilePoint, bool) {
	if mapTable, found := worldmap[_mapid]; found {
		return mapTable.GetTilePoint(_x, _y)
	}

	return nil, false
}

// Waits for all spawned process routines to finish
// This can take a while if there are alot of tiles
func waitForLoadComplete() {
	completedRoutines := 0
	count := 0

	spinner := []string{"|", "/", "-", "\\", "|", "/", "-", "\\"}

	for {
		select {
		case <-processExitChan:
			completedRoutines++
			if completedRoutines == numOfProcessRoutines {
				log.Info("World", "LoadWorldmap", "Loading complete")
				break
			}
		case <-time.After(time.Millisecond * 200):
			fmt.Printf("\rLoading worldmap %s\r", spinner[count])
			count++
			if count > 7 {
				count = 0
			}
		}
	}
}

func internalLoadWorldmap(_hood *hood.Hood) {
	// Load tiles for each map
	for mapId, tilePointTable := range worldmap {
		go loadMap(_hood, mapId, tilePointTable)
	}
}

func loadMap(_hood *hood.Hood, _mapId int, _tilePointTable *TilePointTable) {
	// Create select query to get all tilepoints and layers as one result
	_hood.Select("tilepoint", "tilepoint.x", "tilepoint.y", "tilepointlayer.level", "tilepointlayer.blocking", "tilepointtilelayer.layer", "tilepointtilelayer.tileId")
	_hood.Join(hood.LeftJoin, "tilepointlayer", "tilepointlayer.tilePointId", "tilepoint.tilePointId")
	_hood.Join(hood.LeftJoin, "tilepointtilelayer", "tilepointtilelayer.tilePointLayerId", "tilepointlayer.tilePointLayerId")

	var tilePointRows []TilePointRow
	if err := _hood.Find(&tilePointRows); err != nil {
		log.Error("World", "loadMap", "Error fetching tilepoints from database:\n%s", err.Error())
	} else {
		for _, row := range tilePointRows {
			processTilePoints(row, _tilePointTable)
		}
	}

	processExitChan <- true
}

func processTilePoints(_row TilePointRow, _tilePointTable *TilePointTable) {
	tilePoint, _ := _tilePointTable.getOrAddTilePoint(_row.X, _row.Y)
	tilePointLayer, tplIsNew := tilePoint.getOrAddTilePointLayer(_row.Level)
	if tplIsNew {
		tilePointLayer.SetBlocking(interfaces.GetTileBlockingFromInt(_row.Blocking))
	}

	tilePointLayer.getOrAddTilePointTileLayer(_row.Layer, _row.TileId)
}
