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

func (ph *playerHelper) ConvertPlayerEntityToModel(_entity entities.Player) *models.Player {
	p := &models.Player{}
	p.Name = _entity.Name

	return p
}

func (ph *playerHelper) GetPlayerUsingCredentials(_orm *hood.Hood, _username, _password string) (interfaces.IPlayer, error) {
	var result []entities.Player
	err := _orm.Where("Username", "=", _username).Find(&result)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("Player '%s' not found", _username)
	}
	if len(result) > 1 {
		return nil, fmt.Errorf("Multiple players found with username '%s'", _username)
	}

	playerEntity := result[0]

	// Do some funky stuff to compare passwords
	if playerEntity.Password != _password {
		return nil, fmt.Errorf("Supplied credentials for username '%s' are not valid", _username)
	}

	playerModel := ph.ConvertPlayerEntityToModel(playerEntity)

	return playerModel, nil
}
