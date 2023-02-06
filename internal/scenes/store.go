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
	for i, name := range products.MapsiAvailableProducts.Keys {
		drink := products.MapsiAvailableProducts.Values[i]
		state.ProductsIds = append(state.ProductsIds, name) // Added drink to slice ProductsIds for available by index
		index := len(state.ProductsIds)
		drink.PrintInStore(index)
		drink.PrettyDescription()

		drink.PrintPrices()
	}
	alert.ClueStore() // Print clue how buying drinks
}
