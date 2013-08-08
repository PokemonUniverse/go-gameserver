package models

import (
	"github.com/PokemonUniverse/nonamelib/log"
	"github.com/eaigner/hood"

	"gameserver/data/entities"
	"gameserver/data/pokemon"
)

type PlayerPokemon struct {
	Entity      *entities.PlayerPokemon
	BasePokemon *pokemon.Pokemon

	moves map[int64]*PlayerPokemonMove
}

func NewPlayerPokemon(_hood *hood.Hood, _entity *entities.PlayerPokemon) *PlayerPokemon {
	p := &PlayerPokemon{}
	p.Entity = _entity

	if basePokemon, found := pokemon.GetPokemonById(_entity.PokemonId); found {
		p.BasePokemon = basePokemon
	} else {
		log.Warning("PlayerPokemon", "NewPlayerPokemon", "Unable to find Pokemon with id %d, for PlayerPokemon entity: %d", _entity.PokemonId, _entity.PlayerPokemonId)
		p = nil
	}

	if p != nil {
		p.loadMoves(_hood)
	}

	return p
}

func (p *PlayerPokemon) GetPlayerPokemonId() int64 {
	return int64(p.Entity.PlayerPokemonId)
}

func (p *PlayerPokemon) loadMoves(_hood *hood.Hood) {
	var pokemonMoves []entities.PlayerPokemonMove

	if err := _hood.Where("pokemon_id", "=", p.GetPlayerPokemonId()).Find(&pokemonMoves); err != nil {
		log.Error("PlayerPokemon", "loadMoves", "Failed to load PlayerPokemonMoves for PlayerPokemon: %d. Error: %s", p.GetPlayerPokemonId(), err.Error())
	} else {
		for _, move := range pokemonMoves {
			if playerPokemonMove := NewPlayerPokemonMove(&move); playerPokemonMove != nil {
				p.moves[playerPokemonMove.GetPlayerPokemonMoveId()] = playerPokemonMove
			}
		}
	}
}
