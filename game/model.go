package game

import (
	"github.com/anaseto/gruid"
	"github.com/lecoqjacob/rrouge/gamestate"
	"github.com/lecoqjacob/rrouge/systems"
	"github.com/lecoqjacob/rrouge/utils"
)

type mode int

const (
	modeNormal mode = iota
	modeWelcome
	modeQuit
	modeQuitConfirmation

	AwaitingInput
	PlayerTurn
	MonsterTurn
)

type Model struct {
	Grid   gruid.Grid // drawing grid
	Game   *Game      // Game state
	Gs     *gamestate.GameState
	Action Action // UI action
}

// Update implements gruid.Model.Update. It handles keyboard and mouse input
// messages and updates the model in response to them.
func (m *Model) Update(msg gruid.Msg) gruid.Effect {
	m.Action = Action{} // reset last action information

	switch msg := msg.(type) {
	case gruid.MsgInit:
		m.Game = NewGame()
	case gruid.MsgKeyDown:
		// Update action information on key down.
		m.updateMsgKeyDown(msg)
	}

	// Handle action (if any).
	return m.handleAction()
}

func (m *Model) updateMsgKeyDown(msg gruid.MsgKeyDown) {
	switch msg.Key {
	case gruid.KeyArrowLeft, "h", gruid.KeyArrowDown, "j", gruid.KeyArrowUp, "k", gruid.KeyArrowRight, "l":
		m.Action = Action{Type: ActionMovement, Delta: utils.KeyToDir(msg.Key)}
	case gruid.KeyEscape, "q":
		m.Action = Action{Type: ActionQuit}
	}

	systems.ComputeFov(m.Game.Engine, m.Game.GS)
}

// Draw implements gruid.Model.Draw. It draws a simple map that spans the whole
// grid.
func (m *Model) Draw() gruid.Grid {
	systems.Render(m.Game.Engine, m.Game.GS, m.Grid)
	return m.Grid
}
