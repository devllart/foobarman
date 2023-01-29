package state

import (
	"devllart/foobarman/src/fmtc"
	"strings"

	"github.com/TwiN/go-color"
)

func (command *CommandStruct) ShowClue() {
	commands := strings.Join(command.Aliases, color.Reset+", "+color.Red)
	fmtc.Printf(" %Y\\>%C %R%v%C — %G%s%C\n", commands, command.Description)
}
