package commands

import (
	"devllart/foobarman/internal/state"
	"strconv"
)

func correctIndexProductName(productName string) int {
	if index, err := strconv.Atoi(productName); err == nil {
		if index > len(state.ProductsIds) {
			return -1
		}
		return index - 1
	}

	for i := range state.ProductsIds {
		if productName == state.ProductsIds[i] {
			return i
		}
	}

	return -1
}
