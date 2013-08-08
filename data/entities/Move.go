package entities

import (
	"github.com/eaigner/hood"
)

type Move struct {
	MoveId hood.Id

	Identifier string `sql:"size(255),notnull"`
}
