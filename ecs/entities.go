package ecs

import (
	"github.com/anaseto/gruid"
	"github.com/anaseto/gruid/paths"
	"github.com/anaseto/gruid/rl"
	"github.com/lecoqjacob/rrouge/constants"
	"github.com/lecoqjacob/rrouge/palette"

	. "github.com/lecoqjacob/rrouge/constants"
)

func NewPlayer(entity *Entity) *Entity {
	entity.AddComponent(PlayerComponent{})
	entity.AddComponent(DescriptionComponent{Name: "Player"})
	entity.AddComponent(StyleComponent{Rune: '@', Color: palette.ColorPlayer})
	// entity.AddComponent(Layer400Component{})
	// entity.AddComponent(StatsComponent{
	// 	StatsValues: &StatsValues{},
	// })
	return entity
}

func (entity *Entity) ApplyPosition(p gruid.Point) *Entity {
	entity.AddComponent(PositionComponent{p})
	return entity
}

func (entity *Entity) ApplyFOV(radius int) *Entity {
	entity.AddComponent(FovComponent{
		Dirty:  true,
		Radius: radius,
		View:   rl.NewFOV(gruid.NewRange(-MaxLOS, -MaxLOS, MaxLOS+1, MaxLOS+1)),
	})
	return entity
}

func (entity *Entity) InFOV(p gruid.Point) bool {
	pp := entity.GetPosition().Point
	fov := entity.GetFOV()

	return fov.View.Visible(p) &&
		paths.DistanceManhattan(pp, p) <= constants.MaxLOS
}
