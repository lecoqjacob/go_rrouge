package main

import "github.com/anaseto/gruid"

// action represents information relevant to the last UI action performed.
type action struct {
	Type  actionType  // kind of action (movement, quitting, ...)
	Delta gruid.Point // direction for ActionBump
}

type actionType int

// These constants represent the possible UI actions.
const (
	NoAction actionType = iota

	ActionN
	ActionS
	ActionE
	ActionW

	ActionWait     // wait a turn
	ActionInteract // interact with cell
	ActionQuit     // quit the game
	ActionEscape
	ActionWaitTurn

	ActionBump         // bump request (attack or movement)
	ActionDrop         // menu to drop an inventory item
	ActionInventory    // inventory menu to use an item
	ActionPickup       // pickup an item on the ground
	ActionViewMessages // view history messages
	ActionExamine      // examine map
)

// handleAction updates the model in response to current recorded last action.
func (m *model) handleAction() gruid.Effect {
	return nil
}
