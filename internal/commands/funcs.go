package commands

import (
	"devllart/foobarman/internal/state"
	"devllart/foobarman/src/funcs"
	"strings"
)

func CommandIs(key string) bool {
	state.Command = strings.Title(state.Command)
	return funcs.Contains(AviableCommand[key], state.Command)
}
