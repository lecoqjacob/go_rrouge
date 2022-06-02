package systems

import (
	"log"

	"github.com/lecoqjacob/rrouge/ecs"
	"github.com/lecoqjacob/rrouge/rrouge_game/utils"
	"github.com/lecoqjacob/rrouge/rrouge_game/world"
)

type DamageSystem struct {
	W *world.World
}

func (ds *DamageSystem) Components() []ecs.ComponentType {
	return []ecs.ComponentType{world.ComponentTypes.Stats, world.ComponentTypes.SufferDamage}
}

func (*DamageSystem) Priority() int { return world.Priority_Damage }

func (ds *DamageSystem) Update(entities ecs.EntityList, dt float32) {
	for _, entity := range entities {
		stats := ds.W.GetStats(entity)
		suffer_damage := entity.GetComponent(world.ComponentTypes.SufferDamage).(world.SufferDamage)

		stats.HP -= utils.Sum(suffer_damage.Amount)
		// Dead
		if stats.HP <= 0 {
			log.Printf("%s has been slain\n", ds.W.GetDescription(entity))
			ds.W.ApplyDeath(entity)
			return
		}

		entity.InsertComponent(stats)
	}

	entities.ClearComponents(world.ComponentTypes.SufferDamage)
}
