package commands

import (
	"devllart/foobarman/src/fmtc"
	"strings"

	"github.com/TwiN/go-color"
)

type Command struct {
	Name        string
	Aliases     []string
	Description string
}

func (command Command) ShowClue() {
	commands := strings.Join(command.Aliases, color.Reset+", "+color.Red)
	fmtc.Printf(" %Y\\>%C %R%v%C — %G%s%C\n", commands, command.Description)
}
