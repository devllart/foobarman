package state

import (
	"strconv"
	"strings"
)

// ** For work with commands
var Command = ""      // Current command
var Args = []string{} // Current args of command

func SetCommand(text string) {
	Command = ""
	for i, word := range strings.Split(text, " ") {
		if i == 0 {
			Command += word + " "
		} else if _, err := strconv.ParseFloat(word, 64); err == nil {
			Args = append(Args, word)
		} else {
			Command += word + " "
		}
	}

	for len(Args) < 3 {
		Args = append(Args, "")
	}

	Command = Command[:len(Command)-1] // delete space

	if len(Command) > 0 && Command[len(Command)-1] == 13 {
		Command = Command[:len(Command)-1]
	}
}
