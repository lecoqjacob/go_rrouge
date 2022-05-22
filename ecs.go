package main

import (
	"github.com/anaseto/gruid"
	"github.com/anaseto/gruid/rl"

	ecs "github.com/marioolofo/go-gameengine-ecs"
)

// ECS manages entities, as well as their positions. We don't go full “ECS”
// (Entity-Component-System) in this tutorial, opting for a simpler hybrid
// approach good enough for the tutorial purposes.
type ECS struct {
	PlayerID ecs.ID // index of Player's entity (for convenience)

	World   ecs.World
	Filters map[int]ecs.Filter
}

func InitializeWorld() ecs.World {
	// initial configuration to create the world, new components can be
	// added latter with world.RegisterComponents()
	config := []ecs.ComponentConfig{
		{ID: PointID, Component: &gruid.Point{}},
		{ID: RenderableID, Component: &Renderable{}},
		{ID: PlayerID, Component: &Player{}},
		{ID: FOVID, Component: &rl.FOV{}},
	}

	// NewWorld allocates a world and register the components
	world := ecs.NewWorld(config...)

	return world
}

// NewECS returns an initialized ECS structure.
func NewECS() *ECS {
	m := make(map[int]ecs.Filter)
	w := InitializeWorld()

	m[RenderableFilterID] = w.NewFilter(PointID, RenderableID)

	return &ECS{
		World:   w,
		Filters: m,
	}
}

// Add adds a new entity at a given position and returns its index/id.
func (es *ECS) AddEntity(p gruid.Point, r Renderable, components ...ecs.ID) ecs.ID {
	w := es.World

	// Initialize Entity
	entity := es.World.NewEntity()
	w.Assign(entity, append(components, PointID, RenderableID)...)

	// Positioning
	pos := (*gruid.Point)(w.Component(entity, PointID))
	*pos = p

	// Rendering
	ren := (*Renderable)(w.Component(entity, RenderableID))
	*ren = r

	return entity
}
