package interfaces

import (
	pnet "github.com/PokemonUniverse/nonamelib/network"
)

type IPlayer interface {
	ICreature

	// Generic
	GetPlayerId() int64

	// Networking
	SetNetworkChans(_rx <-chan pnet.INetMessageReader, _tx chan<- pnet.INetMessageWriter)
}
