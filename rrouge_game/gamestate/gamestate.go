package gamestate

import (
	"github.com/anaseto/gruid"
	"github.com/lecoqjacob/rrouge/ecs"
	"github.com/lecoqjacob/rrouge/rrouge_game/game"
	"github.com/lecoqjacob/rrouge/rrouge_game/systems"
	"github.com/lecoqjacob/rrouge/rrouge_game/turnstate"
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
		&systems.RenderSystem{G: g},
		&systems.MovementSystem{G: g, W: w},
		&systems.ComputeFOVSystem{G: g},
		&systems.MaxIndexingSystem{W: w, Dgen: dgen},
		&systems.AISystem{G: g, W: w, Dgen: dgen},
		&systems.DamageSystem{W: w},
		&systems.MeleeCombatSystem{W: w},
	})

	systems.RenderDebugMap(gs.Game, gs.Grid)

	// PreRun Systems
	gs.Game.World.Update(Time.Delta())

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

	// Handle State Machine
	if msg, ok := msg.(msgTurnState); ok {
		// Update systems
		gs.Game.World.Update(Time.Delta())

		// Progress to next state
		gs.Game.TurnState = turnstate.TurnState(msg)
	}

	var eff gruid.Effect
	gs.Action = Action{} // reset last action information

	switch msg := msg.(type) {
	case gruid.MsgKeyDown:
		// Update action information on key down.
		gs.updateMsgKeyDown(msg)
	}

	// Handle action (if any).
	eff = gs.handleAction()
	cmd := gs.turnStateCmd()

	if cmd != nil {
		return gruid.Batch(eff, cmd)
	}

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

type msgTurnState turnstate.TurnState

func (gs *GameState) turnStateCmd() gruid.Cmd {
	if gs.Game.TurnState == turnstate.AwaitingInput {
		return nil
	}

	return func() gruid.Msg {
		return msgTurnState(gs.Game.TurnState.NextState())
	}
}
