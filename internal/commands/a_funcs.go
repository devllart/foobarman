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
	textSubMoney := state.SubtractFromSalary(state.Money / 99)

	alert.MissClient()
	alert.Clue(textSubMoney)

	if len(state.CurrentHistory) > 0 {
		alert.MissHistory()
	}
}

/**
 * Restart game
 */

// Usualy restart
func Restart() {
	state.Restart()
	state.Scene = "Store"
	scenes.Show("Store")
	alert.RestartGame()
}

// Restart and rand buy
func RestartAndRandBuy() {
	Restart()
	buyRandom()
	state.Scene = "Bar"
	scenes.Show("Bar")
}
