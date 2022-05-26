package systems

import (
	"fmt"

	"github.com/anaseto/gruid"
	"github.com/lecoqjacob/rrouge/constants"
	"github.com/lecoqjacob/rrouge/ecs"
	"github.com/lecoqjacob/rrouge/gamestate"
)

func HandleMovementKey(engine *ecs.Engine, gs *gamestate.GameState, p gruid.Point) {
	gs.Player.AddComponent(ecs.MoveComponent{To: p})
}

func Movement(engine *ecs.Engine, gs *gamestate.GameState) {
	entities := engine.Entities.GetEntities(constants.Move)

	for _, entity := range entities {
		move := entity.GetComponent(constants.Move).(ecs.MoveComponent)
		entity.RemoveComponent(constants.Move)

		current_pos, _ := entity.GetComponent(constants.Position).(ecs.PositionComponent)
		newPosition := ecs.PositionComponent{current_pos.Point.Add(move.To)}

		if !gs.Dgen.Walkable(newPosition.Point) {
			fmt.Println("not walkable :(")
			continue
		}

		fov := entity.GetComponent(constants.FOV).(ecs.FovComponent)
		fov.Dirty = true

		entity.ReplaceComponent(newPosition).ReplaceComponent(fov)

	}
}
