package game

import (
	"math/rand"
	"time"

	"github.com/lecoqjacob/rrouge/rrouge_game/dungeon"
	"github.com/lecoqjacob/rrouge/rrouge_game/random"
	"github.com/lecoqjacob/rrouge/rrouge_game/turnstate"
	"github.com/lecoqjacob/rrouge/rrouge_game/world"
)

const Version = "1.0.0"

type Game struct {
	World *world.World
	Dgen  *dungeon.Dungeon
	Rng   *random.Rand

	TurnState turnstate.TurnState
	Log       []LogEntry // log entries
}

func InitializeGame() *Game {
	// Initialize World
	w := world.NewWorld()

	// Initialize map
	dgen := dungeon.NewDungeon()

	// Player
	w.NewPlayer(dgen.Rooms[0].Center())

	rng := random.New(rand.NewSource(time.Now().UnixNano()))
	for _, room := range dgen.Rooms[1:3] {
		// for _, room := range dgen.Rooms[1:] {
		var r rune
		var name string
		switch rng.Roll_Dice(1, 2) {
		case 1:
			r, name = 'g', "Goblin"
		default:
			r, name = 'o', "Orc"
		}

		m := w.NewMonster(room.Center(), r, name)
		m.InsertComponent(world.AIComponent{AiPath: &world.AiPath{Dgen: dgen}})
	}

	return &Game{
		World: w,
		Dgen:  dgen,
		Rng:   rng,
	}
}

func (g *Game) EndPlayerTurn() {
	g.TurnState = turnstate.PlayerTurn
}
