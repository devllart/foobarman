package state

import (
	"devllart/foobarman/src/fmtc"
	"strings"
)

func (command *CommandStruct) ShowClue() {
	separator := fmtc.Sprintf("%C, %B")
	commands := strings.Join(command.Aliases, separator)
	fmtc.Printf(" %Y\\>%C %B%v%C — %G%s%C\n", commands, command.Description)
}
