package entities

import (
	"github.com/eaigner/hood"
)

type PlayerPokemonMove struct {
	PlayerPokemonMoveId hood.Id

	PokemonId int64
	MoveId    int64

	PpUsed int

	Created hood.Created
	Updated hood.Updated
}
