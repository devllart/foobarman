package scenes

import (
	"devllart/foobarman/internal/state"
	"fmt"
)

func Hello() {
	fmt.Print("Ну и как тебя зовут юный бармен: ")
	fmt.Scan(&state.Name)
	fmt.Printf("Отлично %s, ну что пошлите в магазин закупаться? (Нажми Enter)", state.Name)
	state.Scene = Store
}
