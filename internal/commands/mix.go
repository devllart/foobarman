package commands

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/drinks"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
	"fmt"
	"sort"
	"strconv"
)

func correctIndexDrinkName(drinkName string) int {
	if index, err := strconv.Atoi(drinkName); err == nil {
		if index > len(state.DrinksIds) {
			return -1
		}
		return index - 1
	}

	for i := range state.DrinksIds {
		if drinkName == state.DrinksIds[i] {
			return i
		}
	}

	return -1
}

func mix() {
	state.Command = ""
	recipes := []string{}
	fmtc.Printf(texts.SelectIngredients)
	state.Mix = true
	for !CommandIs("mix") {
		if CommandIs("exit") {
			state.Run = false
			return
		} else if CommandIs("desc") {
			config.ShowDescription = !config.ShowDescription
		}

		state.Command = ""
		fmt.Print(" | ")
		fmt.Scanln(&state.Command)
		index := correctIndexDrinkName(state.Command)
		if index != -1 {
			drink := state.Bar[index]
			recipes = append(recipes, drink.Type)
			if err := (&state.Bar[index]).SubVolume(); err != nil {
				fmtc.Printf("%s\n", err)
			} else {
				fmtc.Printf("%Y+ %B%s%C\n", drink.Name)
			}
		} else {
			fmtc.Printf("%RУ вас нет такого напитка%C\n")
		}
	}

	for name, coctail := range drinks.AviableCoctail {
		sort.Sort(sort.StringSlice(recipes))
		if slicesEqual(coctail.Ingredients, recipes) {
			alert.CoctailIsReady(name)
			return
		}
	}
	alert.DontTheRecipies()
	state.Mix = false
}

func slicesEqual(firstSlice, secondSlice []string) bool {
	if len(firstSlice) != len(secondSlice) {
		return false
	}
	for i, val := range firstSlice {
		if secondSlice[i] != val {
			return false
		}
	}
	return true 


}
