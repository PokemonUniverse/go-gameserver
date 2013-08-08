package entities

import (
	"github.com/eaigner/hood"
)

type PlayerPokemon struct {
	PlayerPokemonId hood.Id

	PlayerId  int64
	PokemonId int64
	InParty   bool
	PartySlot int
	HasSeen   bool
	HasCaught bool
	IsTraded  bool

	Experience  int64
	DamageTaken int

	BaseStatHp        int
	BaseStatAttack    int
	BaseStatDefense   int
	BaseStatSpAttack  int
	BaseStatSpDefense int
	BaseStatSpeed     int

	Created hood.Created
	Updated hood.Updated
}
