package world

import (
	"github.com/lecoqjacob/rrouge/ecs"
)

// Higher Priority means the system is executed sooner

const (
	Priority_AI       = 10
	Priority_Movement = 5
	Priority_Indexing = 0
	Priority_FOV      = 0
	Priority_EndTurn  = 0
	Priority_Render   = -1 // Last System to Run
)

func (w *World) AddSystem(system ecs.System) {
	w.engine.AddSystem(system)
}

func (w *World) AddSystems(systems []ecs.System) {
	for _, s := range systems {
		w.engine.AddSystem(s)
	}
}
