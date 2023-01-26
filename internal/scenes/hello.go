package scenes

import (
	"devllart/foobarman/internal/state"
	"devllart/foobarman/src/funcs"
	"fmt"

	"github.com/TwiN/go-color"
)

func Hello() {
	for state.Name == "" {
		funcs.CliClear()
		fmt.Printf("%sНу и как тебя зовут юный бармен: %s", color.Reset, color.Blue)
		fmt.Scanf("%s\n", &state.Name)
	}

	state.Name = color.Blue + state.Name + color.Reset
	fmt.Printf("%sОтлично %s, ну что пошлите в магазин закупаться? (Нажми Enter)", color.Reset, state.Name)
	fmt.Scanf("%s\n", &state.TempStr)
	state.Scene = Store
}
