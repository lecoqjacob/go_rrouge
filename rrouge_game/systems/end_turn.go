package systems

import (
	"github.com/lecoqjacob/rrouge/ecs"
	"github.com/lecoqjacob/rrouge/rrouge_game/game"
	"github.com/lecoqjacob/rrouge/rrouge_game/world"
)

type EndTurnSystem struct {
	G *game.Game
}

func (s *EndTurnSystem) GetEntities() ecs.EntityList {
	return []*ecs.Entity{}
}

// Dummy tag
func (ets *EndTurnSystem) Remove(basic *ecs.Entity) {}

func (*EndTurnSystem) Priority() int { return world.Priority_EndTurn }

func (ets *EndTurnSystem) Update(entities ecs.EntityList, dt float32) {
	// ets.G.IsPlayerTurn = !ets.G.IsPlayerTurn
}
