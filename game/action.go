package game

import (
	"github.com/anaseto/gruid"
	"github.com/lecoqjacob/rrouge/systems"
)

// action represents information relevant to the last UI action performed.
type Action struct {
	Type  actionType  // kind of action (movement, quitting, ...)
	Delta gruid.Point // direction for ActionMovement
}

type actionType int

// These constants represent the possible UI actions.
const (
	NoAction       actionType = iota
	ActionMovement            // movement request
	ActionQuit                // quit the game
)

// handleAction updates the model in response to current recorded last action.
func (m *Model) handleAction() gruid.Effect {

	switch m.Action.Type {
	case ActionMovement:
		m.Game.MovePlayer(m.Action.Delta)
	case ActionQuit:
		// for now, just terminate with gruid End command: this will
		// have to be updated later when implementing saving.
		return gruid.End()
	}

	return nil
}

// MovePlayer moves the player to a given position and updates FOV information.
func (g *Game) MovePlayer(to gruid.Point) {
	systems.HandleMovementKey(g.Engine, g.GS, to)
	systems.Movement(g.Engine, g.GS)
}
