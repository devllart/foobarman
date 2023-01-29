package alert

import (
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
)

/**
 * With panic!
 */

func PanicNotAvailableDrink(drinkName string) {
	state.AddInfof(texts.NotAvailableDrink, drinkName)
	panic(texts.IncorrectInput)
}

func PanicNotAvailableIndexDrink(drinkIndex int) {
	state.AddInfof(texts.NotAvailableIndexDrink, drinkIndex)
	panic(texts.IncorrectInput)
}

func PanicNotEnoughtFundsToBuy(sumPrice float64) {
	if state.Command != "rand" {
		state.AddInfof(texts.NotEnoughFundsToBuy, sumPrice)
	}
	panic(texts.IncorrectInput)
}

func PanicNotVolumeOfDrink(drinkName string, volume float64) {
	state.AddInfof(texts.NotVolumeOfDrink, drinkName, volume)
	panic(texts.IncorrectInput)
}
