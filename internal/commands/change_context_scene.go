package commands

import (
	"devllart/foobarman/internal/scenes"
	"devllart/foobarman/internal/state"
)

func SetScenesBar() {
	state.Scene = scenes.Bar
}
func SetScenesRecipes() {
	state.Scene = scenes.Recipes
}
func SetScenesStore() {
	state.Scene = scenes.Store
}
