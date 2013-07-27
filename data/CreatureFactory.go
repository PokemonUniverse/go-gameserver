package data

import (
	"gameserver/data/models"
	"gameserver/interfaces"
)

func CreatePlayer() interfaces.IPlayer {
	return models.NewPlayer()
}
