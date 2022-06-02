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
		gs.Game.EndPlayerTurn()
	case ActionWait:
		// wait
		gs.Game.EndPlayerTurn()
	case ActionQuit:
		// for now, just terminate with gruid End command: this will
		// have to be updated later when implementing saving.
		return gruid.End()
	}

	return nil
}

// Bump moves the player to a given position and updates FOV information,
// or attacks if there is a monster.
func (gs *GameState) Bump(delta gruid.Point) {
	w := gs.Game.World
	to := w.GetPosition(w.Player).Add(delta)

	// Check if there is a target at new position
	if target, found := w.IsTargetAt(to); found {
		w.ApplyMelee(w.Player, target)
		return
	}

	w.ApplyMovement(w.Player, to)
}
