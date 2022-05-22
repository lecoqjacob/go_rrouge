package main

import (
	"github.com/anaseto/gruid"
	"github.com/anaseto/gruid/rl"
	ecs "github.com/marioolofo/go-gameengine-ecs"
)

func (es *ECS) SpawnPlayer(p gruid.Point) ecs.ID {
	w := es.World

	ent := es.AddEntity(p, Renderable{Rune: '@', Color: ColorFgPlayer}, PlayerID, FOVID)

	// FOV
	pfov := (*rl.FOV)(w.Component(ent, FOVID))
	*pfov = *rl.NewFOV(gruid.NewRange(-maxLOS, -maxLOS, maxLOS+1, maxLOS+1))

	return ent
}
