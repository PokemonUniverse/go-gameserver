package pokemon

import (
	"gameserver/data/entities"
)

type Move struct {
	MoveEntity *entities.Move
}

func NewMove(_entity *entities.Move) *Move {
	return &Move{MoveEntity: _entity}
}

func (m *Move) GetMoveId() int64 {
	return int64(m.MoveEntity.MoveId)
}

func (m *Move) GetIdentifier() string {
	return m.MoveEntity.Identifier
}
