package main

import (
	"math/rand"

	"github.com/anaseto/gruid"
	"github.com/anaseto/gruid/rl"
)

var Version string = "v0.0.1"

// game represents information relevant the current game's state.
type game struct {
	ECS *ECS // entities present on the map
	Map *Map // the game map, made of tiles

	Log         []logEntry
	LogIndex    int
	LogNextTick int
}

// Map represents the rectangular map of the game's level.
type Map struct {
	Grid     rl.Grid
	Rand     *rand.Rand           // random number generator
	Explored map[gruid.Point]bool // explored cells
}
