package ecs

import (
	"sort"
	"sync/atomic"
)

type EntityList []*Entity

type Engine struct {
	Entities EntityList
	systems  Systems
}

////////////////////////////////
// Private functions
////////////////////////////////
func (engine *Engine) addEntity(newEntity *Entity) *Entity {
	// Add entity to the engine
	engine.Entities = append(engine.Entities, newEntity)
	return newEntity
}

////////////////////////////////
// Public functions
////////////////////////////////

func NewEngine() *Engine {
	return &Engine{}
}

////////////////////////////////////////////////////////////////////////////////
// Entity Utility Functions
////////////////////////////////////////////////////////////////////////////////

func (engine *Engine) CreateEntity() *Entity {
	newEntity := &Entity{
		id:         atomic.AddUint64(&idInc, 1),
		Components: make(map[ComponentType]Component),
	}

	engine.Entities = append(engine.Entities, newEntity)

	return newEntity
}

func (engine *Engine) DestroyEntity(entity *Entity) {
	// Remove all components to trigger possible actions
	for k := range entity.Components {
		entity.RemoveComponent(k)
	}
	engine.Entities = engine.Entities.RemoveEntity(entity)

	entity.Components = nil
	entity.Engine = nil
	entity = nil
}

func (engine *Engine) GetEntities(types ...ComponentType) EntityList {
	return engine.Entities.GetEntities(types...)
}

////////////////////////////////////////////////////////////////////////////////
// EntityList Utility Functions
////////////////////////////////////////////////////////////////////////////////

// TODO: move to GameState to allow queries with values, ie Z=1
func (entityList EntityList) GetEntities(types ...ComponentType) EntityList {
	found := EntityList{}
	for _, entity := range entityList {
		if entity.HasComponents(types...) {
			found = append(found, entity)
		}
	}
	return found
}

func (entityList EntityList) GetEntitiesWithFilter(fn func(*Entity) bool) EntityList {
	result := []*Entity{}

	for _, entity := range entityList {
		if fn(entity) {
			result = append(result, entity)
		}
	}

	return result
}

// Only one Entity expected, nil if not
func (entityList EntityList) GetEntity(search_entity *Entity) *Entity {
	for _, entity := range entityList {
		if entity.Id == search_entity.Id {
			return entity
		}
	}

	return nil
}

func (entityList *EntityList) RemoveEntity(entity *Entity) EntityList {
	old := *entityList
	for i, e := range old {
		if e.Id == entity.Id {
			return append(old[:i], old[i+1:]...)
		}
	}
	return old
}

func (entityList EntityList) Concat(other EntityList) EntityList {
	return append(entityList, other...)
}

////////////////////////////////////////////////////////////////////////////////
// System Utility Functions
////////////////////////////////////////////////////////////////////////////////

// AddSystem adds the given System to the World, sorted by priority.
func (w *Engine) AddSystem(system System) {
	if initializer, ok := system.(Initializer); ok {
		initializer.New(w)
	}

	w.systems = append(w.systems, system)
	w.SortSystems()
}

// Systems returns the list of Systems managed by the World.
func (e *Engine) Systems() []System {
	return e.systems
}

// Update updates each System managed by the World. It is invoked by the engine
// once every frame, with dt being the duration since the previous update.
func (engine *Engine) Update(dt float32) {
	for _, system := range engine.Systems() {

		var entities []*Entity
		if get, ok := system.(SystemGetEntities); ok {
			entities = get.GetEntities()
		} else {
			cmps := system.(SystemComponents).Components()
			entities = engine.Entities.GetEntities(cmps...)
		}

		system.Update(entities, dt)
	}
}

// SortSystems sorts the systems in the world.
func (w *Engine) SortSystems() {
	sort.Sort(w.systems)
}
