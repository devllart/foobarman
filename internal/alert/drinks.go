package alert

import (
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
)

/**
 * Alert for drinks don't panic
 */

// Bought drink
func IncorrectAmountOfDrink() {
	state.AddInfof(texts.IncorrectAmountOfDrink)
}

func IncorrectVolumeOfDrink() {
	state.AddInfof(texts.IncorrectVolumeOfDrink)
}

func DrinkBought(drinkName, typeVolume string, volume float64, sumPrice float64, count int) {
	text := fmtc.Sprintf(texts.DrinkBought, drinkName, volume, typeVolume, count, sumPrice)
	state.Alerts = append(state.Alerts, text)
}

func DrinkBoughtYet(drinkName, typeVolume string, volume, sumPrice float64, count, countSum int) {
	text := fmtc.Sprintf(texts.DrinkBoughtYet, drinkName, volume, typeVolume, count, countSum, sumPrice)
	state.Alerts = append(state.Alerts, text)
}
