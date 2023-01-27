package drinks

import (
	"devllart/foobarman/internal/texts"
	"strings"
)

var AviableDrinks = map[string]DrinkInfo{
	"Содовая": {
		Type:          "Содовая",
		Alc:           0,
		AviableVolume: []float64{.25, .5, 1, 1.5},
		Prices:        []float64{.3, .5, .7, 1},
		Description:   texts.DrinkDescSoda,
	},
	"Эбботтс Биттер": {
		Type:          "Ароматический биттер",
		Alc:           41.5,
		AviableVolume: []float64{.1},
		Prices:        []float64{20.50},
		Description:   texts.DrinkDescAbbotsBitter,
	},
	"Апероль": {
		Type:          "Апероль",
		Alc:           22,
		AviableVolume: []float64{0.75, 1},
		Prices:        []float64{27.99, 34.71},
		Description:   texts.DrinkDescAperrol,
	},
	"Просекко": {
		Type:          "Сухое вино",
		Alc:           11,
		AviableVolume: []float64{0.75},
		Prices:        []float64{14.94},
		Description:   texts.DrinkDescProsecco,
	},
	"Cruzan Aged Light Rum": {
		Type:          "Белый ром",
		Alc:           40,
		AviableVolume: []float64{0.375, 0.75},
		Prices:        []float64{7.99, 13.19},
		Description:   texts.DrinkDescWhiteRom,
	},
	"Rom Mocambo 20 Year Old": {
		Type:          "Ром",
		Alc:           40,
		AviableVolume: []float64{0.75},
		Prices:        []float64{42.99},
		Description:   texts.DrinkDescWhiteRom,
	},
	"Tito's Handmade Vodka": {
		Type:          "Водка",
		Alc:           40,
		AviableVolume: []float64{1, 1.75},
		Prices:        []float64{28.40, 36.29},
		Description:   texts.DrinkDescVodka,
	},
	"Belsazar Vermouth": {
		Type:          "Красный вермут",
		Alc:           40,
		AviableVolume: []float64{0.375, 0.75},
		Prices:        []float64{7.99, 13.19},
		Description:   texts.DrinkDescVermouth,
	},
}

func init() {
	for name, drink := range AviableDrinks {
		name = strings.Title(name)
		drink.Name = name
		AviableDrinks[name] = drink
	}
}
