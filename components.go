package main

import (
	"github.com/anaseto/gruid"
	ecs "github.com/marioolofo/go-gameengine-ecs"
)

// Component IDs
const (
	PointID ecs.ID = iota
	RenderableID
	PlayerID
	FOVID
)

// Filter IDs
const (
	RenderableFilterID = iota
	PlayerFilterID
)

////////////////////////////////////////////////////////////////////////////////
/// Components
////////////////////////////////////////////////////////////////////////////////

type Renderable struct {
	Rune  rune
	Color gruid.Color
}

type Player struct{}
