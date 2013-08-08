package interfaces

import (
	"time"

	"github.com/PokemonUniverse/nonamelib/position"
)

// Constants
const NANOSECONDS_TO_MILLISECONDS float64 = 0.000001
const MILLISECONDS_TO_SECOND float64 = 0.001
const MILLISECONDS_TO_MINUTES float64 = (MILLISECONDS_TO_SECOND / 60)
const MINUTES_TO_SECOND = 60
const SECONDS_TO_MILLISECOND = 1000
const MINUTES_TO_MILLISECONDS = 60000

func PUSYS_TIME() int64 {
	timeNano := float64(time.Now().UnixNano())
	return int64(timeNano * NANOSECONDS_TO_MILLISECONDS)
}

var (
	CLIENT_VIEWPORT        position.Position = position.Position{36, 28, 1, 0}
	CLIENT_VIEWPORT_CENTER position.Position = position.Position{17, 14, 1, 0}

	VIEWABLE_TILES int = (CLIENT_VIEWPORT.X * CLIENT_VIEWPORT.Y)
)

const (
	CONFIGURATION_SERVER      string = ""
	CONFIGURATION_SERVER_PORT string = ""
)

type CreatureType uint8

const (
	CREATURE_TYPE_UNKNOWN CreatureType = 0
	CREATURE_TYPE_PLAYER  CreatureType = 1
)

type TileBlocking int

const (
	TILEBLOCK_BLOCK       TileBlocking = 1
	TILEBLOCK_WALK                     = 2
	TILEBLOCK_SURF                     = 3
	TILEBLOCK_TOP                      = 4
	TILEBLOCK_BOTTOM                   = 5
	TILEBLOCK_RIGHT                    = 6
	TILEBLOCK_LEFT                     = 7
	TILEBLOCK_TOPRIGHT                 = 8
	TILEBLOCK_BOTTOMRIGHT              = 9
	TILEBLOCK_BOTTOMLEFT               = 10
	TILEBLOCK_TOPLEFT                  = 11
)

func GetTileBlockingFromInt(_blocking int) TileBlocking {
	ret := TILEBLOCK_BLOCK
	switch _blocking {
	case 1:
		ret = TILEBLOCK_BLOCK
	case 2:
		ret = TILEBLOCK_WALK
	case 3:
		ret = TILEBLOCK_SURF
	case 4:
		ret = TILEBLOCK_TOP
	case 5:
		ret = TILEBLOCK_BOTTOM
	case 6:
		ret = TILEBLOCK_RIGHT
	case 7:
		ret = TILEBLOCK_LEFT
	case 8:
		ret = TILEBLOCK_TOPRIGHT
	case 9:
		ret = TILEBLOCK_BOTTOMRIGHT
	case 10:
		ret = TILEBLOCK_BOTTOMLEFT
	case 11:
		ret = TILEBLOCK_TOPLEFT
	}
	return ret
}

const (
	DIR_NULL  = 0
	DIR_SOUTH = 1
	DIR_WEST  = 2
	DIR_NORTH = 3
	DIR_EAST  = 4
)
