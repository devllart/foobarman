package commands

import (
	"devllart/foobarman/internal/data"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/src/funcs"
	"strings"
)

/**
 * Check context command matches key (key is expected command)
 * The key is also used to refer to the map AvailableCommands
 */

func CommandIs(key string) bool {
	command := strings.ToLower(state.Command)                           // All commands must be in lowercase
	key = strings.ToLower(key)                                          // Of course the key too
	return funcs.Contains(data.AvailableCommands[key].Aliases, command) // true if command match to aleasses command
}
