package main

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

var (
	ColorBg,
	ColorBgDark,
	ColorBgLOS,

	ColorFg,
	ColorFgObject,
	ColorFgDark,
	ColorFgLOS,
	ColorFgLOSLight,
	ColorFgPlace,
	ColorFgPlayer gruid.Color
)

func init() {
	ColorBg = ColorBackground
	ColorBgDark = ColorBackground
	ColorBgLOS = ColorBackgroundSecondary

	ColorFg = ColorForeground
	ColorFgDark = ColorForegroundSecondary
	ColorFgLOS = ColorForegroundEmph
	ColorFgLOSLight = ColorYellow
	ColorFgObject = ColorYellow
	ColorFgPlace = ColorMagenta
	ColorFgPlayer = ColorBlue
}

func ColorToRGBA(c gruid.Color, fg bool) color.Color {
	cl := color.RGBA{}
	opaque := uint8(255)
	switch c {
	case ColorBackgroundSecondary:
		if GameConfig.DarkLOS {
			cl = color.RGBA{7, 54, 66, opaque}
		} else {
			cl = color.RGBA{238, 232, 213, opaque}
		}
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
		if GameConfig.DarkLOS {
			cl = color.RGBA{147, 161, 161, opaque}
		} else {
			cl = color.RGBA{88, 110, 117, opaque}
		}
	case ColorForegroundSecondary:
		if GameConfig.DarkLOS {
			cl = color.RGBA{88, 110, 117, opaque}
		} else {
			cl = color.RGBA{147, 161, 161, opaque}
		}
	default:
		if GameConfig.DarkLOS {
			cl = color.RGBA{0, 43, 54, opaque}
			if fg {
				cl = color.RGBA{131, 148, 150, opaque}
			}
		} else {
			cl = color.RGBA{253, 246, 227, opaque}
			if fg {
				cl = color.RGBA{101, 123, 131, opaque}
			}
		}
	}
	return cl
}

var Only8Colors bool

const (
	AttrInMap gruid.AttrMask = 1 + iota
	AttrReverse
)
