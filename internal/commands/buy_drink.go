package commands

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/drinks"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/src/funcs"
)

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