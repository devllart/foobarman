package alert

import (
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
)

/**
 * With panic!
 */

func NotAvailableDrink(drinkName string) {
	state.AddInfof(texts.NotAvailableDrink, drinkName)
	panic("IncorrectInput")
}

func NotAvailableIndexDrink(drinkIndex int) {
	state.AddInfof(texts.NotAvailableIndexDrink, drinkIndex)
	panic("IncorrectInput")
}

func NotEnoughtFundsToBuy(sumPrice float64) {
	if state.Command != "rand" {
		state.AddInfof(texts.NotEnoughFundsToBuy, sumPrice)
	}
	panic("IncorrectInput")
}

func NotVolumeOfDrink(drinkName string, volume float64) {
	state.AddInfof(texts.NotVolumeOfDrink, drinkName, volume)
	panic("IncorrectInput")
}

/**
 * Don't panic
 */

func IncorrectAmountOfDrink() {
	state.AddInfof(texts.IncorrectAmountOfDrink)
}

func IncorrectVolumeOfDrink() {
	state.AddInfof(texts.IncorrectVolumeOfDrink)
}

func DrinkBought(drinkName string, volume float64, count int, sumPrice float64) {
	state.AddInfof("%Y+%C %B%s%C %G%.3f .л%C %Y%dX%C куплено (%Y-%.2f $%C)\n", drinkName, volume, count, sumPrice)
}

func DrinkBoughtYet(drinkName string, volume float64, count, countSum int, sumPrice float64) {

	state.AddInfof("%Y+%C %B%s%C %G%.3f .л%C %Y%dX%C куплено ещё (общее количество: %G%d%C) (%Y-%.2f $%C)\n", drinkName, volume, count, countSum, sumPrice)
}
