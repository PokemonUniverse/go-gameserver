package interfaces

import (
	"github.com/PokemonUniverse/nonamelib/position"
)

type CreatureMap map[uint64]ICreature

type ICreature interface {
	GetCreatureType() CreatureType
	GetUID() uint64
	GetName() string
	
	GetMapId() int
	GetPosition() position.Position
	SetPosition(_position position.Position)

	LoadCharacterData() bool

	CanMove() bool
	Walk(_from position.Position, _to position.Position, _teleported bool, _direction uint16)
	
	// Methods for all moving creatures
	OnCreatureMove(_creature ICreature, _from position.Position, _to position.Position, _teleport bool)
	OnCreatureTurn(_creature ICreature)
	OnCreatureAppear(_creature ICreature, _isLogin bool)
	OnCreatureDisappear(_creature ICreature, _isLogout bool)
	
	// Methods for all creatures who need to see other creatures	
	AddVisibleCreature(_creature ICreature) bool
	RemoveVisibleCreature(_creature ICreature) bool
	KnowsVisibleCreature(_creature ICreature) bool
	GetVisibleCreatures() CreatureMap
}
