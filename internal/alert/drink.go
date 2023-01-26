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
	panic("UncorrectInput")
}

func NotAvailableIndexDrink(drinkIndex int) {
	state.AddInfof(texts.NotAvailableIndexDrink, drinkIndex)
	panic("UncorrectInput")
}

func NotEnoughtFundsToBuy(sumPrice float64) {
	state.AddInfof(texts.NotEnoughFundsToBuy, sumPrice)
	panic("UncorrectInput")
}

func NotVolumeOfDrink(drinkName string, volume float64) {
	state.AddInfof(texts.NotVolumeOfDrink, drinkName, volume)
	panic("UncorrectInput")
}

/**
 * Don't panic
 */

func IncorrectAmountOfDrink() {
	state.AddInfo(texts.IncorrectAmountOfDrink)
}

func IncorrectVolumeOfDrink() {
	state.AddInfo(texts.IncorrectVolumeOfDrink)
}

func DrinkBought(drinkName string, volume float64, count int, sumPrice float64) {
	state.AddInfof("%Y+%C %B%s%C %G%.3f .л%C %Y%dX%C куплено (%Y-%.2f $%C)\n", drinkName, volume, count, sumPrice)
}

func DrinkBoughtYet(drinkName string, volume float64, count, countSum int, sumPrice float64) {

	state.AddInfof("%Y+%C %B%s%C %G%.3f .л%C %Y%dX%C куплено ещё (общее количество: %G%d%C) (%Y-%.2f $%C)\n", drinkName, volume, count, countSum, sumPrice)
}
