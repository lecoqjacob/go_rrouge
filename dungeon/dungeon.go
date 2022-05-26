package dungeon

import (
	"math/rand"
	"time"

	"github.com/anaseto/gruid"
	"github.com/anaseto/gruid/rl"
	"github.com/lecoqjacob/rrouge/cell"
	"github.com/lecoqjacob/rrouge/random"
)

const (
	DungeonHeight = 21
	DungeonWidth  = 80
	DungeonNCells = DungeonWidth * DungeonHeight
)

// Dungeon represents the rectangular map of the game's level.
type Dungeon struct {
	Grid rl.Grid
	Rand *random.Rand // random number generator

	rooms []Room

	width, height int
}

//////////////////////////////////
// Utility Functions
//////////////////////////////////

func (d *Dungeon) Cell(p gruid.Point) rl.Cell {
	return d.Grid.At(p)
}

func (dgen *Dungeon) SetCell(p gruid.Point, c rl.Cell) {
	oc := dgen.Cell(p)
	dgen.Grid.Set(p, rl.Cell(c|oc&cell.ExploredCell))
}

func (dgen *Dungeon) SetExplored(p gruid.Point) {
	oc := dgen.Grid.At(p)
	dgen.Grid.Set(p, rl.Cell(oc|cell.ExploredCell))
}

// func idxtopos(i int) gruid.Point {
// 	return gruid.Point{X: i % DungeonWidth, Y: i / DungeonWidth}
// }

func (m *Dungeon) xy_idx(x, y int) int {
	return y*m.width + x
}

func (m Dungeon) Player_Start() gruid.Point {
	return m.rooms[0].Center()
}

//////////////////////////////////
/// Dungeon
//////////////////////////////////

func (m *Dungeon) apply_room_to_map(room *Room) {
	// r := gruid.NewRange(room.x1, room.y1, room.x2, room.y2)
	gd := m.Grid.Slice(*room.Range)
	gd.Fill(rl.Cell(cell.FloorCell))
}

func (m *Dungeon) apply_horizontal_tunnel(x1, x2, y int) {
	// lo := Min(x1, x2)
	// hi := Max(x1, x2)
	// s := arr_range(lo, hi)

	// for _, x := range s {
	// 	idx := m.xy_idx(x, y)
	// 	if idx > 0 && idx < m.width*m.height {
	// 		p := idxtopos(idx)
	// 		m.Grid.Set(p, rl.Cell(FloorCell))
	// 	}
	// }
}

func (m *Dungeon) apply_vertical_tunnel(y1, y2, x int) {
	// lo := min(y1, y1)
	// hi := max(y1, y2)
	// s := arr_range(lo, hi)

	// for _, y := range s {
	// 	idx := m.xy_idx(x, y)
	// 	if idx > 0 && idx < m.width*m.height {
	// 		p := idxtopos(idx)
	// 		m.Grid.Set(p, rl.Cell(FloorCell))
	// 	}
	// }
}

// Walkable returns true if at the given position there is a floor tile.
func (m *Dungeon) Walkable(p gruid.Point) bool {
	if !m.Grid.Contains(p) {
		return false
	}

	return cell.Cell(m.Grid.At(p)) != cell.WallCell
}

// NewMap returns a new map with given size.
func NewDungeon() *Dungeon {
	m := &Dungeon{
		Grid:   rl.NewGrid(DungeonWidth, DungeonHeight),
		Rand:   random.New(rand.NewSource(time.Now().UnixNano())),
		width:  DungeonWidth,
		height: DungeonHeight,
	}

	m.Grid.Fill(rl.Cell(cell.WallCell))
	m.Generate()

	return m
}

// Generate fills the Grid attribute of m with a procedurally generated map.
func (m *Dungeon) Generate() {
	const MAX_ROOMS int = 30
	const MIN_SIZE int = 6
	const MAX_SIZE int = 10

	for range [MAX_ROOMS]int{} {
		w := m.Rand.Range(MIN_SIZE, MAX_SIZE)
		h := m.Rand.Range(MIN_SIZE, MAX_SIZE)

		x := m.Rand.Roll_Dice(1, m.width-w-1) - 1
		y := m.Rand.Roll_Dice(1, m.height-w-1) - 1
		// new_room := Room{Range: &gruid.Range{Min: gruid.Point{X: x, Y: y}, Max: gruid.Point{X: x + w, Y: y + h}}}
		r := gruid.NewRange(x, y, w, h)
		new_room := Room{Range: &r}
		ok := true

		for _, or := range m.rooms {
			if new_room.Intersect(or) {
				ok = false
				break
			}
		}

		if ok {
			m.apply_room_to_map(&new_room)

			if len(m.rooms) > 0 {
				new := new_room.Center()
				prev := m.rooms[len(m.rooms)-1].Center()

				if m.Rand.Range(0, 2) == 1 {
					m.apply_horizontal_tunnel(prev.X, new.Y, prev.Y)
					m.apply_vertical_tunnel(prev.Y, new.Y, new.X)
				} else {
					m.apply_vertical_tunnel(prev.Y, new.Y, prev.X)
					m.apply_horizontal_tunnel(prev.X, new.X, new.Y)
				}

			}

			m.rooms = append(m.rooms, new_room)
		}
	}
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
// 		func(r gruid.Point) bool { return p.m.Walkable(r) })
// }
