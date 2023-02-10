package scenes

import (
	"devllart/foobarman/internal/products"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"
	"fmt"
)

/**
 * The best recipes only with us | ru.inshaker.com
 */

func Recipes() {
	fmtc.Printf(texts.SceneRecipes, funcs.Indent(15)) // Print scene

	i := 1 // index for cycle
	// Print all recipes to console
	for _, coctail := range products.MapsiAvailableCoctail.Values {
		state.ProductsIds = append(state.ProductsIds, coctail.Name) // Added coctail to slice ProductsIds for available by index
		fmt.Printf("%d. ", i)                                       // Print index
		coctail.Show()                                              // Print info of coctail
		coctail.PrettyDescription()                                 // Print pretty description of coctail
		coctail.PrettyInstruction()                                 // Print pretty instruction of coctail
		i += 1                                                      // iteration index of cycle
	}
}
