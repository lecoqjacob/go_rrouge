package systems

import (
	"github.com/anaseto/gruid"
	"github.com/anaseto/gruid/paths"

	"github.com/lecoqjacob/rrouge/ecs"
	"github.com/lecoqjacob/rrouge/rrouge_game/dungeon"
	"github.com/lecoqjacob/rrouge/rrouge_game/game"
	"github.com/lecoqjacob/rrouge/rrouge_game/world"
)

type ComputeFOVSystem struct {
	G *game.Game
}

func (fov_s *ComputeFOVSystem) Components() []ecs.ComponentType {
	return []ecs.ComponentType{world.ComponentTypes.Position, world.ComponentTypes.FOV}
}

func (*ComputeFOVSystem) Priority() int { return world.Priority_FOV }

func (fov_s *ComputeFOVSystem) Update(entities ecs.EntityList, dt float32) {
	g := fov_s.G
	w := g.World

	for _, entity := range entities {
		fov := entity.GetComponent(world.ComponentTypes.FOV).(world.FovComponent)

		if fov.Dirty {
			fov.Dirty = false
			pos := w.GetPosition(entity)

			// We shift the FOV's Range so that it will be centered on the new position.
			fov.View.SetRange(visionRange(pos, fov.Radius))

			// Update FOV Component
			entity.InsertComponent(fov)

			// We mark cells in field of view as explored. We use the symmetric
			// shadow casting algorithm provided by the rl package.
			vision_map := fov.View.SSCVisionMap(pos, fov.Radius, g.Dgen.BlocksSSCLOS, false)

			if entity == w.Player {
				for _, p := range vision_map {
					if paths.DistanceManhattan(p, pos) > fov.Radius {
						continue
					}

					if !g.Dgen.Explored[p] {
						g.Dgen.Explored[p] = true
					}
				}
			}
		}
	}
}

////////////////////////////////////////////////////////////////////////////////
/// Utilities
////////////////////////////////////////////////////////////////////////////////

func visionRange(p gruid.Point, radius int) gruid.Range {
	drg := gruid.NewRange(0, 0, dungeon.DungeonWidth, dungeon.DungeonHeight)
	delta := gruid.Point{X: radius, Y: radius}
	return drg.Intersect(gruid.Range{Min: p.Sub(delta), Max: p.Add(delta).Shift(1, 1)})
}
