package world

import (
	"sync"

	"gameserver/interfaces"
)

type TilePointLayer struct {
	tilePointTileLayers map[int]*TilePointTileLayer

	// List of creatures whom are active on this TilePointLayer
	creatures      interfaces.CreatureMap
	creaturesMutex sync.RWMutex

	blocking interfaces.TileBlocking
	layer    int
}

func NewTilePointLayer(_layer int) *TilePointLayer {
	return &TilePointLayer{tilePointTileLayers: make(map[int]*TilePointTileLayer),
		creatures: make(interfaces.CreatureMap),
		layer:     _layer,
		blocking:  interfaces.TILEBLOCK_BLOCK}
}

func (t *TilePointLayer) GetBlocking() interfaces.TileBlocking {
	return t.blocking
}

func (t *TilePointLayer) SetBlocking(_blocking interfaces.TileBlocking) {
	t.blocking = _blocking
}

func (t *TilePointLayer) GetTilePointTileLayers() map[int]*TilePointTileLayer {
	return t.tilePointTileLayers
}

// Gets a TilePointTileLayer using layer index.
// Returns:	TilePointTileLayer object
//			Bool, true if the TilePointTileLayer was found, otherwise false
func (t *TilePointLayer) GetTilePointTileLayer(_layer int) (*TilePointTileLayer, bool) {
	tptl, found := t.tilePointTileLayers[_layer]
	return tptl, found
}

func (t *TilePointLayer) AddCreature(_creature interfaces.ICreature, _checkEvents bool) ReturnValue {
	ret := RET_NOERROR

	/*if _checkEvents && t.Events.Len() > 0 {
		for e := t.events.Front(); e != nil; e = e.Next() {
			event, valid := e.Value.(pulogic.ITileEvent)
			if valid {
				ret = event.OnCreatureEnter(_creature, ret)
			}

			if ret == RET_NOTPOSSIBLE {
				return
			}
		}
	}*/

	t.creaturesMutex.Lock()
	defer t.creaturesMutex.Unlock()
	_, found := t.creatures[_creature.GetUID()]
	if !found {
		t.creatures[_creature.GetUID()] = _creature
	}

	return ret
}

func (t *TilePointLayer) RemoveCreature(_creature interfaces.ICreature, _checkEvents bool) ReturnValue {
	ret := RET_NOERROR

	/*if _checkEvents && t.Events.Len() > 0 {
		for e := t.Events.Front(); e != nil; e = e.Next() {
			event, valid := e.Value.(pulogic.ITileEvent)
			if valid {
				ret = event.OnCreatureLeave(_creature, ret)
			}

			if ret == RET_NOTPOSSIBLE {
				return
			}
		}
	}*/

	t.creaturesMutex.Lock()
	defer t.creaturesMutex.Unlock()
	delete(t.creatures, _creature.GetUID())

	return ret
}

// Gets a TilePointTileLayer using layer index. If the TilePointTileLayer doesn't exists it will be created.
// Returns:	TilePointTileLayer object
//			Bool, true if TilePointTileLayer is new, otherwise false
func (t *TilePointLayer) getOrAddTilePointTileLayer(_layer int, _tileId string) (*TilePointTileLayer, bool) {
	tptl, found := t.GetTilePointTileLayer(_layer)
	if !found {
		tptl = NewTilePointTileLayer(_layer, _tileId)
		t.tilePointTileLayers[_layer] = tptl
	}

	return tptl, !found
}
