package data

import (
	"fmt"

	"github.com/eaigner/hood"

	"gameserver/data/entities"
	"gameserver/data/models"
	"gameserver/interfaces"
)

var PlayerHelper *playerHelper

type playerHelper struct{}

func init() {
	PlayerHelper = &playerHelper{}
}

func (ph *playerHelper) ConvertPlayerEntityToModel(_entity *entities.Player) (*models.Player, error) {
	dbConn := CreateDatabaseConnection()
	
	if dbConn == nil {
		return nil, fmt.Errorf("Could not create database connection for player '%s'", _entity.Name)
	}
	
	p := models.NewPlayer(_entity, dbConn)
	p.Name = _entity.Name

	return p, nil
}

func (ph *playerHelper) GetPlayerUsingCredentials(_orm *hood.Hood, _username, _password string) (interfaces.IPlayer, error) {
	var result *entities.Player
	err := _orm.Where("Username", "=", _username).Find(result)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, fmt.Errorf("Player '%s' not found", _username)
	}

	// Do some funky stuff to compare passwords
	if result.Password != _password {
		return nil, fmt.Errorf("Supplied credentials for username '%s' are not valid", _username)
	}

	playerModel, err := ph.ConvertPlayerEntityToModel(result)
	if err != nil {
		return nil, err
	}

	return playerModel, nil
}