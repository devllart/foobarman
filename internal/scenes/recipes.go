package scenes

import (
	"devllart/foobarman/internal/drinks"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"
	"fmt"
)

func Recipes() {
	fmtc.Printf(texts.SceneRecipes, funcs.Indent(15)) // Print scene

	i := 1
	// Print all recipes to console
	for _, coctail := range drinks.AviableCoctail {
		state.DrinksIds = append(state.DrinksIds, coctail.Name) // Added coctail to slice DrinksIds for available by index
		fmt.Printf("%d. ", i)                                   // Print index
		coctail.Show()                                          // Print info of coctail
		i += 1
	}
}
