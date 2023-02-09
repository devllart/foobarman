package main

import (
	"bufio"
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/commands"
	"devllart/foobarman/internal/sale"
	"devllart/foobarman/internal/scenes"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

/**
 * This game is simulator barman.
 * By playing it, you yourself may become a barman.
 * So be careful... Good luck ;)
 */
var reader = bufio.NewReader(os.Stdin)

func main() {
	funcs.CliClear() // Clear console
	scenes.Hello()   // Run scene "Hello", In the scene ask name's barman

	// Run game cycle
	for state.Run == true {
		state.HandlerStatus()    // Turning status of barman
		scenes.Show(state.Scene) // In scenes.Hello global scene's context was changed to "Store"
		alert.ClearInfo()        // Clear hints and warning message
		sale.Sale()

		// fmtc.Printf("\n%Y\\>%B ") // User will enter command after symbol > |  — TerminalStyle ;)
		// state.SetCommand(fmtc.Scan('\n')) // Gets command and args from STDIN
		fmtc.Printf("\n")
		prompt := promptui.Prompt{
			Label:    fmtc.Sprintf("%Y \\> "),
			Validate: Inputer,
		}

		result, err := prompt.Run()

		if err != nil {
			panic(err.Error())
		}
		state.SetCommand(result)

		fmtc.Printf("%C")
		commands.Exec() // Try execute user command's
	}
}

func Inputer(input string) error {
	state.SetCommand(strings.ToLower(input))
	return commands.FakeExec()
}
