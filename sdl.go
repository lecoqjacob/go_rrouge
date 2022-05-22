//go:build sdl
// +build sdl

package main

import (
	"log"

	"github.com/anaseto/gruid"
	"github.com/anaseto/gruid-sdl"
)

var driver gruid.Driver

func initDriver(fullscreen bool) {
	fullscreen = fullscreen

	t, err := GetTileDrawer()
	if err != nil {
		log.Fatal(err)
	}

	dr := sdl.NewDriver(sdl.Config{
		TileManager: t,
		Fullscreen:  fullscreen,
		WindowTitle: "rrouge",
	})

	//dr.SetScale(2.0, 2.0)
	dr.PreventQuit()
	driver = dr
}

func clearCache() {
	dr := driver.(*sdl.Driver)
	dr.ClearCache()
}
