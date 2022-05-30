package gamestate

import (
	"github.com/anaseto/gruid"
	"github.com/lecoqjacob/rrouge/ecs"
	"github.com/lecoqjacob/rrouge/rrouge_game/game"
	"github.com/lecoqjacob/rrouge/rrouge_game/systems"
	"github.com/lecoqjacob/rrouge/rrouge_game/utils"
)

var Time *Clock

type GameState struct {
	Game   *game.Game
	Grid   gruid.Grid // drawing grid
	Action Action     // UI action
}

func init() {
	Time = NewClock()
}

func (gs *GameState) NewGameState() gruid.Effect {
	gs.Game = game.InitializeGame()

	g := gs.Game
	w := gs.Game.World
	dgen := gs.Game.Dgen

	// Add systems to the world
	gs.Game.World.AddSystems([]ecs.System{
		&systems.MovementSystem{G: g, W: w},
		&systems.RenderSystem{G: g},
		&systems.ComputeFOVSystem{G: g},
		&systems.EndTurnSystem{G: g},
		&systems.MaxIndexingSystem{Dgen: dgen},
		&systems.AISystem{G: g, W: w, Dgen: dgen},
	})

	systems.RenderDebugMap(gs.Game, gs.Grid)
	return nil
}

// Update implements gruid.Model.Update. It handles keyboard and mouse input
// messages and updates the model in response to them.
func (gs *GameState) Update(msg gruid.Msg) gruid.Effect {
	Time.Tick()

	// Initialize Game
	if _, ok := msg.(gruid.MsgInit); ok {
		return gs.NewGameState()
	}

	// Handle Quit Action
	if _, ok := msg.(gruid.MsgQuit); ok {
		// md.mode = modeQuit
		return gruid.End()
	}

	gs.Action = Action{} // reset last action information
	switch msg := msg.(type) {
	case gruid.MsgKeyDown:
		// Update action information on key down.
		gs.updateMsgKeyDown(msg)
	}

	// Handle action (if any).
	eff := gs.handleAction()

	// Update Systems
	gs.Game.World.Update(Time.Delta())

	return eff
}

func (gs *GameState) updateMsgKeyDown(msg gruid.MsgKeyDown) {
	switch msg.Key {
	case gruid.KeyArrowLeft, "h", gruid.KeyArrowDown, "j", gruid.KeyArrowUp, "k", gruid.KeyArrowRight, "l":
		gs.Action = Action{Type: ActionBump, Delta: utils.KeyToDir(msg.Key)}
	case gruid.KeyEnter, gruid.KeySpace:
		gs.Action = Action{Type: ActionWait}
	case gruid.KeyEscape, "q":
		gs.Action = Action{Type: ActionQuit}
	}
}

// Draw implements gruid.Model.Draw. It draws a simple map that spans the whole
// grid.
func (gs *GameState) Draw() gruid.Grid {
	return systems.GetRenderables(gs.Grid)
}
