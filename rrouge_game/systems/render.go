package systems

import (
	"github.com/anaseto/gruid"
	"github.com/anaseto/gruid/rl"
	"github.com/mitchellh/colorstring"

	"github.com/lecoqjacob/rrouge/ecs"
	"github.com/lecoqjacob/rrouge/rrouge_game/dungeon"
	"github.com/lecoqjacob/rrouge/rrouge_game/game"
	"github.com/lecoqjacob/rrouge/rrouge_game/palette"
	"github.com/lecoqjacob/rrouge/rrouge_game/world"
)

type RenderSystem struct {
	G *game.Game
}

func (rs *RenderSystem) Components() []ecs.ComponentType {
	return []ecs.ComponentType{world.ComponentTypes.Position, world.ComponentTypes.Style}
}

// Dummy tag
func (rs *RenderSystem) Remove(basic *ecs.Entity) {
}

func (*RenderSystem) Priority() int { return world.Priority_Movement }

var grid gruid.Grid

func (rs *RenderSystem) Update(entities ecs.EntityList, dt float32) {
	grid.Fill(gruid.Cell{Rune: ' '})

	renderMap(rs.G, grid)
	renderEntities(rs.G, grid)
}

func GetRenderables(mdg gruid.Grid) gruid.Grid {
	grid = mdg
	return grid
}

func RenderDebugMap(g *game.Game, mgd gruid.Grid) {
	gd := mgd.Slice(mgd.Range())
	pp := g.World.GetPosition(g.World.Player)

	g.Dgen.Grid.Iter(func(p gruid.Point, c rl.Cell) {
		ent, found_ai := g.World.PosCache.GetOneByCoordAndComponents(p, world.ComponentTypes.AI)

		switch {
		case p == pp:
			colorstring.Printf("[yellow]%s", "@")
		case found_ai:
			colorstring.Printf("[red]%c", g.World.GetStyle(ent).Rune)
		case g.Dgen.Cell(p) == dungeon.WallCell:
			colorstring.Printf("[green]%s", "#")
		case g.Dgen.Cell(p) == dungeon.FloorCell:
			colorstring.Printf("[white]%s", ".")
		}

		if p.X == gd.Rg.Max.X-1 {
			colorstring.Println("")
		}
	})
}

////////////////////////////////////////////////////////////////////////////////
/// Rendering Utilities
////////////////////////////////////////////////////////////////////////////////

func renderMap(g *game.Game, mgd gruid.Grid) {
	// We draw the map tiles.
	it := g.Dgen.Grid.Iterator()
	for it.Next() {
		if !g.Dgen.Explored[it.P()] {
			continue
		}

		c := gruid.Cell{Rune: g.Dgen.Rune(it.Cell())}
		if g.World.InFOV(g.World.Player, it.P()) {
			c.Style.Bg = palette.ColorFOV
		}

		mgd.Set(it.P(), c)
	}
}

func renderEntities(g *game.Game, mgd gruid.Grid) {
	entities := g.World.GetEntities(world.ComponentTypes.Position, world.ComponentTypes.Style)

	for _, entity := range entities {
		p := g.World.GetPosition(entity)

		if !g.Dgen.Explored[p] || !g.World.InFOV(g.World.Player, p) {
			continue
		}

		s := g.World.GetStyle(entity)
		c := mgd.At(p)

		c.Rune = s.Rune
		c.Style.Fg = s.Color

		// NOTE: We retrieved current cell at e.Pos() to preserve
		// background (in FOV or not).
		mgd.Set(p, c)
	}
}
