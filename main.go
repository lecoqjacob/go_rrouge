//go:build !js
// +build !js

package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/anaseto/gruid"
)

const (
	LogLines  = 2
	MapWidth  = UIWidth
	MapHeight = UIHeight - 1 - LogLines
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
		fmt.Println(Version)
		os.Exit(0)
	}

	if runtime.GOOS != "windows" {
		Xterm256Color = true
	} else {
		Xterm256Color = false
		Only8Colors = true
	}

	if *opt256colors {
		Xterm256Color = true
		Only8Colors = false
	} else if *opt16colors {
		Xterm256Color = false
		Only8Colors = false
	}

	err := initConfig()
	if err != nil {
		log.Print(err)
	}

	applyThemeConf()
	initDriver(*optFullscreen)

	RunGame(*optLogFile)
}

func RunGame(logfile string) {
	grid := gruid.NewGrid(UIWidth, UIHeight)
	m := &model{grid: grid, game: &game{}}

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

	if !Tiles && !LogGame {
		log.SetOutput(ioutil.Discard)
	}

	app := gruid.NewApp(gruid.AppConfig{
		Driver: driver,
		Model:  m,
	})

	err := app.Start(context.Background())
	if !Tiles && !LogGame {
		log.SetOutput(os.Stderr)
	}

	if err != nil {
		log.Fatal(err)
	}
}

func subSig(ctx context.Context, msgs chan<- gruid.Msg) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	defer signal.Stop(sig)

	select {
	case <-ctx.Done():
	case <-sig:
		msgs <- gruid.MsgQuit{}
	}
}
