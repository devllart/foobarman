package commands

import (
	"devllart/foobarman/internal/products"
	"devllart/foobarman/internal/scenes"
	"devllart/foobarman/internal/state"
	"fmt"
	"math/rand"
)

func buy() {
	if CommandIs("ok") {
		state.Scene = scenes.Bar
		return
	} else if CommandIs("rand") {
		state.RandomBuy = true
		buyRandom()
		state.RandomBuy = false
	} else {
		volume := correctVolume(state.Args[0])
		count := correctCount(state.Args[1])
		if count != 0 {
			buyProduct(state.Command, volume, count)
		}
	}
}

func buyRandom() {
	for _, drink := range drinks.AvailableProducts {
		buyProduct(drink.Name, 0, 1)
	}

	for state.Money > 10 {
		state.TempBool = false
		for !state.TempBool {
			index := fmt.Sprintf("%d", rand.Intn(len(drinks.AvailableProducts))+1)

			drinkName := correctProductName(index)
			drink := drinks.AvailableProducts[drinkName]
			aviableVolume := drink.AviableVolume
			indexVolume := rand.Intn(len(aviableVolume))
			price := drink.Prices[indexVolume]
			volume := aviableVolume[indexVolume]
			count := rand.Intn(int(state.Money/(20*price)+1)) + 1
			sumPrice := price * float64(count)

			existYet, _ := ProductExistYet(drinkName, volume)
			if state.Money/sumPrice < 5 ||
				(state.Money > 1 && state.Money-sumPrice < 0) ||
				(state.Money/sumPrice < 10 && existYet != nil) {
				continue
			}
			go buyTransaction(drinkName, count, volume, price, state.RandomBuy)
		}
	}
}
