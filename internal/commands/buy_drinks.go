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
		buyRandom()
	} else {
		volume := correctVolume(state.Arg1)
		count := correctCount(state.Arg2)
		if count != 0 {
			buyDrink(state.Command, volume, count)
		}
	}
}

func buyRandom() {

	for state.Money > 0 {
		index := fmt.Sprintf("%d", rand.Intn(len(drinks.AviableDrinks))+1)

		buyDrink(index, 0, 1)
	}
}

func buyDrink(drinkNameOrIndex string, volume float64, count int) {
	drinkName := correctDrinkName(drinkNameOrIndex)

	drink, exist := drinks.AviableDrinks[drinkName]
	if exist == false {
		alert.NotAvailableDrink(drinkName)
	}

	// If user not inputed volume or input "0" then assign minimal avliable value (firt in slice)
	if volume == 0 {
		volume = drink.AviableVolume[0]
	}

	index := funcs.IndexOf(volume, drink.AviableVolume)
	if index == -1 {
		index = 0
	}

	sumPrice := float64(count) * drink.Prices[index]
	if state.Money-sumPrice < 0 {
		alert.NotEnoughtFundsToBuy(sumPrice)
	}
	if funcs.Contains(drink.AviableVolume, volume) == false {
		alert.NotVolumeOfDrink(drinkName, volume)
	}

	state.Money -= sumPrice

	for i := range state.Bar {
		drink := state.Bar[i]
		if drink.Name == drinkName && drink.Volume == volume {
			state.Bar[i].Count += count
			alert.DrinkBoughtYet(drinkName, volume, count, state.Bar[i].Count, sumPrice)
			return
		}
	}
	state.Bar = append(state.Bar, drinks.New(drinkName, volume, count))
	alert.DrinkBought(drinkName, volume, count, sumPrice)
}
