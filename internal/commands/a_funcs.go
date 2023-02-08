package commands

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/scenes"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/src/funcs"
	"strings"
)

/**
 * Funcs helpers
 */

// Check context command matches key (key is expected command)
// The key is also used to refer to the map AvailableCommands
func CommandIs(key string) bool {
	command := strings.ToLower(state.Command)                              // All commands must be in lowercase
	key = strings.ToLower(key)                                             // Of course the key too
	return funcs.Contains(state.AvailableCommands()[key].Aliases, command) // true if command match to aleasses command
}

// Skip client
func SkipClient() {
	state.NotSaler = true
	state.AddInfof("%YТы упустил клиента%C\n")
	text := state.SubtractFromSalary(state.Money / 99)
	state.AddInfof(text)

	if len(state.CurrentHistory) > 0 {
		state.AddInfof("\n%YИ ещё похоже ты упустил интересную историю%C\n")
	}
}

/**
 * Restart game
 */

// Usualy restart
func Restart() {
	state.Restart()
	state.Scene = scenes.Store
	scenes.Show(scenes.Store)
	alert.RestartGame()
}

// Restart and rand buy
func RestartAndRandBuy() {
	Restart()
	buyRandom()
	state.Scene = scenes.Bar
	scenes.Show(scenes.Bar)
}
