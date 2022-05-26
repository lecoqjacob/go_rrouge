package dungeon

import (
	"github.com/anaseto/gruid"
)

type Room struct {
	*gruid.Range
}

func (r *Room) Intersect(other Room) bool {
	return r.Overlaps(*other.Range)
}

func (r *Room) Center() gruid.Point {
	x1, y1 := r.Min.X, r.Min.Y
	x2, y2 := r.Max.X, r.Max.Y

	return gruid.Point{X: (x1 + x2) / 2, Y: (y1 + y2) / 2}
}
