package world

import (
	"github.com/lecoqjacob/rrouge/ecs"
)

// Higher Priority means the system is executed sooner
type Bits uint8

const (
	Priority_Render = -1 // Last System to Run
	Priority_Damage = iota
	Priority_MeleeCombat
	Priority_Indexing
	Priority_Movement
	Priority_AI
	Priority_FOV
)

func (w *World) AddSystem(system ecs.System) {
	w.engine.AddSystem(system)
}

func (w *World) AddSystems(systems []ecs.System) {
	for _, s := range systems {
		w.engine.AddSystem(s)
	}
}
