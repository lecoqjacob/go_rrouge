package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/anaseto/gruid"
	"github.com/anaseto/gruid/ui"
)

const (
	UIWidth  = 80
	UIHeight = 24
)

var (
	Xterm256Color = false
	Terminal      = false
	LogGame       = false
)

var CustomKeys bool
var GameConfig config

type mode int

const (
	modeNormal mode = iota
	modeWelcome
	modeQuit
	modeQuitConfirmation
)

type model struct {
	grid gruid.Grid // drawing grid
	game *game      // game state
	mode mode

	keysNormal map[gruid.Key]actionType
	keysTarget map[gruid.Key]actionType

	log *ui.Label
}

func initConfig() error {
	GameConfig.DarkLOS = true
	GameConfig.Version = Version
	GameConfig.Tiles = true

	load, err := LoadConfig()
	if err != nil {
		err = fmt.Errorf("error loading config: %v", err)
		saverr := SaveConfig()
		if saverr != nil {
			log.Printf("Error resetting badly loaded config: %v", err)
		}
		return err
	}

	if load {
		CustomKeys = true
	}

	return err
}

func (md *model) initKeys() {
	md.keysNormal = map[gruid.Key]actionType{
		gruid.KeyArrowLeft:  ActionW,
		gruid.KeyArrowDown:  ActionS,
		gruid.KeyArrowUp:    ActionN,
		gruid.KeyArrowRight: ActionE,
		"h":                 ActionW,
		"j":                 ActionS,
		"k":                 ActionN,
		"l":                 ActionE,
		"a":                 ActionW,
		"s":                 ActionS,
		"w":                 ActionN,
		"d":                 ActionE,
		"4":                 ActionW,
		"2":                 ActionS,
		"8":                 ActionN,
		"6":                 ActionE,
		".":                 ActionWaitTurn,
		gruid.KeyEnter:      ActionWaitTurn,
		"5":                 ActionWaitTurn,
		"e":                 ActionInteract,
		"E":                 ActionInteract,
		"Q":                 ActionQuit,
		gruid.KeyEscape:     ActionEscape,
	}

	md.keysTarget = map[gruid.Key]actionType{
		gruid.KeyArrowLeft:  ActionW,
		gruid.KeyArrowDown:  ActionS,
		gruid.KeyArrowUp:    ActionN,
		gruid.KeyArrowRight: ActionE,
		"h":                 ActionW,
		"j":                 ActionS,
		"k":                 ActionN,
		"l":                 ActionE,
		"a":                 ActionW,
		"s":                 ActionS,
		"w":                 ActionN,
		"d":                 ActionE,
		"4":                 ActionW,
		"2":                 ActionS,
		"8":                 ActionN,
		"6":                 ActionE,
		gruid.KeySpace:      ActionEscape,
		gruid.KeyEscape:     ActionEscape,
		"x":                 ActionEscape,
		"X":                 ActionEscape,
		// "?":             ActionHelp,
	}
	CustomKeys = false
}

func (md *model) initWidgets() {
	md.log = ui.NewLabel(ui.StyledText{}.WithStyle(gruid.Style{}).WithMarkup('t', gruid.Style{Fg: ColorYellow}))
}

func (md *model) init() gruid.Effect {
	if runtime.GOOS != "js" {
		md.mode = modeWelcome
	}

	md.initKeys()
	md.initWidgets()

	g := md.game
	_ = g
	md.applyConfig()
	// g.InitLevel()

	// if err != nil {
	// 	g.PrintStyled("Warning: could not load old saved game… starting new game.", logError)
	// 	log.Printf("Error: %v", err)
	// }

	return gruid.Sub(subSig)
}

// Update implements gruid.Model.Update. It handles keyboard and mouse input
// messages and updates the model in response to them.
func (md *model) Update(msg gruid.Msg) gruid.Effect {
	if _, ok := msg.(gruid.MsgInit); ok {
		return md.init()
	}

	if _, ok := msg.(gruid.MsgQuit); ok {
		// The user requested the end of the application (for example
		// by closing the window).
		return gruid.End()
	}

	return nil
}

// Draw implements gruid.Model.Draw. It draws a simple map that spans the whole
// grid.
func (m *model) Draw() gruid.Grid {
	return m.grid
}

func (md *model) applyConfig() {
	if GameConfig.NormalModeKeys != nil {
		md.keysNormal = GameConfig.NormalModeKeys
	}
	if GameConfig.TargetModeKeys != nil {
		md.keysTarget = GameConfig.TargetModeKeys
	}
}

func applyThemeConf() {
	if Only8Colors && !Tiles {
		ColorFgLOS = ColorGreen
	}
}
