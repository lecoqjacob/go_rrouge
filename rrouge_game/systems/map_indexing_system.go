package systems

import (
	"github.com/lecoqjacob/rrouge/ecs"
	"github.com/lecoqjacob/rrouge/rrouge_game/dungeon"
	"github.com/lecoqjacob/rrouge/rrouge_game/world"
)

type MaxIndexingSystem struct {
	Dgen *dungeon.Dungeon
}

func (mis *MaxIndexingSystem) Components() []ecs.ComponentType {
	return []ecs.ComponentType{world.ComponentTypes.Position, world.ComponentTypes.BlocksCell}
}

// Dummy tag
func (mis *MaxIndexingSystem) Remove(basic *ecs.Entity) {}

func (*MaxIndexingSystem) Priority() int { return world.Priority_Indexing }

func (mis *MaxIndexingSystem) Update(entities ecs.EntityList, dt float32) {
	mis.Dgen.Populate_Blocked()

	for _, e := range entities {
		p := e.GetComponent(world.ComponentTypes.Position).(world.PositionComponent).Point
		mis.Dgen.Blocked[p] = true
	}
}
