package scenes

import (
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"
	"fmt"
)

/**
 * For decency, you need to say hello and ask for a name.
 * (Otherwise, how are we going to contact our bartender?)
 */

func Hello() { // Hello
	for state.Name == "" {
		funcs.CliClear()              // Clear console
		fmtc.Printf(texts.WhatIsName) // Ask Name
		state.Name = fmtc.Scan('\n')  // Read line and write value in state.Name
	}

	state.Name = fmtc.Sprintf("%B%s%C", state.Name) // Set color of Name
	fmtc.Printf(texts.NameIsOk, state.Name)         // Say Name is valid
	fmt.Scanln(&state.TempStr)                      // Wait Input
	state.Scene = Store                             // Change context scene
}
