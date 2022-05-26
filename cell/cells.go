package cell

import (
	"github.com/anaseto/gruid/rl"
)

const (
	WallCell rl.Cell = iota
	FloorCell
	ExploredCell rl.Cell = 0b10000000
)

func Cell(c rl.Cell) rl.Cell {
	return c &^ ExploredCell
}

func Explored(c rl.Cell) bool {
	return c&ExploredCell != 0
}

// Rune returns the character rune representing a given terrain.
func Rune(c rl.Cell) (r rune) {
	switch c {
	case WallCell:
		r = '#'
	case FloorCell:
		r = '.'
	}

	return r
}
