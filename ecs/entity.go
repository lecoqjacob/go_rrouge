package ecs

import (
	"strconv"
)

////////////////////////////////////////////////////////////////////////
// Entity
////////////////////////////////////////////////////////////////////////

var (
	idInc uint64
)

type Entity struct {
	id         uint64
	Components map[ComponentType]Component
	engine     *Engine
}

// Identifier is an interface for anything that implements the basic ID() uint64,
// as the BasicEntity does.  It is useful as more specific interface for an
// entity registry than just the interface{} interface
type Identifier interface {
	ID() uint64
}

// IdentifierSlice implements the sort.Interface, so you can use the
// store entites in slices, and use the P=n*log n lookup for them
type IdentifierSlice []Identifier

// ID returns the unique identifier of the entity.
func (e *Entity) ID() uint64 {
	return e.id
}

func (entity *Entity) String() string {
	var str = "Entity " + strconv.FormatUint(entity.ID(), 10) + "["
	for _, c := range entity.Components {
		str += string(c.ComponentType()) + ","
	}
	str += "]"
	return str
}

// Inserts a component onto the entity.
//If a previous component exist, it will be replaced.
func (entity *Entity) InsertComponent(component Component) *Entity {
	entity.Components[component.ComponentType()] = component

	// Call event if possible
	cmp, ok := component.(OnAddComponent)
	if ok {
		cmp.OnAdd(entity.engine, entity)
	}

	return entity
}

func (entity *Entity) InsertComponents(components ...Component) *Entity {
	for _, component := range components {
		entity.InsertComponent(component)
	}

	return entity
}

func (entity *Entity) removeComponent(componentType ComponentType) *Entity {
	component, ok := entity.Components[componentType]
	if ok {
		delete(entity.Components, componentType)

		// Call event if possible
		cmp, ok := component.(OnRemoveComponent)
		if ok {
			cmp.OnRemove(entity.engine, entity)
		}
	}

	return entity
}

func (entity *Entity) RemoveComponents(componentTypes ...ComponentType) *Entity {
	for _, componentType := range componentTypes {
		entity.removeComponent(componentType)
	}

	return entity
}

func (entity Entity) HasComponent(componentType ComponentType) bool {
	_, ok := entity.Components[componentType]
	return ok
}

func (entity Entity) HasComponents(componentTypes ...ComponentType) bool {
	// Check to see if the entity has the given components
	containsAll := true
	for i := 0; i < len(componentTypes); i++ {
		if _, ok := entity.Components[componentTypes[i]]; !ok {
			containsAll = false
			break
		}
	}
	return containsAll
}

func (entity Entity) GetComponent(componentType ComponentType) Component {
	if cmp, ok := entity.Components[componentType]; ok {
		return cmp
	} else {
		return nil
	}
}
