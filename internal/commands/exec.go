package commands

import (
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/dontpanic"
	"devllart/foobarman/internal/scenes"
	"devllart/foobarman/internal/state"
	"strings"
)

func Exec() {
	defer dontpanic.RecoverAll()
	state.Command = strings.ToLower(state.Command)

	if state.Command == "" {
		return
	} else if CommandIs("exit") {
		state.Run = false
	} else if state.Command == "hideall" {
		config.HideAll()
	} else if state.Command == "showall" {
		config.ShowAll()
	} else if state.Command == "cmds" {
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
		state.AddInfof("! Незивестная комманда: %s\n", state.Command)
	}

	state.ClearTemp()
}
