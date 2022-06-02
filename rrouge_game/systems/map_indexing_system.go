package systems

import (
	"github.com/lecoqjacob/rrouge/ecs"
	"github.com/lecoqjacob/rrouge/rrouge_game/dungeon"
	"github.com/lecoqjacob/rrouge/rrouge_game/utils"
	"github.com/lecoqjacob/rrouge/rrouge_game/world"
)

type MaxIndexingSystem struct {
	W    *world.World
	Dgen *dungeon.Dungeon
}

func (mis *MaxIndexingSystem) Components() []ecs.ComponentType {
	return []ecs.ComponentType{world.ComponentTypes.Position}
}

func (*MaxIndexingSystem) Priority() int { return world.Priority_Indexing }

func (mis *MaxIndexingSystem) Update(entities ecs.EntityList, dt float32) {
	mis.Dgen.Populate_Blocked()
	mis.Dgen.Clear_Content_Index()

	for _, e := range entities {
		p := mis.W.GetPosition(e)

		if e.HasComponent(world.ComponentTypes.BlocksCell) {
			mis.Dgen.Blocked[p] = true
		}

		utils.AppendSlice(mis.Dgen.Cell_Content[p], e)
	}
}
