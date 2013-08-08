package pokemon

import (
	"fmt"

	"github.com/PokemonUniverse/nonamelib/log"
	"github.com/eaigner/hood"

	"gameserver/data/entities"
)

type Pokemon struct {
	PokemonEntity *entities.Pokemon

	moves map[int]*Move // Key = level
}

func NewPokemon(_entity *entities.Pokemon) *Pokemon {
	p := &Pokemon{PokemonEntity: _entity,
		moves: make(map[int]*Move)}
	return p
}

func (p *Pokemon) LinkMoves(_hood *hood.Hood) error {
	var pokemonMoves []entities.PokemonMove
	if err := _hood.Where("pokemon_id", "=", p.GetPokemonId()).Find(&pokemonMoves); err != nil {
		return err
	}

	if len(pokemonMoves) == 0 {
		return fmt.Errorf("No moves found for pokemon %d", p.GetPokemonId())
	}

	for _, pokeMove := range pokemonMoves {
		if move, found := GetMoveById(pokeMove.MoveId); found {
			p.moves[pokeMove.Level] = move
		} else {
			log.Error("Pokemon", "LinkMoves", "Move with id %d could not be found.", pokeMove.MoveId)
		}
	}

	return nil
}

func (p *Pokemon) GetPokemonId() int64 {
	return int64(p.PokemonEntity.PokemonId)
}
