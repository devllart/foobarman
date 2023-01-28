package scenes

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/drinks"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
)

func Store() {
	fmtc.Printf(texts.StoreHello, state.Name)
	alert.ClueStore()
	for name, drink := range drinks.AviableDrinks {
		state.DrinksIds = append(state.DrinksIds, name)
		index := len(state.DrinksIds)
		drink.PrintInStore(index)
		drink.PrettyDescription()

		drink.PrintPrices()
	}
}
