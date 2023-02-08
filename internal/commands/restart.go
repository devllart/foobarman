package commands

import (
	"devllart/foobarman/internal/scenes"
	"devllart/foobarman/internal/state"
)

/**
 * Restart game
 */

// Usualy restart
func Restart() {
	state.Restart()
	state.Scene = scenes.Store
	scenes.Show(scenes.Store)
	state.ClearTemp()
	state.Info = ""
	state.BarOpen = false
	state.NotSaler = true
	state.AddInfof("%B ! Игра перезапущена%C\n")
}

// Restart and rand buy
func RestartAndRandBuy() {
	Restart()
	state.TempBool = false
	state.RandomBuy = true
	buyRandom()
	state.RandomBuy = false
	state.Scene = scenes.Bar
	scenes.Show(scenes.Bar)
}
