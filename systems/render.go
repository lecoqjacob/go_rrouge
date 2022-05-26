package systems

import (
	"github.com/anaseto/gruid"
	"github.com/lecoqjacob/rrouge/cell"
	"github.com/lecoqjacob/rrouge/constants"
	"github.com/lecoqjacob/rrouge/dungeon"
	"github.com/lecoqjacob/rrouge/ecs"
	"github.com/lecoqjacob/rrouge/gamestate"
	"github.com/lecoqjacob/rrouge/palette"
)

func Render(engine *ecs.Engine, gs *gamestate.GameState, mgd gruid.Grid) {
	mgd.Fill(gruid.Cell{Rune: ' '})

	renderMap(gs, gs.Dgen, mgd)
	renderEntities(engine, gs, gs.Dgen, mgd)
}

func renderMap(gs *gamestate.GameState, dgen *dungeon.Dungeon, mgd gruid.Grid) {
	// We draw the map tiles.
	it := dgen.Grid.Iterator()
	for it.Next() {
		if !cell.Explored(it.Cell()) {
			continue
		}

		c := gruid.Cell{Rune: cell.Rune(cell.Cell(it.Cell()))}

		if gs.Player.InFOV(it.P()) {
			c.Style.Bg = palette.ColorFOV
		}

		mgd.Set(it.P(), c)
	}
}

func renderEntities(engine *ecs.Engine, gs *gamestate.GameState, dgen *dungeon.Dungeon, mgd gruid.Grid) {
	entities := engine.Entities.GetEntities(constants.Position, constants.Style)

	for _, entity := range entities {
		p := entity.GetPosition().Point

		if !cell.Explored(dgen.Cell(p)) || !gs.Player.InFOV(p) {
			continue
		}

		s := entity.GetStyle()
		c := mgd.At(p)

		c.Rune = s.Rune
		c.Style.Fg = s.Color

		// NOTE: We retrieved current cell at e.Pos() to preserve
		// background (in FOV or not).
		mgd.Set(p, c)
	}
}
