package systems

import (
	"github.com/anaseto/gruid/paths"
	"github.com/lecoqjacob/rrouge/ecs"
	"github.com/lecoqjacob/rrouge/rrouge_game/dungeon"
	"github.com/lecoqjacob/rrouge/rrouge_game/game"
	"github.com/lecoqjacob/rrouge/rrouge_game/turnstate"
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

func (*AISystem) Priority() int { return world.Priority_AI }

func (ais *AISystem) Update(ai_entities ecs.EntityList, dt float32) {
	if ais.G.TurnState != turnstate.MonsterTurn {
		return
	}

	w := ais.W
	dgen := ais.Dgen
	pp := ais.W.GetPosition(w.Player)

	for _, ai := range ai_entities {
		p := w.GetPosition(ai)
		ai_cmp := ai.GetComponent(world.ComponentTypes.AI).(world.AIComponent)

		if paths.DistanceManhattan(p, pp) == 1 {
			// If the monster is adjacent to the player, attack.
			// g.BumpAttack(i, g.ECS.PlayerID)
			w.ApplyMelee(ai, w.Player)
			continue
		}

		// NOTE: this base AI can be improved for example to avoid
		// monster's getting stuck between them. It's enough to get
		// started, though.
		if !w.IsInFOV(ai, pp) {
			if len(ai_cmp.Path) < 1 {
				// Pick new path to a random floor tile.
				ai_cmp.Path = dgen.PR.AstarPath(ai_cmp.AiPath, p, dgen.RandomFloor())
			}

			ais.aiMove(ai, ai_cmp)
			continue
		}

		// The monster is in player's FOV, so we compute a suitable path to
		// reach the player.
		ai_cmp.Path = dgen.PR.AstarPath(ai_cmp.AiPath, p, pp)
		ais.aiMove(ai, ai_cmp)
	}
}

func (ais *AISystem) aiMove(ai *ecs.Entity, ai_cmp world.AIComponent) {
	w := ais.W
	dgen := ais.Dgen

	if len(ai_cmp.Path) > 0 && ai_cmp.Path[0] == w.GetPosition(ai) {
		ai_cmp.Path = ai_cmp.Path[1:]
	}

	if len(ai_cmp.Path) > 0 && !dgen.Blocked[ai_cmp.Path[0]] {
		// Only move if there is no blocking entity.
		w.ApplyMovement(ai, ai_cmp.Path[0])
		ai_cmp.Path = ai_cmp.Path[1:]
	}

	// Update AI Component
	ai.InsertComponent(ai_cmp)
}
