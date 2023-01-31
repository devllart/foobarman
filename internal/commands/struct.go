package commands

import (
	"devllart/foobarman/src/fmtc"
	"strings"
)

type Command struct {
	Name        string
	Aliases     []string
	Description string
}

func (command Command) ShowClue() {
	separator := fmtc.Sprintf("%C, %B")
	commands := strings.Join(command.Aliases, separator)
	fmtc.Printf(" %Y\\>%C %R%v%C — %G%s%C\n", commands, command.Description)
}
