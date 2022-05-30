package world

import (
	"github.com/anaseto/gruid"
	"github.com/anaseto/gruid/paths"
	"github.com/anaseto/gruid/rl"

	"github.com/lecoqjacob/rrouge/ecs"
	"github.com/lecoqjacob/rrouge/rrouge_game/log"
	"github.com/lecoqjacob/rrouge/rrouge_game/palette"
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

////////////
// Getters
////////////

func (ecs *World) GetStyle(entity *ecs.Entity) StyleComponent {
	return entity.GetComponent(ComponentTypes.Style).(StyleComponent)
}

func (ecs *World) GetPosition(entity *ecs.Entity) gruid.Point {
	cmp, ok := entity.GetComponent(ComponentTypes.Position).(PositionComponent)
	if ok {
		return cmp.Point
	}

	panic(log.PrintNotFoundComponentError(entity.Id, ComponentTypes.Description))
}

func (ecs *World) GetDescription(entity *ecs.Entity) string {
	cmp, ok := entity.GetComponent(ComponentTypes.Description).(DescriptionComponent)
	if ok {
		return cmp.Name
	}

	log.PrintNotFoundComponentError(entity.Id, ComponentTypes.Description)
	return "UnDeFiNeD!"
}

func (ecs *World) GetFOV(entity *ecs.Entity) FovComponent {
	cmp, ok := entity.GetComponent(ComponentTypes.FOV).(FovComponent)
	if ok {
		return cmp
	}

	panic(log.PrintNotFoundComponentError(entity.Id, ComponentTypes.Description))
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
		entity.RemoveComponent(pos_cmp.ComponentType())
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

////////////
// Misc
////////////

func (ecs *World) InFOV(entity *ecs.Entity, p gruid.Point) bool {
	pp := ecs.GetPosition(entity)
	fov := ecs.GetFOV(entity)

	return fov.View.Visible(p) &&
		paths.DistanceManhattan(pp, p) <= fov.Radius
}

// NoBlockingEntityAt returns true if there is no blocking entity at p (no
// player nor monsters in this tutorial).
// func (es *World) NoBlockingEntityAt(p gruid.Point) bool {
// 	_, ok := es.PosCache.GetOneByCoordAndComponents(p, ComponentTypes.BlocksCell)
// 	return !ok
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
	)

	w.ApplyPosition(player, start_pos)
	w.ApplyFOV(player, 6)

	w.Player = player
	return player
}

func (w *World) NewMonster(start_pos gruid.Point, r rune, name string) *ecs.Entity {
	monster := w.engine.CreateEntity()

	monster.InsertComponents(
		StyleComponent{Rune: r, Color: palette.ColorMonster},
		DescriptionComponent{Name: name},
		BlocksCellComponent{},
	)

	w.ApplyPosition(monster, start_pos)
	w.ApplyFOV(monster, 8)
	return monster
}
