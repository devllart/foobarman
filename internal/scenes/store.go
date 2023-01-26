package scenes

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/drinks"
	"devllart/foobarman/internal/ptf"
	"devllart/foobarman/internal/state"
	"fmt"

	"github.com/TwiN/go-color"
)

func Store() {
	ptf.StoreHello(state.Name)
	alert.ClueStore()
	i := 1
	for name, drink := range drinks.AviableDrinks {
		fmt.Printf("%d. %s%s%s (%s%.2f%s процентов алкоголя) | %s%s%s (вкус %s%s%s)\n", i, color.Blue, name, color.Reset, color.Red, drink.Alc, color.Reset, color.Yellow, drink.Type, color.Reset, color.Green, drink.Taste, color.Reset)
		drink.PrettyDescription()

		state.DrinksIds = append(state.DrinksIds, name)
		for i := range drink.Prices {
			price := drink.Prices[i]
			fmt.Printf("| %s%.3f$%s за %s%.3f .л%s", color.Yellow, price, color.Reset, color.Green, drink.AviableVolume[i], color.Reset)
		}
		fmt.Print(" |\n\n")
		i += 1
	}
}
