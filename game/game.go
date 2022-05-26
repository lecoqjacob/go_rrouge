package game

import (
	"github.com/lecoqjacob/rrouge/gamestate"
	"github.com/lecoqjacob/rrouge/systems"

	"github.com/lecoqjacob/rrouge/ecs"
)

const Version = "1.0.0"

type Game struct {
	Engine *ecs.Engine
	GS     *gamestate.GameState
}

func NewGame() *Game {
	// Initialize entities
	engine := ecs.NewEngine()
	gs := gamestate.NewGameState(engine)
	systems.ComputeFov(engine, gs)

	return &Game{
		// ECS engine
		Engine: engine,
		GS:     gs,
	}
}
