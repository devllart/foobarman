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
	// scenes.Hello()

	state.Name = "Пётр"
	state.Scene = scenes.Store

	for state.Run == true {
		scenes.Show(state.Scene) // В scenes.Hello контекст сцены меняется на Store
		state.Info = ""          // Отчищаем подсказки и варнинги

		fmt.Print("> ") // Пользователь вводит команды после > |  — TerminalStyle ;)
		fmt.Scanf("%s %s %s", &state.Command, &state.Arg1, &state.Arg2)
		commands.Exec()
	}
}
