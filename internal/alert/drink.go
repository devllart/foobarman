package alert

import (
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
)

/**
 * Don't panic
 */

// Bought drink
func IncorrectAmountOfDrink() {
	state.AddInfof(texts.IncorrectAmountOfDrink)
}

func IncorrectVolumeOfDrink() {
	state.AddInfof(texts.IncorrectVolumeOfDrink)
}

func DrinkBought(drinkName, typeVolume string, volume float64, sumPrice float64, count int) {
	state.AddInfof(texts.DrinkBought, drinkName, volume, typeVolume, count, sumPrice)
}

func DrinkBoughtYet(drinkName, typeVolume string, volume, sumPrice float64, count, countSum int) {
	state.AddInfof(texts.DrinkBoughtYet, drinkName, volume, typeVolume, count, countSum, sumPrice)
}
