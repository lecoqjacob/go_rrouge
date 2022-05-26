package systems

import (
	"github.com/anaseto/gruid"
	"github.com/anaseto/gruid/paths"
	"github.com/lecoqjacob/rrouge/cell"
	"github.com/lecoqjacob/rrouge/constants"
	"github.com/lecoqjacob/rrouge/ecs"
	"github.com/lecoqjacob/rrouge/gamestate"
	"github.com/lecoqjacob/rrouge/utils"
)

func ComputeFov(engine *ecs.Engine, gs *gamestate.GameState) {
	entities := engine.Entities.GetEntities(constants.Position, constants.FOV)
	dgen := gs.Dgen

	for _, entity := range entities {
		fov := entity.GetComponent(constants.FOV).(ecs.FovComponent)

		if fov.Dirty {
			fov.Dirty = false
			pos := entity.GetPosition().Point

			// We shift the FOV's Range so that it will be centered on the new position.
			fov.View.SetRange(utils.VisionRange(pos, fov.Radius))

			// Update FOV Component
			entity.ReplaceComponent(fov)

			if entity == gs.Player {
				// We mark cells in field of view as explored. We use the symmetric
				// shadow casting algorithm provided by the rl package.
				passable := func(p gruid.Point) bool {
					return cell.Cell(dgen.Grid.At(p)) != cell.WallCell
				}

				for _, p := range fov.View.SSCVisionMap(pos, constants.MaxLOS, passable, false) {
					if paths.DistanceManhattan(p, pos) > constants.MaxLOS {
						continue
					}

					c := dgen.Grid.At(p)
					if !cell.Explored(c) {
						dgen.SetExplored(p)
						// dgen.Grid.Set(p, c|cell.ExploredCell)
					}
				}
			}
		}
	}
}
