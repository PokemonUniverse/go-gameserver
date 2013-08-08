package models

import (
	"github.com/PokemonUniverse/nonamelib/log"

	"gameserver/data/entities"
	"gameserver/data/pokemon"
)

type PlayerPokemonMove struct {
	Entity   *entities.PlayerPokemonMove
	BaseMove *pokemon.Move
}

func NewPlayerPokemonMove(_entity *entities.PlayerPokemonMove) *PlayerPokemonMove {
	m := &PlayerPokemonMove{}
	m.Entity = _entity

	if move, found := pokemon.GetMoveById(_entity.MoveId); found {
		m.BaseMove = move
	} else {
		log.Error("PlayerPokemonMove", "New", "Could not found move with id %d for PlayerPokemonMove %d", _entity.MoveId, _entity.PlayerPokemonMoveId)
		m = nil
	}

	return m
}

func (m *PlayerPokemonMove) GetPlayerPokemonMoveId() int64 {
	return int64(m.Entity.PlayerPokemonMoveId)
}
