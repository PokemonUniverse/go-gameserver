package entities

import (
	"github.com/eaigner/hood"
)

type Configuration struct {
	ConfigurationId hood.Id

	Section string `sql:"size(45),notnull"`
	Name    string `sql:"size(45),notnull"`
	Value   string `sql:"size(255),notnull"`

	Created hood.Created
	Updated hood.Updated
}
