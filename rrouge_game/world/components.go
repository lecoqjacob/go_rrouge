package world

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/anaseto/gruid"
	"github.com/anaseto/gruid/rl"

	"github.com/lecoqjacob/rrouge/ecs"
	"github.com/lecoqjacob/rrouge/rrouge_game/utils"
)

var ComponentTypes _componentTypes

type _componentTypes struct {
	Position     ecs.ComponentType
	Style        ecs.ComponentType
	Description  ecs.ComponentType
	Player       ecs.ComponentType
	AI           ecs.ComponentType
	FOV          ecs.ComponentType
	Move         ecs.ComponentType
	BlocksCell   ecs.ComponentType
	Stats        ecs.ComponentType
	WantsToMelee ecs.ComponentType
	SufferDamage ecs.ComponentType
}

func initComponentTypes() {
	ComponentTypes = _componentTypes{
		Position:     "position",
		Style:        "style",
		Description:  "description",
		Player:       "player",
		AI:           "ai",
		FOV:          "fov",
		Move:         "move",
		BlocksCell:   "blocks_cell",
		Stats:        "stats",
		WantsToMelee: "wants_to_melee",
		SufferDamage: "suffer_damage",
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

func (s StyleComponent) String() string {
	return fmt.Sprintf("Rune: %c / Color: %v", s.Rune, s.Color)
}

// Position
type PositionComponent struct{ gruid.Point }

func (a PositionComponent) ComponentType() ecs.ComponentType {
	return ComponentTypes.Position
}

func (a PositionComponent) GetKey() string {
	return strconv.Itoa(a.X) + "," + strconv.Itoa(a.Y)
}

// func (a PositionComponent) String() string {
// 	return a.GetKey()
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

// Stats
type StatsValues struct {
	HP      int
	Max_HP  int
	Defense int
	Power   int
	Fov     int
}

func (stats StatsValues) String() string {
	var sb strings.Builder
	sb.WriteString(utils.StatStrMax("Health", stats.HP, stats.Max_HP))
	sb.WriteString(utils.StatStr("Pow", stats.Power))
	sb.WriteString(utils.StatStr("Def", stats.Defense))
	sb.WriteString(utils.StatStr("FOV", stats.Fov))

	return strings.Trim(sb.String(), " ")
}

type StatsComponent struct {
	*StatsValues
}

func (a StatsComponent) ComponentType() ecs.ComponentType {
	return ComponentTypes.Stats
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

type WantsToMelee struct {
	Target *ecs.Entity
}

func (a WantsToMelee) ComponentType() ecs.ComponentType {
	return ComponentTypes.WantsToMelee
}

////////////////////////////////////////////////////////////////////////////////
// Utility Components
////////////////////////////////////////////////////////////////////////////////

type SufferDamage struct {
	Amount []int
}

func (sd SufferDamage) ComponentType() ecs.ComponentType {
	return ComponentTypes.SufferDamage
}
