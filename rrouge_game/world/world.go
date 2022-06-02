package world

import (
	"github.com/anaseto/gruid"
	"github.com/anaseto/gruid/paths"
	"github.com/anaseto/gruid/rl"

	"github.com/lecoqjacob/rrouge/ecs"
	"github.com/lecoqjacob/rrouge/rrouge_game/log"
	"github.com/lecoqjacob/rrouge/rrouge_game/palette"
	"github.com/lecoqjacob/rrouge/rrouge_game/utils"
)

type World struct {
	engine   *ecs.Engine
	PosCache PositionCache
	Player   *ecs.Entity
}

func init() {
	initComponentTypes()
}

func NewWorld() *World {
	engine := ecs.NewEngine()

	return &World{
		engine: engine,
		PosCache: PositionCache{
			Entities: map[string]EntitySet{},
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Utility
////////////////////////////////////////////////////////////////////////////////

func (w *World) Update(dt float32) {
	w.engine.Update(dt)
}

func (w *World) GetEntities(types ...ecs.ComponentType) ecs.EntityList {
	return w.engine.GetEntities(types...)
}

func (w *World) DestroyEntity(entity *ecs.Entity) {
	w.engine.DestroyEntity(entity)
}

////////////
// Getters
////////////

func (ecs *World) GetStyle(entity *ecs.Entity) StyleComponent {
	return entity.GetComponent(ComponentTypes.Style).(StyleComponent)
}

func (ecs *World) GetPosition(entity *ecs.Entity) gruid.Point {
	if cmp, found := entity.GetComponent(ComponentTypes.Position).(PositionComponent); found {
		return cmp.Point
	}

	panic(log.PrintNotFoundComponentError(entity.ID(), ComponentTypes.Description))
}

func (ecs *World) GetDescription(entity *ecs.Entity) string {
	if cmp, found := entity.GetComponent(ComponentTypes.Description).(DescriptionComponent); found {
		return cmp.Name
	}

	log.PrintNotFoundComponentError(entity.ID(), ComponentTypes.Description)
	return "UnDeFiNeD!"
}

func (ecs *World) GetFOV(entity *ecs.Entity) FovComponent {
	if cmp, found := entity.GetComponent(ComponentTypes.FOV).(FovComponent); found {
		return cmp
	}

	panic(log.PrintNotFoundComponentError(entity.ID(), ComponentTypes.Description))
}

func (ecs *World) GetStats(entity *ecs.Entity) StatsComponent {
	if cmp, found := entity.GetComponent(ComponentTypes.Stats).(StatsComponent); found {
		return cmp
	}

	panic(log.PrintNotFoundComponentError(entity.ID(), ComponentTypes.Description))
}

////////////
// Apply
////////////

func (ecs *World) ApplyPosition(entity *ecs.Entity, p gruid.Point) *ecs.Entity {
	// Remove previous position if it exist
	ecs.RemovePosition(entity)

	// Apply new position
	pos_cmp := PositionComponent{Point: p}
	entity.InsertComponent(pos_cmp)

	// Update cache
	ecs.PosCache.Add(pos_cmp.GetKey(), entity)

	return entity
}

func (ecs *World) RemovePosition(entity *ecs.Entity) *ecs.Entity {
	pos_cmp, ok := entity.GetComponent(ComponentTypes.Position).(PositionComponent)

	if ok {
		entity.RemoveComponents(pos_cmp.ComponentType())
		ecs.PosCache.Delete(pos_cmp.GetKey(), entity)
	}

	return entity
}

func (ecs *World) ApplyFOV(entity *ecs.Entity, radius int) *ecs.Entity {
	entity.InsertComponent(FovComponent{
		Dirty:  true,
		Radius: radius,
		View:   rl.NewFOV(gruid.NewRange(-radius, -radius, radius+1, radius+1)),
	})
	return entity
}

func (ecs *World) ApplyMovement(entity *ecs.Entity, to gruid.Point) {
	entity.InsertComponent(MoveComponent{To: to})
}

func (ecs *World) ApplyMelee(entity, target *ecs.Entity) {
	entity.InsertComponent(WantsToMelee{Target: target})
}

func (ecs *World) ApplyDamage(victim *ecs.Entity, amount int) {
	if cmp, ok := victim.GetComponent(ComponentTypes.SufferDamage).(SufferDamage); ok {
		utils.AppendSlice(cmp.Amount, amount)
	} else {
		dmg := SufferDamage{Amount: []int{amount}}
		victim.InsertComponent(dmg)
	}
}

func (w *World) ApplyDeath(victim *ecs.Entity) {
	victim.RemoveComponents(
		ComponentTypes.AI,
		ComponentTypes.Stats,
		ComponentTypes.BlocksCell,
	)

	// Update name
	description := victim.GetComponent(ComponentTypes.Description).(DescriptionComponent)
	description.Name += " corpse"

	// Update styling
	style := w.GetStyle(victim)
	style.Rune = '%'
	style.Color = palette.ColorDead

	victim.InsertComponents(description, style)
}

////////////
// Misc
////////////

func (w *World) IsInFOV(entity *ecs.Entity, p gruid.Point) bool {
	pp := w.GetPosition(entity)
	fov := w.GetFOV(entity)

	return fov.View.Visible(p) &&
		paths.DistanceManhattan(pp, p) <= fov.Radius
}

func (w *World) IsTargetAt(p gruid.Point) (*ecs.Entity, bool) {
	target, found := w.PosCache.GetOneByCoordAndComponents(p, ComponentTypes.Stats)
	return target, found
}

// NoBlockingEntityAt returns true if there is no blocking entity at p (no
// player nor monsters in this tutorial).
// func (es *World) NoBlockingEntityAt(p gruid.Point) bool {
// 	_, isBlocked := es.PosCache.GetOneByCoordAndComponents(p, ComponentTypes.BlocksCell)
// 	return !isBlocked
// }

////////////////////////////////////////////////////////////////////////////////
// Spawner
////////////////////////////////////////////////////////////////////////////////

func (w *World) NewPlayer(start_pos gruid.Point) *ecs.Entity {
	player := w.engine.CreateEntity()

	player.InsertComponents(
		PlayerComponent{},
		StyleComponent{Rune: '@', Color: palette.ColorPlayer},
		DescriptionComponent{Name: "Player"},
		BlocksCellComponent{},
		StatsComponent{StatsValues: &StatsValues{Max_HP: 30, HP: 30, Defense: 2, Power: 5, Fov: 8}},
	)

	w.ApplyPosition(player, start_pos)
	w.ApplyFOV(player, 8)

	w.Player = player
	return player
}

func (w *World) NewMonster(start_pos gruid.Point, r rune, name string) *ecs.Entity {
	monster := w.engine.CreateEntity()

	monster.InsertComponents(
		StyleComponent{Rune: r, Color: palette.ColorMonster},
		DescriptionComponent{Name: name},
		BlocksCellComponent{},
		StatsComponent{StatsValues: &StatsValues{Max_HP: 16, HP: 16, Defense: 1, Power: 4, Fov: 4}},
	)

	w.ApplyPosition(monster, start_pos)
	w.ApplyFOV(monster, 4)
	return monster
}
