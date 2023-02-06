package alert

import (
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
)

/**
 * Alert for drinks don't panic
 */

// Bought drink
func IncorrectAmountOfProduct() {
	state.AddInfof(texts.IncorrectAmountOfProduct)
}

func IncorrectVolumeOfProduct() {
	state.AddInfof(texts.IncorrectVolumeOfProduct)
}

func ProductBought(drinkName, typeVolume string, volume float64, sumPrice float64, count int) {
	state.AddInfof(texts.ProductBought, drinkName, volume, typeVolume, count, sumPrice)
}

func ProductBoughtYet(drinkName, typeVolume string, volume, sumPrice float64, count, countSum int) {
	state.AddInfof(texts.ProductBoughtYet, drinkName, volume, typeVolume, count, countSum, sumPrice)
}
