package scenes

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/products"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
)

/**
 * Store "Bri Larson" has a variety of products especially for you.
 */

func Store() {
	fmtc.Printf(texts.StoreHello, state.Name) // Print greetings
	// Cycle by available products
	for i, name := range products.MapsiAvailableProducts.Keys {
		product := products.MapsiAvailableProducts.Values[i]
		state.ProductsIds = append(state.ProductsIds, name) // Added product to slice ProductsIds for available by index
		index := len(state.ProductsIds)
		product.PrintInStore(index)
		product.PrettyDescription()

		product.PrintPrices()
	}
	alert.ClueStore() // Print clue how buying products
}
