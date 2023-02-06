package commands

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/products"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"
	"fmt"
	"sort"
	"strconv"
)

func correctIndexProductName(drinkName string) int {
	if index, err := strconv.Atoi(drinkName); err == nil {
		if index > len(state.ProductsIds) {
			return -1
		}
		return index - 1
	}

	for i := range state.ProductsIds {
		if drinkName == state.ProductsIds[i] {
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
		if state.Command == "mix" {
			break
		}
		index := correctIndexProductName(state.Command)
		if index != -1 {
			drink := state.Bar[index]
			recipes = append(recipes, drink.Type)
			// if err := (&state.Bar[index]).SubVolume(); err != nil {
			// 	fmtc.Printf("%s\n", err)
			// } else {
			fmtc.Printf("%Y+ %B%s%C\n", drink.Name)
			// }
		} else {
			fmtc.Printf("%RУ вас нет такого напитка%C\n")
		}
	}

	for name, coctail := range drinks.MapsiAvailableCoctail.Data() {
		sort.Sort(sort.StringSlice(recipes))
		ingredients := coctail.Ingredients
		sort.Sort(sort.StringSlice(ingredients))
		if funcs.SlicesEqual(ingredients, recipes) {
			state.YourCoctail = *coctail
			state.CoctailReady = true
			alert.CoctailIsReady(name)
			return
		}
		// fmt.Println(recipes)
		// fmt.Println(ingredients)
		// panic("aaa")
	}
	alert.DontTheRecipies()
	state.Mix = false
}
