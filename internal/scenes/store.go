package scenes

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/drinks"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
)

/**
 * Store "Bri Larson" has a variety of drinks especially for you.
 */

func Store() {
	fmtc.Printf(texts.StoreHello, state.Name) // Print greetings
	// Cycle by available drinks
	for name, drink := range drinks.AviableDrinks {
		state.DrinksIds = append(state.DrinksIds, name) // Added drink to slice DrinksIds for available by index
		index := len(state.DrinksIds)
		drink.PrintInStore(index)
		drink.PrettyDescription()

		drink.PrintPrices()
	}
	alert.ClueStore() // Print clue how buying drinks
}
