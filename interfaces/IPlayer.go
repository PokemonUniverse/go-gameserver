package interfaces

import (
	pnet "github.com/PokemonUniverse/nonamelib/network"
)

type IPlayer interface {

	// ICreature
	GetCreatureType() CreatureType
	GetName() string
	GetUID() uint64

	LoadCharacterData() bool

	Walk()

	// Generic
	GetPlayerId() int

	// Networking
	SetNetworkChans(_rx <-chan pnet.INetMessageReader, _tx chan<- pnet.INetMessageWriter)

	// Misc
	DoSomething()
}
