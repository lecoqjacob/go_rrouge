package main

import "github.com/anaseto/gruid/rl"

type cell rl.Cell

const (
	WallCell cell = iota
	GroundCell
	Explored = 0b10000000
)

func terrain(c cell) cell {
	return c &^ Explored
}

func explored(c cell) bool {
	return c&Explored != 0
}
