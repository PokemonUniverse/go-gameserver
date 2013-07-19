package interfaces

type ICreature interface {
	GetCreatureType() CreatureType
	GetUID() uint64
	GetName() string

	LoadCharacterData() bool

	Walk()
}
