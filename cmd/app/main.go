package main

import (
	"devllart/foobarman/internal/commands"
	"devllart/foobarman/internal/scenes"
	"devllart/foobarman/internal/state"
	"fmt"
	"strings"
)

func main() {
	var command string
	for state.Run == true {
		scenes.Show(scenes.Bar)
		fmt.Print("> ")
		fmt.Scan(&command)
		commands.Exec(strings.ToLower(command))
	}
}
