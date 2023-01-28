package drinks

import (
	"devllart/foobarman/internal/texts"
	"strings"
)

var AviableDrinks = map[string]DrinkInfo{

	"Лимон": {
		Type:          "Фрукт",
		Alc:           0,
		AviableVolume: []float64{1, 2, 3, 5},
		Prices:        []float64{5, 8, 12, 14},
		Description:   texts.DescLimon,
	},
	"Лайм": {
		Type:          "Фрукт",
		Alc:           0,
		AviableVolume: []float64{1, 2, 3, 5},
		Prices:        []float64{6, 10, 13, 16.5},
		Description:   texts.DescLaim,
	},
	"Содовая": {
		Type:          "Содовая",
		Alc:           0,
		AviableVolume: []float64{.25, .5, 1, 1.5},
		Prices:        []float64{.3, .5, .7, 1},
		Description:   texts.DescSoda,
	},
	"Эбботтс Биттер": {
		Type:          "Ароматический биттер",
		Alc:           41.5,
		AviableVolume: []float64{.1},
		Prices:        []float64{20.50},
		Description:   texts.DescAbbotsBitter,
	},
	"Апероль": {
		Type:          "Апероль",
		Alc:           22,
		AviableVolume: []float64{0.75, 1},
		Prices:        []float64{27.99, 34.71},
		Description:   texts.DescAperrol,
	},
	"Просекко": {
		Type:          "Просекко",
		Alc:           11,
		AviableVolume: []float64{0.75},
		Prices:        []float64{14.94},
		Description:   texts.DescProsecco,
	},
	"Cruzan Aged Light Rum": {
		Type:          "Белый ром",
		Alc:           40,
		AviableVolume: []float64{0.375, 0.75},
		Prices:        []float64{7.99, 13.19},
		Description:   texts.DescWhiteRom,
	},
	"Rom Mocambo 20 Year Old": {
		Type:          "Ром",
		Alc:           40,
		AviableVolume: []float64{0.75},
		Prices:        []float64{42.99},
		Description:   texts.DescWhiteRom,
	},
	"Titos Handmade Vodka": {
		Type:          "Водка",
		Alc:           40,
		AviableVolume: []float64{1, 1.75},
		Prices:        []float64{28.40, 36.29},
		Description:   texts.DescVodka,
	},
	"Belsazar Vermouth": {
		Type:          "Красный вермут",
		Alc:           40,
		AviableVolume: []float64{0.375, 0.75},
		Prices:        []float64{7.99, 13.19},
		Description:   texts.DescVermouth,
	},
}

func init() {
	for name, drink := range AviableDrinks {
		name = strings.Title(name)
		drink.Name = name
		AviableDrinks[name] = drink
	}
}
