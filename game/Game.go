package game

import (
	"github.com/PokemonUniverse/nonamelib/log"
	"github.com/PokemonUniverse/nonamelib/position"

	"gameserver/interfaces"
	"gameserver/world"
)

var g *Game
type Game struct {
	creatures	map[uint64]interfaces.ICreature
}

func init() {
	g = &Game{	creatures: make(map[uint64]interfaces.ICreature) }
}

func AddPlayer(_player interfaces.IPlayer) {
	internalAddCreature(_player)
}

func GetPlayerByUID(_uid uint64) (interfaces.IPlayer, bool) {
	return nil, false
}

func OnPlayerMove(_player interfaces.IPlayer, _direction uint16) {
	basePosition := _player.GetPosition()
	retValue := internalCreatureMove(_player, _direction)
 	if retValue == world.RET_NOTPOSSIBLE {
		// Send current position to client, so position is synced
		_player.Walk(basePosition, position.ZP, false, _direction)
	}
}

/***********************
	Private Methods 
************************/
func internalAddCreature(_creature interfaces.ICreature) bool {
	_, found := g.creatures[_creature.GetUID()]
	
	if !found {
		g.creatures[_creature.GetUID()] = _creature
	} else {
		log.Warning("Game", "AddCreature", "Could not add creature. Creature with UID '%d' already exists", _creature.GetUID())
	}
	
	return !found
}

func internalCreatureMove(_creature interfaces.ICreature, _direction uint16) world.ReturnValue {
	ret := world.RET_NOTPOSSIBLE
	
	if _creature.CanMove() {
		sourcePosition := _creature.GetPosition()
		destinationPosition := _creature.GetPosition()
	
		switch _direction {
			case interfaces.DIR_NORTH:
				destinationPosition.Y -= 1
			case interfaces.DIR_SOUTH:
				destinationPosition.Y += 1
			case interfaces.DIR_WEST:
				destinationPosition.X -= 1
			case interfaces.DIR_EAST:
				destinationPosition.X += 1
		}	
		
		// Current TilePointLayer
		srcTpl, retValue := world.GetTilePointLayer(sourcePosition.MapId, sourcePosition.X, sourcePosition.Y, sourcePosition.Z)
		if retValue != world.RET_NOERROR {
			log.Debug("Game", "internalCreatureMove", "Source '[%d] %s' not found!", sourcePosition.MapId, sourcePosition.String())
			return retValue
		}
		
		// Destination TilePointLayer
		dstTpl, retValue := world.GetTilePointLayer(destinationPosition.MapId, destinationPosition.X, destinationPosition.Y, destinationPosition.Z)
		if retValue != world.RET_NOERROR {
			log.Debug("Game", "internalCreatureMove", "Destination '[%d] %s' not found!", destinationPosition.MapId, destinationPosition.String())
			return retValue
		}
		
		// Move creature to destination tile
		if ret = srcTpl.RemoveCreature(_creature, true); ret == world.RET_NOTPOSSIBLE {
			return ret
		}
		
		if ret = dstTpl.AddCreature(_creature, true); ret == world.RET_NOTPOSSIBLE {
			srcTpl.AddCreature(_creature, false) // Something went wrong. Put creature back on src TilePoint
			_creature.SetPosition(sourcePosition)
			return ret
		}
		
		if ret == world.RET_NOERROR {
			_creature.SetPosition(destinationPosition)	
			finalDestination := _creature.GetPosition()
			
			// Get list of creatures who're "new"
			visibleCreatures := world.GetVisibleCreaturesInDirection(finalDestination.MapId, finalDestination.X, finalDestination.Y, finalDestination.Z, _direction)
		
			for _, value := range(visibleCreatures) {
				if value != nil {
					_creature.AddVisibleCreature(value)
				}
			}
			
			// Tell Creature he has moved
			_creature.Walk(sourcePosition, destinationPosition, false, _direction)
		}
	}
	
	return ret
}
