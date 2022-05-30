package world

import (
	"strconv"

	"github.com/anaseto/gruid"
	"github.com/anaseto/gruid/rl"

	"github.com/lecoqjacob/rrouge/ecs"
)

var ComponentTypes _componentTypes

type _componentTypes struct {
	Position    ecs.ComponentType
	Style       ecs.ComponentType
	Description ecs.ComponentType
	Player      ecs.ComponentType
	AI          ecs.ComponentType
	FOV         ecs.ComponentType
	Move        ecs.ComponentType
	BlocksCell  ecs.ComponentType
}

func initComponentTypes() {
	ComponentTypes = _componentTypes{
		Position:    "position",
		Style:       "style",
		Description: "description",
		Player:      "player",
		AI:          "ai",
		FOV:         "fov",
		Move:        "move",
		BlocksCell:  "blocks_cell",
	}
}

////////////////////////////////////////////////////////////////////////////////
// Base Entity Components
////////////////////////////////////////////////////////////////////////////////

// Style
type StyleComponent struct {
	Rune  rune
	Color gruid.Color
}

func (a StyleComponent) ComponentType() ecs.ComponentType {
	return ComponentTypes.Style
}

// Position
type PositionComponent struct{ gruid.Point }

func (a PositionComponent) ComponentType() ecs.ComponentType {
	return ComponentTypes.Position
}

func (a PositionComponent) GetKey() string {
	return strconv.Itoa(a.X) + "," + strconv.Itoa(a.Y)
}

func (a PositionComponent) String() string {
	return a.GetKey()
}

// func (a PositionComponent) OnAdd(engine *ecs.Engine, entity *ecs.Entity) {
// 	engine.PosCache.Add(a.GetKey(), entity)
// }

// func (a PositionComponent) OnRemove(engine *ecs.Engine, entity *ecs.Entity) {
// 	engine.PosCache.Delete(a.GetKey(), entity)
// }

// Description
type DescriptionComponent struct {
	Name string
}

func (a DescriptionComponent) ComponentType() ecs.ComponentType {
	return ComponentTypes.Description
}

// FOV
type FovComponent struct {
	View   *rl.FOV
	Dirty  bool
	Radius int
}

func (a FovComponent) ComponentType() ecs.ComponentType {
	return ComponentTypes.FOV
}

// Blocks Cell
type BlocksCellComponent struct{}

func (a BlocksCellComponent) ComponentType() ecs.ComponentType {
	return ComponentTypes.BlocksCell
}

////////////////////////////////////////////////////////////////////////////////
// Entity Tags
////////////////////////////////////////////////////////////////////////////////

type PlayerComponent struct{}

func (a PlayerComponent) ComponentType() ecs.ComponentType {
	return ComponentTypes.Player
}

type AIComponent struct {
	AiPath *AiPath
	Path   []gruid.Point // path to destination
}

func (a AIComponent) ComponentType() ecs.ComponentType {
	return ComponentTypes.AI
}

////////////////////////////////////////////////////////////////////////////////
// Intent Components
////////////////////////////////////////////////////////////////////////////////
type MoveComponent struct {
	To gruid.Point
}

func (a MoveComponent) ComponentType() ecs.ComponentType {
	return ComponentTypes.Move
}
