package commands

import (
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

func BayDrink(drinkName string, volume float64, count int) {

	indexDrink, err := strconv.Atoi(drinkName)
	if err == nil {
		indexDrink -= 1
		if indexDrink >= len(state.DrinksIds) {
			state.AddInfof("! Напитка под номером %d нет\n", indexDrink)
			return
		}
		drinkName = state.DrinksIds[indexDrink]
	}

	drink, exist := drinks.AviableDrinks[drinkName]
	if exist == false {
		state.AddInfof("! Напитка %s нет продаже\n", drinkName)
		return
	}
	if volume == 0 {
		volume = drink.AviableVolume[0]
	}

	index := funcs.IndexOf(volume, drink.AviableVolume)
	if index == -1 {
		index = 0
	}
	sumPrice := float64(count) * drink.Prices[index]
	if state.Money-sumPrice < 0 {
		state.AddInfof("! Недостаточно средст для покупки (общая сумма составила %.2f $)\n", sumPrice)
		return
	}
	if funcs.Contains(drink.AviableVolume, volume) == false {
		state.AddInfof("! %s с объёмом %.3f .л нет в продаже, возьмите другой объём\n", drinkName, volume)
		return
	}

	state.Money -= sumPrice

	for i := range state.Bar {
		drink := state.Bar[i]
		if drink.Name == drinkName && drink.Volume == volume {
			state.Bar[i].Count += count
			return
		}
	}
	state.Bar = append(state.Bar, drinks.New(drinkName, volume, count))
	state.AddInfof("+ %s %.3f .л (количество: %d) куплено (-%.2f $)\n", drinkName, volume, count, sumPrice)
}

func CorrectCount(arg string) int {
	var count int = 1
	var err error
	if arg != "" {
		count, err = strconv.Atoi(arg)
	}
	if err != nil || count <= 0 {
		state.AddInfo("Неверно указанно количество напитка\n")
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
		state.AddInfo("Неверно указан объём напитка\n")
		return 0
	}

	return volume
}
