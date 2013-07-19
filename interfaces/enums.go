package interfaces

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
