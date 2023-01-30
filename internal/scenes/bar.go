package scenes

import (
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"
)

/**
 * Come often to our modest tavern.
 */

func Bar() {
	fmtc.Printf(texts.SceneBar, funcs.Indent(15)) // Print scene

	// Print all barmans've drink to console
	for i, drink := range state.Bar {
		state.DrinksIds = append(state.DrinksIds, drink.Name) // Added drink to slice DrinksIds for available by index
		fmtc.Printf("%d. ", i+1)                              // Print index
		drink.Show()                                          // Print info of drink
	}
}
