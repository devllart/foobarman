package alert

import (
	"devllart/foobarman/internal/state"

	"github.com/TwiN/go-color"
)

/**
 * With panic!
 */

func NotAvailableDrink(drinkName string) {
	state.AddInfof("! Напитка %s нет продаже в продаже\n", drinkName)
	panic("UncorrectInput")
}

func NotAvailableIndexDrink(drinkIndex int) {
	state.AddInfof("! Напитка под номером %d нет в продаже\n", drinkIndex)
	panic("UncorrectInput")
}

func NotEnoughtFundsToBuy(sumPrice float64) {
	state.AddInfof("! Недостаточно средст для покупки (общая сумма составила %.2f $)\n", sumPrice)
	panic("UncorrectInput")
}

func NotVolumeOfDrink(drinkName string, volume float64) {
	state.AddInfof("! %s с объёмом %.3f .л нет в продаже, возьмите другой объём\n", drinkName, volume)
	panic("UncorrectInput")
}

/**
 * Don't panic
 */

func IncorrectAmountOfDrink() {
	state.AddInfo("Неверно указанно количество напитка\n")
}

func IncorrectVolumeOfDrink() {
	state.AddInfo("Неверно указан объём напитка\n")
}

func DrinkBought(drinkName string, volume float64, count int, sumPrice float64) {
	state.AddInfof("%s+%s %s%s%s %s%.3f .л%s %s%dX%s куплено (%s-%.2f $%s)\n", color.Yellow, color.Reset, color.Red, drinkName, color.Reset, color.Green, volume, color.Reset, color.Yellow, count, color.Reset, color.Yellow, sumPrice, color.Reset)
}

func DrinkBoughtYet(drinkName string, volume float64, count, countSum int, sumPrice float64) {

	state.AddInfof("%s+%s %s%s%s %s%.3f .л%s %s%dX%s куплено ещё (общее количество: %s%d%s) (%s-%.2f $%s)\n", color.Yellow, color.Reset, color.Red, drinkName, color.Reset, color.Green, volume, color.Reset, color.Yellow, count, color.Reset, color.Green, countSum, color.Reset, color.Yellow, sumPrice, color.Reset)
	// state.AddInfof("+ %s %.3f .л (количество: %d) куплено ещё (общее колиство: %d) (-%.2f $)\n", drinkName, volume, count, countSum, sumPrice)
}
