package scenes

import (
	"devllart/foobarman/internal/state"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"
	"fmt"

	"github.com/TwiN/go-color"
)

func Hello() {
	for state.Name == "" {
		funcs.CliClear()
		fmtc.Printf("%CНу и как тебя зовут юный бармен: %B")
		fmt.Scanf("%s\n", &state.Name)
	}

	state.Name = color.Blue + state.Name + color.Reset
	fmtc.Printf("%CОтлично %s, ну что пошлите в магазин закупаться? (Нажми Enter): ", state.Name)
	fmt.Scanf("%s\n", &state.TempStr)
	state.Scene = Store
}
