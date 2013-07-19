package game

import (
	"gameserver/interfaces"
)

func AddCreature(_creature interfaces.ICreature) {

}

func GetPlayerByUID(_uid uint64) (interfaces.IPlayer, bool) {
	return nil, false
}

func OnPlayerMove(_player interfaces.IPlayer, _direction uint16) {
	internalCreatureMove(_player, _direction)
}

/***********************
	Private Methods 
************************/
func internalCreatureMove(_creature interfaces.ICreature, _direction uint16) {

}
