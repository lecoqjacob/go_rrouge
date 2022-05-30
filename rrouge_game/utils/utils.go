package utils

import (
	"github.com/anaseto/gruid"
)

func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func Range(lo, hi int) []int {
	s := make([]int, hi-lo+1)
	for i := range s {
		s[i] = i + lo
	}
	return s
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Distance(p, q gruid.Point) int {
	p = p.Sub(q)
	return Abs(p.X) + Abs(p.Y)
}

func KeyToDir(key gruid.Key) (p gruid.Point) {
	switch key {
	case gruid.KeyArrowLeft, "h":
		p = gruid.Point{X: -1, Y: 0}
	case gruid.KeyArrowDown, "j":
		p = gruid.Point{X: 0, Y: 1}
	case gruid.KeyArrowUp, "k":
		p = gruid.Point{X: 0, Y: -1}
	case gruid.KeyArrowRight, "l":
		p = gruid.Point{X: 1, Y: 0}
	}
	return p
}
