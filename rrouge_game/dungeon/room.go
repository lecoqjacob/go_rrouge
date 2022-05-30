package dungeon

import (
	"github.com/anaseto/gruid"
)

type Rect struct {
	gruid.Range
}

// type Room struct {
// 	x1, x2, y1, y2 int
// }

func NewRoom(x, y, w, h int) Rect {
	// return Room{x1: x, y1: y, x2: x + w, y2: y + h}
	return Rect{gruid.NewRange(x, y, x+w, y+h)}
}

func (r *Rect) Intersect(other Rect) bool {
	return r.Overlaps(other.Range)
	// return r.x1 <= other.x2 && r.x2 >= other.x1 && r.y1 <= other.y2 && r.y2 >= other.y1
}

func (r *Rect) Center() gruid.Point {
	// return gruid.Point{X: (r.x1 + r.x2) / 2, Y: (r.y1 + r.y2) / 2}
	x1, y1 := r.Min.X, r.Min.Y
	x2, y2 := r.Max.X, r.Max.Y
	return gruid.Point{X: (x1 + x2) / 2, Y: (y1 + y2) / 2}
}

func (r *Rect) Width() int {
	return r.Range.Size().X
}

func (r *Rect) Height() int {
	return r.Range.Size().Y
}
