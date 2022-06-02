package systems

import (
	"github.com/lecoqjacob/rrouge/ecs"
	"github.com/lecoqjacob/rrouge/rrouge_game/game"
	"github.com/lecoqjacob/rrouge/rrouge_game/world"
)

type MovementSystem struct {
	G *game.Game
	W *world.World
}

func (ms *MovementSystem) Components() []ecs.ComponentType {
	return []ecs.ComponentType{world.ComponentTypes.Move}
}

func (*MovementSystem) Priority() int { return world.Priority_Movement }

func (ms *MovementSystem) Update(entities ecs.EntityList, dt float32) {
	g := ms.G

	for _, entity := range entities {
		move := entity.GetComponent(world.ComponentTypes.Move).(world.MoveComponent)

		np := move.To
		if !g.Dgen.Walkable(np) || g.Dgen.Blocked[np] { // !g.World.NoBlockingEntityAt(np.Point)
			continue
		}

		// Update POS
		g.World.ApplyPosition(entity, np)
		g.Dgen.Blocked[np] = true

		// Update FOV, if applicable
		if fov, ok := entity.GetComponent(world.ComponentTypes.FOV).(world.FovComponent); ok {
			fov.Dirty = true
			entity.InsertComponent(fov)
		}
	}

	entities.ClearComponents(world.ComponentTypes.Move)
}
