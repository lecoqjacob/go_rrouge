package ecs

type ComponentType string

type Component interface {
	ComponentType() ComponentType
}

type OnAddComponent interface {
	OnAdd(engine *Engine, entity *Entity)
}

type OnRemoveComponent interface {
	OnRemove(engine *Engine, entity *Entity)
}
