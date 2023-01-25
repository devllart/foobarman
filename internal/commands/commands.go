package commands

import (
	"devllart/foobarman/internal/state"
)

func Exec(command string) {
	if command == "exit" {
		state.Run = false
	} else {
    state.Info = "! Незивестная комманда: " + command
	}

}
