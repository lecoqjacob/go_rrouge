package world

import (
	"github.com/anaseto/gruid"
	"github.com/anaseto/gruid/paths"
	"github.com/lecoqjacob/rrouge/rrouge_game/dungeon"
)

// aiPath implements the paths.Astar interface for use in AI pathfinding.
type AiPath struct {
	Dgen *dungeon.Dungeon
	NB   paths.Neighbors
}

// Neighbors returns the list of walkable neighbors of q in the map using 4-way
// movement along cardinal directions.
func (aip *AiPath) Neighbors(q gruid.Point) []gruid.Point {
	return aip.NB.Cardinal(q,
		func(r gruid.Point) bool {
			return aip.Dgen.Walkable(r)
		})
}

// Cost implements paths.Astar.Cost.
func (aip *AiPath) Cost(p, q gruid.Point) int {
	if !aip.Dgen.NoBlockingEntityAt(q) {
		// Extra cost for blocked positions: this encourages the
		// pathfinding algorithm to take another path to reach the
		// player.
		return 8
	}
	return 1
}

// Estimation implements paths.Astar.Estimation. For 4-way movement, we use the
// Manhattan distance.
func (aip *AiPath) Estimation(p, q gruid.Point) int {
	return paths.DistanceManhattan(p, q)
}
