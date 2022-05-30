package world

import (
	"strconv"

	"github.com/anaseto/gruid"
	"github.com/lecoqjacob/rrouge/ecs"
)

/**
 * Position cache
 */

type EntitySet map[*ecs.Entity]bool

type PositionCache struct {
	Entities map[string]EntitySet
}

func createKey(p gruid.Point) string {
	return strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y)
}

func (c *PositionCache) Add(key string, value *ecs.Entity) {
	_, ok := c.Entities[key]
	if !ok {
		c.Entities[key] = make(EntitySet)
	}
	c.Entities[key][value] = true
}

func (c *PositionCache) Delete(key string, value *ecs.Entity) {
	set, ok := c.Entities[key]
	if ok {
		delete(set, value)
		if len(set) == 0 {
			delete(c.Entities, key)
		}
	}
}

func (c *PositionCache) GetByCoord(p gruid.Point) (ecs.EntityList, bool) {
	key := createKey(p)
	return c.Get(key)
}

func (c *PositionCache) GetByCoordAndComponents(p gruid.Point, cmpTypes ...ecs.ComponentType) (ecs.EntityList, bool) {
	key := createKey(p)
	entityList, ok := c.Get(key)
	if ok {
		return entityList.GetEntities(cmpTypes...), true
	}
	return ecs.EntityList{}, false
}

func (c *PositionCache) GetOneByCoordAndComponents(p gruid.Point, cmpTypes ...ecs.ComponentType) (*ecs.Entity, bool) {
	key := createKey(p)
	entityList, ok := c.Get(key)
	if ok {
		found := entityList.GetEntities(cmpTypes...)
		if len(found) == 1 {
			return found[0], true
		}
	}

	return &ecs.Entity{}, false
}

func (c *PositionCache) Get(key string) (ecs.EntityList, bool) {
	set, ok := c.Entities[key]
	entities := make([]*ecs.Entity, 0, len(set))
	if ok {
		for entity := range set {
			entities = append(entities, entity)
		}

	}
	return entities, ok
}
