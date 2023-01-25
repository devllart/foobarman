package commands

import (
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/funcs"
	"devllart/foobarman/internal/scenes"
	"devllart/foobarman/internal/state"
	"fmt"
	"strings"
)

var AviableCommand = map[string][]string{
	"exit":        {"Exit", "Quit", "Выйти", "Пока"},
	"description": {"Desc", "Description", "Описание"},
}

func Exec(command, arg, arg2 string) {
	command = strings.Title(command)

	if funcs.Contains(AviableCommand["exit"], command) {
		state.Run = false
	} else if funcs.Contains(AviableCommand["description"], command) {
		config.ShowDescription = !config.ShowDescription
	} else if scenes.CurrentIs(scenes.Store) {
		if command == "" {
			return
		} else if command == "Ок" {
			state.Scene = scenes.Bar
			return
		} else {
			volume := CorrectVolume(arg)
			count := CorrectCount(arg2)
			if count != 0 {
				BayDrink(command, volume, count)
			}
		}
	} else if scenes.CurrentIs(scenes.Store) == false {
		state.Scene = scenes.Store
	} else {
		state.Info += fmt.Sprintf("! Незивестная комманда: %s\n", command)
	}

}
