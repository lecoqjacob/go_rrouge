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

// Dummy tag
func (ms *MovementSystem) Remove(basic *ecs.Entity) {
}

func (*MovementSystem) Priority() int { return world.Priority_Movement }

func (ms *MovementSystem) Update(entities ecs.EntityList, dt float32) {
	g := ms.G

	for _, entity := range entities {
		move := entity.GetComponent(world.ComponentTypes.Move).(world.MoveComponent)
		entity.RemoveComponent(world.ComponentTypes.Move)

		newPosition := world.PositionComponent{Point: move.To}
		if !g.Dgen.Walkable(newPosition.Point) {
			continue
		}

		// Update POS
		g.World.ApplyPosition(entity, move.To)

		// Update FOV, if applicable
		fov, ok := entity.GetComponent(world.ComponentTypes.FOV).(world.FovComponent)
		if ok {
			fov.Dirty = true
			entity.InsertComponent(fov)
		}
	}
}
