package commands

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/drinks"
	"devllart/foobarman/internal/scenes"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/src/funcs"
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
			buyDrink(state.Command, volume, count)
		}
	}
}

func buyRandom() {
	for state.Money > 10 {
		state.TempBool = false
		for !state.TempBool {
			index := fmt.Sprintf("%d", rand.Intn(len(drinks.AviableDrinks))+1)

			drinkName := correctDrinkName(index)
			drink := drinks.AviableDrinks[drinkName]
			aviableVolume := drink.AviableVolume
			indexVolume := rand.Intn(len(aviableVolume))
			price := drink.Prices[indexVolume]
			volume := aviableVolume[indexVolume]
			count := rand.Intn(int(state.Money/(20*price)+1)) + 1
			sumPrice := price * float64(count)

			totalSumMoreOneFifth := state.Money/sumPrice < 5
			MoneyMoreOneDollarButSumPriceIsMore := state.Money > 1 && state.Money-sumPrice < 0
			totalSumMoreOneTenthAndAlredyHaveDrink := state.Money/sumPrice < 10 && DrinkExistYet(drinkName, volume) != nil
			if totalSumMoreOneFifth || MoneyMoreOneDollarButSumPriceIsMore || totalSumMoreOneTenthAndAlredyHaveDrink {
				continue
			}
			go buyTransaction(drinkName, count, volume, price, state.RandomBuy)
		}
	}
}

func buyDrink(drinkNameOrIndex string, volume float64, count int) {
	drinkName := correctDrinkName(drinkNameOrIndex)

	drink, exist := drinks.AviableDrinks[drinkName]
	if exist == false {
		alert.PanicNotAvailableDrink(drinkName)
	}

	// If user not inputed volume or input "0" then assign minimal avliable value (firt in slice)
	if volume == 0 {
		volume = drink.AviableVolume[0]
	}

	index := funcs.IndexOfOrNull(volume, drink.AviableVolume)

	// Calculate the total sum
	// If selected volume not exist, alert user and panic "IncorrectInput"
	if funcs.Contains(drink.AviableVolume, volume) == false {
		alert.PanicNotVolumeOfDrink(drinkName, volume)
	}

	// Buy drink transaction in goroutine (Oh yes I'm a pervert)
	buyTransaction(drinkName, count, volume, drink.Prices[index], state.RandomBuy)
}

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
