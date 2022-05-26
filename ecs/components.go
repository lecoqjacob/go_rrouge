package ecs

import (
	"strconv"

	"github.com/anaseto/gruid"
	"github.com/anaseto/gruid/rl"
	"github.com/lecoqjacob/rrouge/constants"
)

type PlayerComponent struct {
	FOV *rl.FOV // player's field of view
}

func (a PlayerComponent) ComponentType() string {
	return constants.Player
}

type StyleComponent struct {
	Rune  rune
	Color gruid.Color
}

func (a StyleComponent) ComponentType() string {
	return constants.Style
}

func (entity *Entity) GetStyle() StyleComponent {
	return entity.GetComponent(constants.Style).(StyleComponent)
}

type PositionComponent struct {
	gruid.Point
}

func (a PositionComponent) ComponentType() string {
	return constants.Position
}

func (entity *Entity) GetPosition() PositionComponent {
	return entity.GetComponent(constants.Position).(PositionComponent)
}

func (a PositionComponent) GetKey() string {
	return strconv.Itoa(a.X) + "," + strconv.Itoa(a.Y)
}

func (a PositionComponent) String() string {
	return a.GetKey()
}

func (a PositionComponent) OnAdd(engine *Engine, entity *Entity) {
	engine.PosCache.Add(a.GetKey(), entity)
}

func (a PositionComponent) OnRemove(engine *Engine, entity *Entity) {
	engine.PosCache.Delete(a.GetKey(), entity)
}

type DescriptionComponent struct {
	Name string
}

func (a DescriptionComponent) ComponentType() string {
	return constants.Description
}

type FovComponent struct {
	View   *rl.FOV
	Dirty  bool
	Radius int
}

func (a FovComponent) ComponentType() string {
	return constants.FOV
}

func (entity *Entity) GetFOV() FovComponent {
	return entity.GetComponent(constants.FOV).(FovComponent)
}

type MoveComponent struct {
	To gruid.Point
}

func (a MoveComponent) ComponentType() string {
	return constants.Move
}
