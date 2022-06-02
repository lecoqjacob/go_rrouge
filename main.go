package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"

	"github.com/anaseto/gruid"

	. "github.com/lecoqjacob/rrouge/rrouge_game/constants"
	"github.com/lecoqjacob/rrouge/rrouge_game/game"
	"github.com/lecoqjacob/rrouge/rrouge_game/gamestate"
	"github.com/lecoqjacob/rrouge/rrouge_game/palette"
	"github.com/lecoqjacob/rrouge/rrouge_game/tiles"
)

func main() {
	optVersion := flag.Bool("v", false, "print version number")
	optLogFile := flag.String("o", "", "log to output file")

	opt16colors := new(bool)
	opt256colors := new(bool)
	optFullscreen := new(bool)

	if Terminal {
		opt16colors = flag.Bool("s", false, "use 16-color simple palette")
		opt256colors = flag.Bool("x", false, "use xterm 256-color palette (solarized approximation)")
	} else {
		optFullscreen = flag.Bool("F", false, "fullscreen")
	}
	flag.Parse()

	if *optVersion {
		fmt.Println(game.Version)
		os.Exit(0)
	}

	if runtime.GOOS != "windows" {
		Xterm256Color = true
	} else {
		Xterm256Color = false
		palette.Only8Colors = true
	}

	if *opt256colors {
		Xterm256Color = true
		palette.Only8Colors = false
	} else if *opt16colors {
		Xterm256Color = false
		palette.Only8Colors = false
	}

	tiles.InitDriver(*optFullscreen)
	RunGame(*optLogFile)
}

func RunGame(logfile string) {
	// Create a new grid with standard 80x24 size.
	gd := gruid.NewGrid(UI_WIDTH, UI_HEIGHT)

	// m := &game.Model{Grid: gd, Game: &game.Game{}}
	gs := &gamestate.GameState{Grid: gd, Game: &game.Game{}}

	if logfile != "" {
		f, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Printf("opening log file: %v", err)
		} else {
			defer f.Close()
			LogGame = true
			log.SetOutput(f)
		}
	}

	if !tiles.Tiles && !LogGame {
		log.SetOutput(ioutil.Discard)
	}

	app := gruid.NewApp(gruid.AppConfig{
		Model:  gs,
		Driver: tiles.Driver,
	})

	err := app.Start(context.Background())
	if !tiles.Tiles && !LogGame {
		log.SetOutput(os.Stderr)
	}

	if err != nil {
		log.Fatal(err)
	}
}
