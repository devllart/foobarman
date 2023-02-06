package drinks

import (
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/mapsi"
	"strings"
)

var DrinksStandartTypesPrice = map[string]float64{}

var MapsiAvailableCoctail mapsi.Mapsi[Coctail]
var MapsiAvailableDrinks mapsi.Mapsi[DrinkInfo]

func init() {
	for name, drink := range AviableDrinks {
		delete(AviableDrinks, name)
		name = strings.Title(name)
		drink.Name = name
		drink.Taste = getTaste(drink)
		if _, exist := DrinksStandartTypesPrice[drink.Type]; exist {
			DrinksStandartTypesPrice[drink.Type] += drink.Prices[0] * drink.AviableVolume[0] * 0.5
		} else {
			DrinksStandartTypesPrice[drink.Type] = drink.Prices[0] * drink.AviableVolume[0]
		}
		AviableDrinks[name] = drink
	}

	for name, coctail := range AviableCoctail {
		delete(AviableCoctail, name)
		name = strings.Title(name)
		coctail.Name = name
		coctail.Price = coctail.GetPrice()
		AviableCoctail[name] = coctail
	}

	MapsiAvailableCoctail = mapsi.New(AviableCoctail)
	MapsiAvailableDrinks = mapsi.New(AviableDrinks)
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
