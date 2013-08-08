package data

import (
	"gameserver/data/entities"
	"gameserver/data/models"
	"gameserver/interfaces"
)

func CreatePlayer(_entity *entities.Player) interfaces.IPlayer {
	return models.NewPlayer(_entity)
}
