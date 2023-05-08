package inputer

import (
	"devllart/foobarman/internal/commands"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/src/fmtc"
	"strings"

	"github.com/manifoldco/promptui"
)

var firstRun = true

func GetCommand(label string) (command string) {
	if firstRun {
		firstRun = false
		return ""
	}

	fmtc.Printf("\n")
	prompt := promptui.Prompt{
		Label:    fmtc.Sprintf(label),
		Validate: CheckInput,
	}

	command, err := prompt.Run()

	if err != nil {
		panic(err.Error())
	}

	// fmtc.Printf("%C") // Clear colorize
	return command
}

var err error

func CheckInput(input string) error {
	cmd := strings.ToLower(input)
	state.SetCommand(cmd)
	err = commands.FakeExec()
	return err
}
