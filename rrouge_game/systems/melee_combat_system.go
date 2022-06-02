package systems

import (
	"log"

	"github.com/lecoqjacob/rrouge/ecs"
	"github.com/lecoqjacob/rrouge/rrouge_game/utils"
	"github.com/lecoqjacob/rrouge/rrouge_game/world"
)

type MeleeCombatSystem struct {
	W *world.World
}

func (mcs *MeleeCombatSystem) Components() []ecs.ComponentType {
	return []ecs.ComponentType{
		world.ComponentTypes.WantsToMelee,
		world.ComponentTypes.Description,
		world.ComponentTypes.Stats,
	}
}

func (*MeleeCombatSystem) Priority() int { return world.Priority_MeleeCombat }

func (mcs *MeleeCombatSystem) Update(entities ecs.EntityList, dt float32) {
	for _, entity := range entities {
		wants_melee := entity.GetComponent(world.ComponentTypes.WantsToMelee).(world.WantsToMelee)

		stats := mcs.W.GetStats(entity)
		name := mcs.W.GetDescription(entity)

		if stats.HP > 0 {
			target_stats := mcs.W.GetStats(wants_melee.Target)

			if target_stats.HP > 0 {
				target_name := mcs.W.GetDescription(wants_melee.Target)
				damage := utils.Max(0, stats.Power-target_stats.Defense)

				if damage == 0 {
					log.Printf("%s is unable to hurt %s\n", name, target_name)
				} else {
					log.Printf("%s hits %s, for %d hp.\n", name, target_name, damage)
					mcs.W.ApplyDamage(wants_melee.Target, damage)
				}
			}
		}
	}

	entities.ClearComponents(world.ComponentTypes.WantsToMelee)
}
