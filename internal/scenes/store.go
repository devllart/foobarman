package scenes

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/products"
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
	for i, name := range drinks.MapsiAvailableDrinks.Keys {
		drink := drinks.MapsiAvailableDrinks.Values[i]
		state.DrinksIds = append(state.DrinksIds, name) // Added drink to slice DrinksIds for available by index
		index := len(state.DrinksIds)
		drink.PrintInStore(index)
		drink.PrettyDescription()

		drink.PrintPrices()
	}
	alert.ClueStore() // Print clue how buying drinks
}
