package drinks

import (
	"devllart/foobarman/internal/texts"
	"strings"
)

func init() {
	for name, drink := range AviableDrinks {
		delete(AviableDrinks, name)
		name = strings.Title(name)
		drink.Name = name
		drink.Taste = getTaste(drink)
		AviableDrinks[name] = drink
	}

	for name, coctail := range AviableCoctail {
		delete(AviableCoctail, name)
		name = strings.Title(name)
		coctail.Name = name
		AviableCoctail[name] = coctail
	}
}

func getTaste(drink DrinkInfo) *string {

	if taste := NewTastes.GetValue(drink.Name); taste != nil {
		return taste
	}

	if taste := NewTastes.GetValue(drink.Type); taste != nil {
		return taste
	}

	return NewTastes.GetValue(texts.Unknown)
}
