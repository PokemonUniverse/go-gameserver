package entities

import (
	"github.com/eaigner/hood"
)

type PokemonMove struct {
	PokemonMoveId hood.Id

	PokemonId int64
	MoveId    int64
	Level     int

	Created hood.Created
	Updated hood.Updated
}
