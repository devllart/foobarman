package commands

import (
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/scenes"
	"devllart/foobarman/internal/state"
	"fmt"
	"strings"
)

var AviableCommand = map[string][]string{
	"exit":        {"Exit", "Quit", "Выйти", "Пока"},
	"description": {"Desc", "Description", "Описание"},
	"ok":          {"Ok", "Ок", "Всё", "Закончить"},
	"store":       {"Store", "Магазин"},
	"mix":         {"Mix", "Смешать"},
}

func Exec(command, arg1, arg2 string) {
	command = strings.Title(command)
	state.Command = command
	state.Arg1 = arg1
	state.Arg2 = arg2

	if command == "" {
		return
	} else if CommandIs("exit") {
		state.Run = false
	} else if CommandIs("description") {
		config.ShowDescription = !config.ShowDescription
	} else if scenes.CurrentIs(scenes.Store) {
		Buy()
	} else if !scenes.CurrentIs(scenes.Store) && CommandIs("store") {
		state.Scene = scenes.Store
	} else if CommandIs("mix") {
		Mix()
	} else {
		state.Info += fmt.Sprintf("! Незивестная комманда: %s\n", command)
	}
}
