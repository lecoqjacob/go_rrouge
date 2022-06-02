package utils

import (
	"fmt"
	"strconv"

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

func Ternary(condition bool, ifTrue, ifFalse interface{}) interface{} {
	if condition {
		return ifTrue
	} else {
		return ifFalse
	}
}

func AppendSlice[T any](s []T, v T) []T {
	return append(s, v)
}

func Sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

////////////////////////////////////////////////////////////////////////////////
// Stats
////////////////////////////////////////////////////////////////////////////////

func AddSign(i int) string {
	if i >= 0 {
		return "+" + strconv.Itoa(i)
	}
	return strconv.Itoa(i)
}

func StatStrMax(name string, value int, max int) string {

	if value != 0 || max != 0 {
		return name + " " + AddSign(value) + "/" + AddSign(max) + " "
	}
	return ""
}

func StatStr(name string, value int) string {
	if value != 0 {
		return fmt.Sprintf("%s %d ", name, value)
	}
	return ""
}
