package entities

import (
	"github.com/eaigner/hood"
)

type Map struct {
	MapId hood.Id

	Name string `sql:"size(45),notnull"`

	Created hood.Created
	Updated hood.Updated
}
