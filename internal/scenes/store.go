package scenes

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/drinks"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
	"fmt"
)

func Store() {
	fmtc.Printf(texts.StoreHello, state.Name)
	alert.ClueStore()
	i := 1
	for name, drink := range drinks.AviableDrinks {

		fmtc.Printf(texts.StoreDrinkInfo, i, name, drink.Alc, drink.Type, drink.GetTaste())
		drink.PrettyDescription()

		state.DrinksIds = append(state.DrinksIds, name)
		for i := range drink.Prices {
			price := drink.Prices[i]
			fmtc.Printf(texts.StoreDrinkInfoPrice, price, drink.AviableVolume[i])
		}
		fmt.Print(" |\n\n")
		i += 1
	}
}
