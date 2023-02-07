package commands

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/products"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/src/funcs"
)

func buyProduct(productNameOrIndex string, volume float64, count int) {
	productName := correctProductName(productNameOrIndex)

	product, exist := products.AvailableProducts[productName]
	if exist == false {
		alert.PanicNotAvailableProduct(productName)
	}

	// If user not inputed volume or input "0" then assign minimal avliable value (firt in slice)
	if volume == 0 {
		volume = product.AviableVolume[0]
	}

	index := funcs.IndexOfOrNull(volume, product.AviableVolume)

	// Calculate the total sum
	// If selected volume not exist, alert user and panic "IncorrectInput"
	if funcs.Contains(product.AviableVolume, volume) == false {
		alert.PanicNotVolumeOfProduct(productName, volume)
	}

	// Buy product transaction in goroutine (Oh yes I'm a pervert)
	buyTransaction(productName, count, volume, product.Prices[index], state.RandomBuy)
}
