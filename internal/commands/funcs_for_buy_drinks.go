package commands

import (
	"devllart/foobarman/internal/drinks"
	"devllart/foobarman/src/funcs"
	"devllart/foobarman/internal/state"
	"fmt"
	"strconv"
)

func CorrectCount(arg string) int {
	count := 1
	var err error
	if arg != "" {
		count, err = strconv.Atoi(arg)
	}
	if err != nil || count <= 0 {
		state.Info += "Неверно указанно количество напитка\n"
		return 0
	}

	return count
}

func CorrectVolume(arg string) float64 {
	var volume float64
	var err error
	if arg != "" {
		volume, err = strconv.ParseFloat(arg, 64)
	}
	if err != nil || volume <= 0 {
		state.Info += "Неверно указан объём напитка\n"
		return 0
	}

	return volume
}

func BayDrink(drinkName string, volume float64, count int) {
	drink, exist := drinks.AviableDrinks[drinkName]
	if exist == false {
		state.Info += fmt.Sprintf("! Напитка %s нет продаже\n", drinkName)
		return
	}
	sumPrice := float64(count) * drink.Prices[0]
	if state.Money-sumPrice < 0 {
		state.Info += fmt.Sprintf("! Недостаточно средст для покупки (общая сумма составила %.2f $)\n", sumPrice)
		return
	}
	if funcs.Contains(drink.AviableVolume, volume) == false {
		state.Info += fmt.Sprintf("%s с объёмом %.3f нет в продаже, возьмите другой объём\n", drinkName, volume)
		return
	}

	state.Bar = append(state.Bar, drinks.New(drinkName, volume, count))
	state.Money -= sumPrice
	state.Info += fmt.Sprintf("%s (%d) куплено (-%.2f $)\n", drinkName, count, sumPrice)
}
