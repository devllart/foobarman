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

	// Print all barmans've product to console
	for i, product := range state.Bar {
		state.ProductsIds = append(state.ProductsIds, product.Name) // Added product to slice ProductsIds for available by index
		fmtc.Printf("%d. ", i+1)                                    // Print index
		product.Show()                                              // Print info of product
	}
}
