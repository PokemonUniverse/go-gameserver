package entities

import (
	"github.com/eaigner/hood"
)

type Player struct {
	PlayerId hood.Id

	Name     string `sql:"size(15),notnull"`
	Password string `sql:"size(255),notnull"`
	Salt     string `sql:"size(255),notnull"`

	Money int32

	Created hood.Created
	Updated hood.Updated
}
