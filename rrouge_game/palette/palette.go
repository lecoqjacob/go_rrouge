package palette

import (
	"image/color"

	"github.com/anaseto/gruid"
)

// Thoses are the colors of the main palette. They are given 16-palette color
// numbers compatible with terminals, though they are then mapped to more
// precise colors depending on options and the driver. Dark colorscheme is
// assumed by default, but it can be changed in configuration.
const (
	ColorBackground          gruid.Color = gruid.ColorDefault // background
	ColorBackgroundSecondary gruid.Color = 1 + 0              // black
	ColorForeground          gruid.Color = gruid.ColorDefault
	ColorForegroundSecondary gruid.Color = 1 + 7  // white
	ColorForegroundEmph      gruid.Color = 1 + 15 // bright white
	ColorYellow              gruid.Color = 1 + 3
	ColorOrange              gruid.Color = 1 + 1 // red
	ColorRed                 gruid.Color = 1 + 9 // bright red
	ColorMagenta             gruid.Color = 1 + 5
	ColorViolet              gruid.Color = 1 + 12 // bright blue
	ColorBlue                gruid.Color = 1 + 4
	ColorCyan                gruid.Color = 1 + 6
	ColorGreen               gruid.Color = 1 + 2
)

// Color definitions. For now, we use a special color for FOV. We start from 1,
// because 0 is gruid.ColorDefault, which we use for default foreground and
// background.
var (
	ColorFg,
	ColorFgDark,
	ColorFgLOS,

	ColorBg,
	ColorBgDark,
	ColorBgLOS,

	ColorPlayer,
	ColorMonster,
	ColorFOV gruid.Color
	// ColorFOV,
	// ColorLogPlayerAttack,
	// ColorLogItemUse,
	// ColorLogMonsterAttack,
	// ColorLogSpecial,
	// ColorStatusHealthy,
	// ColorStatusWounded,
	// ColorConsumable gruid.Color
)

func init() {
	ColorFg = ColorForeground
	ColorFgDark = ColorForegroundSecondary
	ColorFgLOS = ColorForegroundEmph

	ColorBg = ColorBackground
	ColorBgDark = ColorBackground
	ColorBgLOS = ColorBackgroundSecondary

	ColorPlayer = ColorYellow
	ColorMonster = ColorRed
	ColorFOV = ColorForegroundEmph
}

func ColorToRGBA(c gruid.Color, fg bool) color.Color {
	cl := color.RGBA{}
	opaque := uint8(255)

	switch c {
	case ColorRed:
		cl = color.RGBA{220, 50, 47, opaque}
	case ColorGreen:
		cl = color.RGBA{133, 153, 0, opaque}
	case ColorYellow:
		cl = color.RGBA{181, 137, 0, opaque}
	case ColorBlue:
		cl = color.RGBA{38, 139, 210, opaque}
	case ColorMagenta:
		cl = color.RGBA{211, 54, 130, opaque}
	case ColorCyan:
		cl = color.RGBA{42, 161, 152, opaque}
	case ColorOrange:
		cl = color.RGBA{203, 75, 22, opaque}
	case ColorViolet:
		cl = color.RGBA{108, 113, 196, opaque}
	case ColorForegroundEmph:
		cl = color.RGBA{88, 110, 117, opaque}
	case ColorForegroundSecondary:
		cl = color.RGBA{147, 161, 161, opaque}
	default:
		cl = color.RGBA{253, 246, 227, opaque}
		if fg {
			cl = color.RGBA{101, 123, 131, opaque}
		}
	}

	return cl
}

var Only8Colors bool

const (
	AttrReverse = 1 << iota
)
