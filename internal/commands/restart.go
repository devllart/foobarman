package commands

import (
	"devllart/foobarman/internal/scenes"
	"devllart/foobarman/internal/state"
)

func Restart() {
	state.Restart()
	state.Scene = scenes.Store
	scenes.Show(scenes.Store)
}

func RestartAndRandBuy() {
	Restart()
	state.TempBool = false
	state.RandomBuy = true
	buyRandom()
	state.RandomBuy = false
	state.Scene = scenes.Bar
	scenes.Show(scenes.Bar)
}
