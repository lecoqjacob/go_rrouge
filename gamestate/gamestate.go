package gamestate

import (
	"math/rand"
	"time"

	"github.com/lecoqjacob/rrouge/constants"
	"github.com/lecoqjacob/rrouge/dungeon"
	"github.com/lecoqjacob/rrouge/ecs"
	"github.com/lecoqjacob/rrouge/random"
)

type GameState struct {
	Engine *ecs.Engine
	Dgen   *dungeon.Dungeon // the game map, made of tiles
	Player *ecs.Entity

	Rand *random.Rand
}

func NewGameState(engine *ecs.Engine) *GameState {
	// Initialize map
	dgen := dungeon.NewDungeon()

	player_start_pos := dgen.Player_Start()

	// Player
	player := ecs.NewPlayer(engine.NewEntity())
	player.ApplyPosition(player_start_pos)
	player.ApplyFOV(constants.MaxLOS)

	return &GameState{
		// ECS engine
		Engine: engine,
		Dgen:   dgen,
		Player: player,

		Rand: random.New(rand.NewSource(time.Now().UnixNano())),
	}
}
