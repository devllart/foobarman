package main

import (
	"devllart/foobarman/internal/commands"
	"devllart/foobarman/internal/scenes"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/src/funcs"
	"fmt"
)

func main() {
	funcs.CliClear()
	scenes.Hello()

	var command, arg, arg2 string
	for state.Run == true {
		scenes.Show(state.Scene) // В scenes.Hello контекст сцены меняется на Store
		state.Info = ""          // Отчищаем подсказки и варнинги

		fmt.Print("> ") // Пользователь вводит команды после > |  — TerminalStyle ;)
		fmt.Scanf("%s %s %s", &command, &arg, &arg2)
		commands.Exec(command, arg, arg2)
	}
}
