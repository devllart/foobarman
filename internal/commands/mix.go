package commands

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/products"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"
	"fmt"
	"sort"
)

func mix() {
	state.Command = ""
	recipes := []string{}
	barIndexes := []int{}
	fmtc.Printf(texts.SelectIngredients)

	state.Mix = true
	for !CommandIs("mix") {
		state.Command = ""
		fmt.Print(" | ")
		fmt.Scanln(&state.Command)

		if state.Command == "mix" {
			break
		} else if CommandIs("exit") {
			state.Run = false
			return
		}

		index := correctIndexProductName(state.Command)
		if index != -1 {
			product := state.Bar[index]
			recipes = append(recipes, product.Type)
			barIndexes = append(barIndexes, index)
			fmtc.Printf("%Y+ %B%s%C\n", product.Name)
		} else {
			fmtc.Printf("%RУ вас нет такого напитка%C\n")
		}
	}

	for name, coctail := range products.MapsiAvailableCoctail.Data() {
		sort.Sort(sort.StringSlice(recipes))
		ingredients := coctail.Ingredients
		sort.Sort(sort.StringSlice(ingredients))
		if funcs.SlicesEqual(ingredients, recipes) {
			for i, vol := range coctail.Grammar {
				if err := state.Bar[barIndexes[i]].SubVolume(vol * 0.001); err != nil {
					state.AddInfof(err.Error())
					return
				}
			}
			state.YourCoctail = *coctail
			state.CoctailReady = true
			alert.CoctailIsReady(name)
			return
		}
	}

	for _, index := range barIndexes {
		if err := state.Bar[index].SubVolume(0.1); err != nil {
			state.AddInfof(err.Error())
			return
		}
	}
	alert.DontTheRecipies()
	state.Mix = false
}
