package commands

import (
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/scenes"
	"devllart/foobarman/internal/state"
	"strings"
)

var AviableCommand = map[string][]string{
	"exit":        {"Exit", "Quit", "Выйти", "Пока"},
	"description": {"Desc", "Description", "Описание"},
	"ok":          {"Ok", "Ок", "Всё", "Закончить"},
	"store":       {"Store", "Магазин"},
	"mix":         {"Mix", "Смешать"},
}

func Exec() {
	state.Command = strings.Title(state.Command)

	if state.Command == "" {
		return
	} else if CommandIs("exit") {
		state.Run = false
	} else if state.Command == "Hideall" {
		config.HideAll()
	} else if state.Command == "Showall" {
		config.ShowAll()
	} else if state.Command == "Cmds" {
		config.ShowCommands = !config.ShowCommands
	} else if CommandIs("description") {
		config.ShowDescription = !config.ShowDescription
	} else if scenes.CurrentIs(scenes.Store) {
		Buy()
	} else if !scenes.CurrentIs(scenes.Store) && CommandIs("store") {
		state.Scene = scenes.Store
	} else if CommandIs("mix") {
		Mix()
	} else {
		state.AddInfof("! Незивестная комманда: %s\n", state.Command)
	}

	state.ClearTemp()
}
