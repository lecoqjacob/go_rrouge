package dungeon

import (
	"math/rand"
	"time"

	"github.com/anaseto/gruid"
	"github.com/anaseto/gruid/paths"
	"github.com/anaseto/gruid/rl"

	"github.com/lecoqjacob/rrouge/ecs"
	"github.com/lecoqjacob/rrouge/rrouge_game/random"
	"github.com/lecoqjacob/rrouge/rrouge_game/utils"
)

const (
	DungeonHeight = 21
	DungeonWidth  = 80
	DungeonNCells = DungeonWidth * DungeonHeight
)

const (
	WallCell rl.Cell = iota
	FloorCell
)

// Dungeon represents the rectangular map of the game's level.
type Dungeon struct {
	Grid  rl.Grid
	Rooms []Rect
	Rand  *random.Rand     // random number generator
	PR    *paths.PathRange // path finding in the grid range

	Tile_content []*ecs.Entity
	Explored     map[gruid.Point]bool // explored cells
	Blocked      map[gruid.Point]bool // explored cells

}

func NewDungeon() *Dungeon {
	grid := rl.NewGrid(DungeonWidth, DungeonHeight)

	dgen := &Dungeon{
		Grid: grid,
		Rand: random.New(rand.NewSource(time.Now().UnixNano())),
		PR:   paths.NewPathRange(grid.Bounds()),

		Explored: make(map[gruid.Point]bool),
		Blocked:  make(map[gruid.Point]bool),
	}

	return dgen.Generate()
}

// func (dgen *Dungeon) IDX_P(i int) gruid.Point {
// 	return gruid.Point{X: i % DungeonWidth, Y: i / DungeonWidth}
// }

// func (dgen *Dungeon) xy_idx(x, y int) int {
// 	return y*dgen.width() + x
// }

//////////////////////////////////
// Private Fns
//////////////////////////////////

func (dgen *Dungeon) apply_room_to_map(room Rect) {
	// gd := dgen.Grid.Slice(gruid.NewRange(room.x1, room.y1, room.x2, room.y2))
	gd := dgen.Grid.Slice(room.Range)
	gd.Fill(rl.Cell(FloorCell))
}

func (dgen *Dungeon) apply_horizontal_tunnel(x1, x2, y int) {
	lo := utils.Min(x1, x2)
	hi := utils.Max(x1, x2)

	dgen.Grid.Slice(gruid.Range{
		Min: gruid.Point{X: lo, Y: y},
		Max: gruid.Point{X: hi, Y: y},
	}).Fill(rl.Cell(FloorCell))
}

func (dgen *Dungeon) apply_vertical_tunnel(y1, y2, x int) {
	lo := utils.Min(y1, y2)
	hi := utils.Max(y1, y2) + 1

	dgen.Grid.Slice(gruid.Range{
		Min: gruid.Point{X: x - 1, Y: lo},
		Max: gruid.Point{X: x, Y: hi},
	}).Fill(rl.Cell(FloorCell))
}

func (dgen *Dungeon) width() int {
	return dgen.Grid.Size().X
}

func (dgen *Dungeon) height() int {
	return dgen.Grid.Size().Y
}

//////////////////////////////////
// Public Fns
//////////////////////////////////

func (d *Dungeon) Cell(p gruid.Point) rl.Cell {
	return d.Grid.At(p)
}

func (dgen *Dungeon) Rune(c rl.Cell) (r rune) {
	switch c {
	case WallCell:
		r = '#'
	case FloorCell:
		r = '.'
	}

	return r
}

// Walkable returns true if at the given position is within map bounds and
// there is a floor tile
func (dgen *Dungeon) Walkable(p gruid.Point) bool {
	return dgen.Grid.Contains(p) && dgen.Grid.At(p) == FloorCell
}

// Checks if position blocks FOV
func (dgen *Dungeon) BlocksSSCLOS(p gruid.Point) bool {
	return dgen.Cell(p) != WallCell
}

// RandomFloor returns a random floor cell in the map. It assumes that such a
// floor cell exists (otherwise the function does not end).
func (dgen *Dungeon) RandomFloor() gruid.Point {
	size := dgen.Grid.Size()
	for {
		freep := gruid.Point{X: dgen.Rand.Intn(size.X), Y: dgen.Rand.Intn(size.Y)}
		if dgen.Grid.At(freep) == FloorCell {
			return freep
		}
	}
}

// RandomFloor returns a random floor cell in the map. It assumes that such a
// floor cell exists (otherwise the function does not end).
func (dgen *Dungeon) Populate_Blocked() {
	dgen.Grid.Iter(func(p gruid.Point, c rl.Cell) {
		if c == WallCell {
			dgen.Blocked[p] = true
		}
	})
}

func (dgen *Dungeon) NoBlockingEntityAt(p gruid.Point) bool {
	return !dgen.Blocked[p]
}

// NewMap returns a new map with given size.
func (dgen *Dungeon) Generate() *Dungeon {
	dgen.Grid.Fill(rl.Cell(WallCell))

	const MAX_ROOMS int = 30
	const MIN_SIZE int = 6
	const MAX_SIZE int = 10

	for range [MAX_ROOMS]int{} {
		w := dgen.Rand.Range(MIN_SIZE, MAX_SIZE)
		h := dgen.Rand.Range(MIN_SIZE, MAX_SIZE)

		// TODO: Come back and evaluate the need for -1 bounds.
		// I removed it due to keeping the rooms inside the entire grid container
		x := dgen.Rand.Roll_Dice(1, dgen.width()-w-1)  //- 1
		y := dgen.Rand.Roll_Dice(1, dgen.height()-h-1) //- 1

		new_room := NewRoom(x, y, w, h)
		ok := true

		for _, or := range dgen.Rooms {
			if new_room.Intersect(or) {
				ok = false
				break
			}
		}

		if ok {
			dgen.apply_room_to_map(new_room)

			if len(dgen.Rooms) > 0 {
				new := new_room.Center()
				prev := dgen.Rooms[len(dgen.Rooms)-1].Center()

				if dgen.Rand.Range(0, 2) == 1 {
					dgen.apply_horizontal_tunnel(prev.X, new.X, prev.Y)
					dgen.apply_vertical_tunnel(prev.Y, new.Y, new.X)
				} else {
					dgen.apply_vertical_tunnel(prev.Y, new.Y, prev.X)
					dgen.apply_horizontal_tunnel(prev.X, new.X, new.Y)
				}
			}

			dgen.Rooms = append(dgen.Rooms, new_room)
		}
	}

	return dgen
}

// // path implements the paths.Pather interface and is used to provide pathing
// // information in map generation.
// type path struct {
// 	m  *Dungeon
// 	nb paths.Neighbors
// }

// // Neighbors returns the list of walkable neighbors of q in the map using 4-way
// // movement along cardinal directions.
// func (p *path) Neighbors(q gruid.Point) []gruid.Point {
// 	return p.nb.Cardinal(q,
// 		func(r gruid.Point) bool { return p.dgen.Walkable(r) })
// }
