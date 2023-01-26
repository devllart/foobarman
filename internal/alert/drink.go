package alert

import "devllart/foobarman/internal/state"

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

func DrinkBought(drinkName string, volume float64, count int, sumPrice float64) {
	state.AddInfof("+ %s %.3f .л (количество: %d) куплено (-%.2f $)\n", drinkName, volume, count, sumPrice)
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
