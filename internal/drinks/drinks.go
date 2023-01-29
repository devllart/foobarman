package drinks

import (
	"devllart/foobarman/internal/texts"
	"strings"
)

func init() {
	for name, drink := range AviableDrinks {
		name = strings.Title(name)
		drink.Name = name
		drink.Taste = getTaste(drink)
		AviableDrinks[name] = drink
	}

	for name, coctail := range AviableCoctail {
		name = strings.Title(name)
		coctail.Name = name
		AviableCoctail[name] = coctail
	}
}

func getTaste(drink DrinkInfo) string {
	if taste, exist := Tastes[drink.Name]; exist == true {
		return taste
	}
	if taste, exist := Tastes[drink.Type]; exist == true {
		return taste
	}

	return texts.Unknown
}
