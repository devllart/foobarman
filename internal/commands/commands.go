package commands

import (
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/funcs"
	"devllart/foobarman/internal/state"
)

func Exec(command string) {
	if funcs.Contains([]string{"exit", "выйти"}, command) {
		state.Run = false
	} else if funcs.Contains([]string{"desc", "description", "описание"}, command) {
		config.ShowDescription = !config.ShowDescription
	} else {
		state.Info = "! Незивестная комманда: " + command
	}

}
