package tiles

import (
	"image"
	"image/color"

	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/opentype"

	"github.com/anaseto/gruid"
	"github.com/anaseto/gruid/tiles"

	"github.com/lecoqjacob/rrouge/rrouge_game/palette"
)

const Tiles = true

// TileDrawer implements TileManager from the gruid-sdl module. It is used to
// provide a mapping from virtual grid cells to images using the tiles package.
// In this tutorial, we just draw a font with a given foreground and
// background, but it would be possible to make a tiles version with custom
// drawings for cells.
type TileDrawer struct {
	drawer *tiles.Drawer
}

// TileSize implements TileManager.TileSize. It returns the tile size, in
// pixels. In this tutorial, it corresponds to the size of a character with the
// font we use.
func (t *TileDrawer) TileSize() gruid.Point {
	return t.drawer.Size()
}

// GetImage implements TileManager.GetImage.
func (t *TileDrawer) GetImage(c gruid.Cell) image.Image {
	// We use some colors from https://github.com/jan-warchol/selenized,
	// using the palette variant with dark backgound and light foreground.

	fg := image.NewUniform(palette.ColorToRGBA(c.Style.Fg, true))
	bg := image.NewUniform(color.RGBA{0x10, 0x3c, 0x48, 255})
	// bg := image.NewUniform(palette.ColorToRGBA(c.Style.Bg, false))

	// We define non default-colors (for FOV, ...).
	switch c.Style.Bg {
	case palette.ColorFOV:
		bg = image.NewUniform(color.RGBA{0x18, 0x49, 0x56, 255})
		// bg = image.NewUniform(palette.ColorToRGBA(c.Style.Bg, false))
	}

	// switch c.Style.Fg {
	// case palette.ColorPlayer:
	// 	fg = image.NewUniform(color.RGBA{0x46, 0x95, 0xf7, 255})
	// case palette.ColorMonster:
	// 	fg = image.NewUniform(color.RGBA{0xfa, 0x57, 0x50, 255})
	// }

	if c.Style.Attrs&palette.AttrReverse != 0 {
		fg, bg = bg, fg
	}

	return t.drawer.Draw(c.Rune, fg, bg)
}

// GetTileDrawer returns a TileDrawer that implements TileManager for the sdl
// driver, or an error if there were problems setting up the font face.
func GetTileDrawer() (*TileDrawer, error) {
	t := &TileDrawer{}
	var err error
	// We get a monospace font TTF.
	font, err := opentype.Parse(gomono.TTF)
	if err != nil {
		return nil, err
	}
	// We retrieve a font face.
	face, err := opentype.NewFace(font, &opentype.FaceOptions{
		Size: 24,
		DPI:  72,
	})
	if err != nil {
		return nil, err
	}
	// We create a new drawer for tiles using the previous face. Note that
	// if more than one face is wanted (such as an italic or bold variant),
	// you would have to create drawers for thoses faces too, and then use
	// the relevant one accordingly in the GetImage method.
	t.drawer, err = tiles.NewDrawer(face)
	if err != nil {
		return nil, err
	}
	return t, nil
}
