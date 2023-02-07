package alert

import (
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
)

/**
 * Alert for products don't panic
 */

// Bought product
func IncorrectAmountOfProduct() {
	state.AddInfof(texts.IncorrectAmountOfProduct)
}

func IncorrectVolumeOfProduct() {
	state.AddInfof(texts.IncorrectVolumeOfProduct)
}

func ProductBought(productName, typeVolume string, volume float64, sumPrice float64, count int) {
	state.AddInfof(texts.ProductBought, productName, volume, typeVolume, count, sumPrice)
}

func ProductBoughtYet(productName, typeVolume string, volume, sumPrice float64, count, countSum int) {
	state.AddInfof(texts.ProductBoughtYet, productName, volume, typeVolume, count, countSum, sumPrice)
}
