package commands

import (
	"devllart/foobarman/internal/products"
	"devllart/foobarman/internal/state"
	"fmt"
	"math/rand"
)

func buy() {
	if CommandIs("ok") {
		state.Scene = "Bar"
		return
	} else if CommandIs("rand") {
		buyRandom()
	} else {
		volume := correctVolume(state.Args[0])
		count := correctCount(state.Args[1])
		if count != 0 {
			buyProduct(state.Command, volume, count)
		}
	}
}

func buyRandom() {
	state.RandomBuy = true
	for _, product := range products.MapsiAvailableProducts.Values {
		buyProduct(product.Name, 0, 1)
	}

	state.TempBool = false
	for state.Money > 10 {
		for !state.TempBool {
			index := fmt.Sprintf("%d", rand.Intn(products.MapsiAvailableProducts.Len())+1)

			productName := correctProductName(index)
			product := products.AvailableProducts[productName]
			aviableVolume := product.AviableVolume
			indexVolume := rand.Intn(len(aviableVolume))
			price := product.Prices[indexVolume]
			volume := aviableVolume[indexVolume]
			count := rand.Intn(int(state.Money/(20*price)+1)) + 1
			sumPrice := price * float64(count)

			existYet, _ := ProductExistYet(productName, volume)
			if state.Money/sumPrice < 5 ||
				(state.Money > 1 && state.Money-sumPrice < 0) ||
				(state.Money/sumPrice < 10 && existYet != nil) {
				continue
			}
			go buyTransaction(productName, count, volume, price, state.RandomBuy)
		}
	}
	state.RandomBuy = false
}
