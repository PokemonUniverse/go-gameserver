package data

import (
	"fmt"

	"github.com/eaigner/hood"
	_ "github.com/ziutek/mymysql/godrv"

	"github.com/PokemonUniverse/nonamelib/configuration"
	"github.com/PokemonUniverse/nonamelib/log"

	"gameserver/config"
)

func CreateDatabaseConnection() *hood.Hood {
	username, _ := configuration.GetString(config.CONFIG_DB_USER)
	password, _ := configuration.GetString(config.CONFIG_DB_PASSWORD)
	scheme, _ := configuration.GetString(config.CONFIG_DB_SCHEME)
	connectionString := fmt.Sprintf("%v/%v/%v", scheme, username, password)

	hd, err := hood.Open("mymysql", connectionString)
	if err != nil {
		log.Error("DatabaseHelper", "CreateDatabaseConnection", "Unable to connect to database scheme '%s' with supplied credentials '%s'.\nError: %s", scheme, username, err.Error())
		return nil
	}

	return hd
}
