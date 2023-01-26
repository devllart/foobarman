package commands

import (
	"devllart/foobarman/internal/data"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/src/funcs"
	"strings"
)

func CommandIs(key string) bool {
	command := strings.ToLower(state.Command)
	key = strings.ToLower(key)
	return funcs.Contains(data.AvailableCommands[key].Aliases, command)
}
