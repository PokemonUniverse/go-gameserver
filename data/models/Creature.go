package models

import (
	"sync"
	
	"github.com/PokemonUniverse/nonamelib/position"
	
	"gameserver/interfaces"
)

type Creature struct {
	position	position.Position
	
	Name		string
	
	movementSpeed	int
	lastStep		int64
	
	visibleCreatures interfaces.CreatureMap
	visibleCreaturesMutex sync.RWMutex
}

func (c *Creature) init() {
	c.visibleCreatures = make(interfaces.CreatureMap)
	
	c.movementSpeed = 250
	c.lastStep = interfaces.PUSYS_TIME()
}

func (c *Creature) GetUID() uint64 {
	return 0
}

func (c *Creature) GetName() string {
	return c.Name
}

func (c *Creature) GetCreatureType() interfaces.CreatureType {
	return interfaces.CREATURE_TYPE_UNKNOWN
}

func (c *Creature) LoadCharacterData() bool {
	return true
}

func (c *Creature) GetMapId() int {
	return c.position.MapId
}

func (c *Creature) GetPosition() position.Position {
	return c.position
}

func (c *Creature) SetPosition(_position position.Position) {
	c.position = _position
}

func (c *Creature) CanMove() bool {
	return (c.getTimeSinceLastMove() >= c.movementSpeed)
}

func (c *Creature) Walk(_from position.Position, _to position.Position, _teleported bool, _direction uint16) {
	c.visibleCreaturesMutex.RLock()
	defer c.visibleCreaturesMutex.RUnlock()
	
	for _, creature := range(c.visibleCreatures) {
		creature.OnCreatureMove(c, _from, _to, _teleported)
	}
}

// Methods for all moving creatures
func (c *Creature) OnCreatureMove(_creature interfaces.ICreature, _from position.Position, _to position.Position, _teleport bool) {
	if _creature.GetUID() == c.GetUID() {
		c.lastStep = interfaces.PUSYS_TIME()
	}
}

func (c *Creature) OnCreatureTurn(_creature interfaces.ICreature) {
}

func (c *Creature) OnCreatureAppear(_creature interfaces.ICreature, _isLogin bool) {
}

func (c *Creature) OnCreatureDisappear(_creature interfaces.ICreature, _isLogout bool) {
}

// Methods for all creatures who need to see other creatures	
func (c *Creature) AddVisibleCreature(_creature interfaces.ICreature) bool {
	if _creature.GetUID() != c.GetUID()	{
		c.visibleCreaturesMutex.Lock()
		defer c.visibleCreaturesMutex.Unlock()
	
		if _, found := c.visibleCreatures[_creature.GetUID()]; !found {
			c.visibleCreatures[_creature.GetUID()] = _creature
			
			return true
		}
	}
	
	return false
}

func (c *Creature) RemoveVisibleCreature(_creature interfaces.ICreature) bool {
	if _creature.GetUID() != c.GetUID()	{
		c.visibleCreaturesMutex.Lock()
		defer c.visibleCreaturesMutex.Unlock()
	
		delete(c.visibleCreatures, _creature.GetUID())
		
		return true
	}
	
	return false
}

func (c *Creature) KnowsVisibleCreature(_creature interfaces.ICreature) bool {
	c.visibleCreaturesMutex.RLock()
	defer c.visibleCreaturesMutex.RUnlock()

	_, found := c.visibleCreatures[_creature.GetUID()]
	return found
}

func (c *Creature) GetVisibleCreatures() interfaces.CreatureMap {
	return c.visibleCreatures
}

func (c *Creature) CanSeePosition(_position position.Position) bool {
	if c.GetMapId() != _position.MapId {
		return false
	}

	return c.position.IsInRange3p(_position, interfaces.CLIENT_VIEWPORT_CENTER)
}

func (c *Creature) getTimeSinceLastMove() int {
	return int(interfaces.PUSYS_TIME() - c.lastStep)
}
