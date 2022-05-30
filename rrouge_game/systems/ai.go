package systems

import (
	"fmt"

	"github.com/anaseto/gruid/paths"
	"github.com/lecoqjacob/rrouge/ecs"
	"github.com/lecoqjacob/rrouge/rrouge_game/dungeon"
	"github.com/lecoqjacob/rrouge/rrouge_game/game"
	"github.com/lecoqjacob/rrouge/rrouge_game/log"
	"github.com/lecoqjacob/rrouge/rrouge_game/world"
)

type AISystem struct {
	G    *game.Game
	W    *world.World
	Dgen *dungeon.Dungeon
}

func (ais *AISystem) Components() []ecs.ComponentType {
	return []ecs.ComponentType{world.ComponentTypes.AI, world.ComponentTypes.FOV}
}

// Dummy tag
func (ais *AISystem) Remove(basic *ecs.Entity) {
}

func (*AISystem) Priority() int { return world.Priority_AI }

func (ais *AISystem) Update(ai_entities ecs.EntityList, dt float32) {
	if ais.G.IsPlayerTurn {
		return
	}

	fmt.Printf("AISystem.Update\n")

	g := ais.G
	w := ais.W
	dgen := ais.Dgen
	pp := ais.W.GetPosition(w.Player)

	for _, ai := range ai_entities {
		p := w.GetPosition(ai)
		ai_cmp := ai.GetComponent(world.ComponentTypes.AI).(world.AIComponent)

		if paths.DistanceManhattan(p, pp) == 1 {
			log.Debug("%s shouts `I can smell your breath!`\n", w.GetDescription(ai))
			// If the monster is adjacent to the player, attack.
			// g.BumpAttack(i, g.ECS.PlayerID)
			return
		}

		// NOTE: this base AI can be improved for example to avoid
		// monster's getting stuck between them. It's enough to get
		// started, though.
		if !w.InFOV(ai, pp) {
			if len(ai_cmp.Path) < 1 {
				// Pick new path to a random floor tile.
				ai_cmp.Path = dgen.PR.AstarPath(ai_cmp.AiPath, p, dgen.RandomFloor())
			}

			aiMove(ai, ai_cmp, w, dgen)
			return
		}

		// The monster is in player's FOV, so we compute a suitable path to
		// reach the player.
		ai_cmp.Path = dgen.PR.AstarPath(ai_cmp.AiPath, p, pp)

		aiMove(ai, ai_cmp, w, dgen)
	}

	g.IsPlayerTurn = true
}

func aiMove(ai *ecs.Entity, ai_cmp world.AIComponent, ecs *world.World, dgen *dungeon.Dungeon) {
	if len(ai_cmp.Path) > 0 && ai_cmp.Path[0] == ecs.GetPosition(ai) {
		ai_cmp.Path = ai_cmp.Path[1:]
	}

	if len(ai_cmp.Path) > 0 && dgen.NoBlockingEntityAt(ai_cmp.Path[0]) {
		// Only move if there is no blocking entity.
		ecs.ApplyMovement(ai, ai_cmp.Path[0])
		ai_cmp.Path = ai_cmp.Path[1:]
	}

	ai.InsertComponent(ai_cmp)
}
