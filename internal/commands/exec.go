package commands

import (
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/dontpanic"
	"devllart/foobarman/internal/scenes"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
	"strings"
)

func Exec() {
	defer dontpanic.RecoverAll()
	state.Command = strings.ToLower(state.Command)

	if state.Command == "" { // Empty command -> None action
		return
	} else if CommandIs("exit") { // command exit -> close app, exit from app, .... you undestened
		state.Run = false
	} else if CommandIs("restart") { // command exit -> close app, exit from app, .... you undestened
		Restart()
	} else if CommandIs("restartrand") { // command exit -> close app, exit from app, .... you undestened
		RestartAndRandBuy()
	} else if CommandIs("hideall") { // Standart commands start here
		config.HideAll()
	} else if CommandIs("showall") {
		config.ShowAll()
	} else if CommandIs("cmds") {
		config.TurnShowCommands()
	} else if CommandIs("description") {
		config.TurnShowDescription()
	} else if CommandIs("instruction") {
		config.TurnShowInstruction()
	} else if CommandIs("bar") { // Change scenes context commands start here
		state.Scene = scenes.Bar
	} else if CommandIs("recipes") { // Change scenes context commands start here
		state.Scene = scenes.Recipes
	} else if CommandIs("store") {
		state.Scene = scenes.Store
	} else if scenes.CurrentIs(scenes.Store) { // Context dependent commands start here
		buy()
	} else if scenes.CurrentIs(scenes.Bar) && CommandIs("mix") {
		mix()
	} else { // otherwise say user "not command"
		state.AddInfof(texts.UnknownCommand, "%B"+state.Command+"%C")
	}

	state.ClearTemp() // Cleaning
}
