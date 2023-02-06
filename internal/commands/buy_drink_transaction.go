package commands

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/products"
	"devllart/foobarman/internal/state"
)

func buyTransaction(drinkName string, count int, volume, price float64, rand bool) {
	if state.TempBool {
		return
	}

	sumPrice := float64(count) * price

	if state.Money-sumPrice < 0 {
		if !rand {
			alert.NotEnoughtFundsToBuy(sumPrice)
		}
		state.TempBool = true
		return
	}

	state.Money -= sumPrice

	newDrink := drinks.New(drinkName, volume, count)

	// Buy the not first drink (exitst in the bar yet)
	if drink := DrinkExistYet(newDrink.Name, newDrink.Volume); drink != nil {
		drink.Count += count
		alert.DrinkBoughtYet(drinkName, newDrink.TypeVolume(), volume, sumPrice, count, drink.Count)
		return
	}

	// Buy the first drink (not in the bar yet)
	state.Bar = append(state.Bar, newDrink)
	alert.DrinkBought(drinkName, newDrink.TypeVolume(), volume, sumPrice, count)
}
