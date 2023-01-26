package structs

import (
	"fmt"
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
	fmt.Printf(" %s\\>%s %s%v%s — %s%s%s\n", color.Yellow, color.Reset, color.Red, commands, color.Reset, color.Green, command.Description, color.Reset)
}
