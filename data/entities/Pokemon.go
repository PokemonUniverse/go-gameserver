package entities

import (
	"github.com/eaigner/hood"
)

type Pokemon struct {
	PokemonId hood.Id

	DexId           int
	Identifier      string
	Height          string
	Weight          string
	NextEvolutionId int64
	TypeId          int
	CaptureRate     int
	GenderRate      int
	BaseHappiness   int
	IsBaby          bool
	HatchRate       int
	FormsSwitchable bool

	Created hood.Created
	Updated hood.Updated
}
