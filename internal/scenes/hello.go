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
	for state.RawName == "" {
		funcs.CliClear()                // Clear console
		fmtc.Printf(texts.WhatIsName)   // Ask Name
		state.RawName = fmtc.Scan('\n') // Read line and write value in state.Name
	}

	state.Name = fmtc.Sprintf("%B%s%C", state.RawName) // Set color of Name

	// 	fmtc.Printf(`
	//         Добро пожаловать %s
	// `, state.Name)

	fmtc.Printf(texts.NameIsOk, state.Name) // Say Name is valid
	fmt.Scanln(&state.TempStr)              // Wait Input
	state.Scene = "Store"                   // Change context scene
}
