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
		volume := correctVolume(state.Args[0])
		count := correctCount(state.Args[1])
		if count != 0 {
			buyDrink(state.Command, volume, count)
		}
	}
}

func buyRandom() {
	for state.Money > 0 {
		index := fmt.Sprintf("%d", rand.Intn(len(drinks.AviableDrinks))+1)

		drinkName := correctDrinkName(index)
		aviableVolume := drinks.AviableDrinks[drinkName].AviableVolume
		// if len(aviableVolume) != 0 {
		indexVolume := rand.Intn(len(aviableVolume))
		// }
		buyDrink(index, aviableVolume[indexVolume], rand.Intn(2)+1)
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
	sumPrice := float64(count) * drink.Prices[index]
	// If money not enought, alert user and panic "IncorrectInput"
	if state.Money-sumPrice < 0 {
		alert.PanicNotEnoughtFundsToBuy(sumPrice)
	}
	// If selected volume not exist, alert user and panic "IncorrectInput"
	if funcs.Contains(drink.AviableVolume, volume) == false {
		alert.PanicNotVolumeOfDrink(drinkName, volume)
	}

	// Buy transaction
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
	// Buy the first drink (not in the bar yet)
	state.Bar = append(state.Bar, newDrink)
	alert.DrinkBought(drinkName, newDrink.TypeVolume(), volume, sumPrice, count)
}
