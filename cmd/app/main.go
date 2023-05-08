package main

import (
	"bufio"
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/commands"
	"devllart/foobarman/internal/inputer"
	"devllart/foobarman/internal/products"
	"devllart/foobarman/internal/sale"
	"devllart/foobarman/internal/scenes"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/src/funcs"
	"os"
	"strings"
)

/**
 * This game is simulator barman.
 * By playing it, you yourself may become a barman.
 * So be careful... Good luck ;)
 */

var reader = bufio.NewReader(os.Stdin)

func main() {
	state.Load()

	for i := 2; i <= state.Stage; i++ {
		products.AddAvailablesCoctail(i)
	} // duct tape

	// fmt.Print(products.NotExitst)
	// panic("Aaa")

	cliApp()
}

func cliApp() {
	lenArgs := len(os.Args)
	if lenArgs == 1 {
		gamePlay()
		return
	}

	commandExec(strings.Join(os.Args[1:lenArgs], "")) // Execute command
	gameActions()
	state.Save() // Save state game
}

func gamePlay() {
	funcs.CliClear() // Clear console
	scenes.Hello()   // Run scene "Hello", In the scene ask name's barman

	// Run game cycle
	for state.Run == true {
		gameActions()                              // Run game actions
		commandExec(inputer.GetCommand("%Y \\> ")) // Get command from input of user and execute
		state.Save()                               // Save state game
	}
}

func gameActions() {
	sale.Sale()              // If coctail is ready then sale to client the coctail
	state.HandlerStatus()    // Turning status of barman
	scenes.Show(state.Scene) // In scenes.Hello global scene's context was changed to "Store"
	alert.ClearInfo()        // Clear hints and warning message
}

func commandExec(command string) {
	state.SetCommand(command) // Set command
	commands.Exec()           // Try execute user command's
}
