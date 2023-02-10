package commands

import (
	"devllart/foobarman/internal/state"
)

func SetScenesBar() {
	state.Scene = "Bar"
}
func SetScenesRecipes() {
	state.Scene = "Recipes"
}
func SetScenesStore() {
	state.Scene = "Store"
}
