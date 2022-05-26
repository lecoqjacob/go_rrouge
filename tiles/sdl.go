package tiles

import (
	"log"

	"github.com/anaseto/gruid"
	sdl "github.com/anaseto/gruid-sdl"
	"github.com/lecoqjacob/rrouge/constants"
)

var Driver gruid.Driver
var IsFullscreen bool

func InitDriver(fullscreen bool) {
	IsFullscreen = fullscreen

	t, err := GetTileDrawer()
	if err != nil {
		log.Fatal(err)
	}

	dr := sdl.NewDriver(sdl.Config{
		TileManager: t,
		Fullscreen:  fullscreen,
		WindowTitle: constants.GameName,
	})

	//dr.SetScale(2.0, 2.0)
	dr.PreventQuit()
	Driver = dr
}

func ClearCache() {
	dr := Driver.(*sdl.Driver)
	dr.ClearCache()
}
