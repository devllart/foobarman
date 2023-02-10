package commands

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/structs"
	"strconv"
	"strings"
)

/**
 * Get Name or Index product's and return correct its name.
 * + alerting user if his index out of range product's menu.
 */

func ProductExistYet(productName string, volume float64) (*structs.Product, int) {
	// Cycle cycle through the bar's products list
	for i := range state.Bar {
		product := state.Bar[i]

		// If the product has in the bar yet, then added to exist product
		if product.Name == productName && product.Volume == volume {
			return &product, i
		}
	}
	return nil, -1
}

func correctProductName(productNameOrIndex string) string {
	productIndex, err := strconv.Atoi(productNameOrIndex)
	if err == nil {
		if productIndex > len(state.ProductsIds) {
			alert.PanicNotAvailableIndexProduct(productIndex)
		}
		return state.ProductsIds[productIndex-1]
	}

	return strings.Title(productNameOrIndex)
}

func correctCount(arg string) int {
	var count int = 1
	var err error
	if arg != "" {
		count, err = strconv.Atoi(arg)
	}
	if err != nil || count <= 0 {
		alert.IncorrectAmountOfProduct()
		return 0
	}

	return count
}

func correctVolume(arg string) float64 {
	var volume float64 = 0
	var err error
	if arg != "" {
		volume, err = strconv.ParseFloat(arg, 64)
	}
	if err != nil || volume < 0 {
		alert.IncorrectVolumeOfProduct()
		return 0
	}

	return volume
}
