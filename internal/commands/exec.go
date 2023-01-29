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

	if state.Command == "" {
		return
	} else if CommandIs("exit") {
		state.Run = false
	} else if CommandIs("recipes") {
		state.Scene = scenes.Recipes
	} else if CommandIs("hideall") {
		config.HideAll()
	} else if CommandIs("showall") {
		config.ShowAll()
	} else if CommandIs("cmds") {
		config.ShowCommands = !config.ShowCommands
	} else if CommandIs("description") {
		config.ShowDescription = !config.ShowDescription
	} else if scenes.CurrentIs(scenes.Store) {
		buy()
	} else if !scenes.CurrentIs(scenes.Store) && CommandIs("store") {
		state.Scene = scenes.Store
	} else if CommandIs("mix") {
		Mix()
	} else {
		state.AddInfof(texts.UnknownCommand, "%B"+state.Command+"%C")
	}

	state.ClearTemp()
}
