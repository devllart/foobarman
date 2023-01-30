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
	} else {
		volume := correctVolume(state.Args[0])
		count := correctCount(state.Args[1])
		if count != 0 {
			buyDrink(state.Command, volume, count)
		}
	}
	if CommandIs("rand") {
		state.Scene = scenes.Bar
	}
	//  else {
	// 	alert.Show()
	// }
}

func buyRandom() {
	for state.Money > 10 {
		index := fmt.Sprintf("%d", rand.Intn(len(drinks.AviableDrinks))+1)

		drinkName := correctDrinkName(index)
		drink := drinks.AviableDrinks[drinkName]
		aviableVolume := drink.AviableVolume
		indexVolume := rand.Intn(len(aviableVolume))

		count := rand.Intn(int(state.Money/(10*drink.Prices[indexVolume])+1)) + 1
		buyDrink(index, aviableVolume[indexVolume], count)
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

	index := funcs.IndexOf(volume, drink.AviableVolume)
	if index == -1 {
		index = 0
	}

	// Calculate the total sum
	// If selected volume not exist, alert user and panic "IncorrectInput"
	if funcs.Contains(drink.AviableVolume, volume) == false {
		alert.PanicNotVolumeOfDrink(drinkName, volume)
	}

	// Buy transaction
	// Buy the first drink (not in the bar yet)
	// state.Bar = append(state.Bar, newDrink)
	go bbuy(drinkName, count, volume, drink.Prices[index])
}

func bbuy(drinkName string, count int, volume, price float64) {
	sumPrice := float64(count) * price

	if state.Money-sumPrice < 0 {
		// alert.NotEnoughtFundsToBuy(sumPrice)
		state.TempBool = true
		return
	}

	state.Money -= sumPrice

	newDrink := drinks.New(drinkName, volume, count)

	// Cycle cycle through the bar's drinks list
	for i := range state.Bar {
		drink := state.Bar[i]

		// If the drink has in the bar yet, then added to exist drink
		if drink.Name == drinkName && drink.Volume == volume {
			state.Bar[i].Count += count
			alert.DrinkBoughtYet(drinkName, newDrink.TypeVolume(), volume, sumPrice, count, state.Bar[i].Count)
			return
		}
	}
	state.Bar = append(state.Bar, newDrink)
	alert.DrinkBought(drinkName, newDrink.TypeVolume(), volume, sumPrice, count)
}
