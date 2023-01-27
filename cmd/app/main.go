package main

import (
	"devllart/foobarman/internal/commands"
	"devllart/foobarman/internal/scenes"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"
)

/**
 * This game is simulator barman.
 * By playing it, you yourself may become a barman.
 * So be careful... Good luck ;)
 */
func main() {
	funcs.CliClear() // Clear console
	scenes.Hello()   // Run scene "Hello", In the scene ask name's barman

	// Run game cycle
	for state.Run == true {
		scenes.Show(state.Scene) // In scenes.Hello global scene's context was changed to "Store"
		state.Info = ""          // Clear hints and warning message

		fmtc.Printf("\n%Y\\>%B ")         // User will enter command after symbol > |  — TerminalStyle ;)
		state.SetCommand(fmtc.Scan('\n')) // Gets command and args from STDIN
		// fmt.Scanf("%s %s %s", &state.Command, &state.Arg1, &state.Arg2) // Gets command and args from STDIN
		fmtc.Printf("%C")
		commands.Exec() // Try execute user command's
	}
}
