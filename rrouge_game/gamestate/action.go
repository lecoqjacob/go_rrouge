package gamestate

import (
	"github.com/anaseto/gruid"
)

// action represents information relevant to the last UI action performed.
type Action struct {
	Type  actionType  // kind of action (movement, quitting, ...)
	Delta gruid.Point // direction for ActionMovement
}

type actionType int

// These constants represent the possible UI actions.
const (
	NoAction   actionType = iota
	ActionBump            // movement request
	ActionWait            // wait a turn
	ActionQuit            // quit the game
)

func (gs *GameState) handleAction() gruid.Effect {
	switch gs.Action.Type {
	case ActionBump:
		gs.Bump(gs.Action.Delta)
	case ActionWait:
		// systems.EndTurn(gs.Game)
	case ActionQuit:
		// for now, just terminate with gruid End command: this will
		// have to be updated later when implementing saving.
		return gruid.End()
	}

	gs.Game.IsPlayerTurn = false

	return nil
}

// Bump moves the player to a given position and updates FOV information,
// or attacks if there is a monster.
func (gs *GameState) Bump(delta gruid.Point) {
	w := gs.Game.World
	to := w.GetPosition(w.Player).Add(delta)
	w.ApplyMovement(w.Player, to)
}
