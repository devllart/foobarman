package commands

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/drinks"
	"devllart/foobarman/internal/scenes"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/src/funcs"
	"strconv"
)

func Buy() {
	if CommandIs("ok") {
		state.Scene = scenes.Bar
		return
	} else {
		volume := CorrectVolume(state.Arg1)
		count := CorrectCount(state.Arg2)
		if count != 0 {
			BayDrink(state.Command, volume, count)
		}
	}
}

/**
 * Get Name or Index drink's and return correct its name.
 * + alerting user if his index out of range drink's menu.
 */
func CorrectDrinkName(drinkNameOrIndex string) string {
	drinkIndex, err := strconv.Atoi(drinkNameOrIndex)
	if err == nil {
		if drinkIndex > len(state.DrinksIds) {
			alert.NotAvailableIndexDrink(drinkIndex)
		}
		return state.DrinksIds[drinkIndex-1]
	}

	return drinkNameOrIndex
}

func BayDrink(drinkNameOrIndex string, volume float64, count int) {
	drinkName := CorrectDrinkName(drinkNameOrIndex)

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

func CorrectCount(arg string) int {
	var count int = 1
	var err error
	if arg != "" {
		count, err = strconv.Atoi(arg)
	}
	if err != nil || count <= 0 {
		alert.IncorrectAmountOfDrink()
		return 0
	}

	return count
}

func CorrectVolume(arg string) float64 {
	var volume float64 = 0
	var err error
	if arg != "" {
		volume, err = strconv.ParseFloat(arg, 64)
	}
	if err != nil || volume <= 0 {
		alert.IncorrectVolumeOfDrink()
		return 0
	}

	return volume
}
