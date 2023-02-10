package commands

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/products"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"
	"fmt"
	"math/rand"
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
		fmt.Println(coctail.Ingredients)
		fmt.Println(coctail.Grammar)

		// sort.Sort(sort.StringSlice(recipes))
		// ingredients := []string{}
		// ingredients = append(ingredients, coctail.Ingredients...)
		// sort.Sort(sort.StringSlice(ingredients))
		if funcs.SlicesSortAndEqual(coctail.Ingredients, recipes) {
			for i, vol := range coctail.Grammar {

				for _, indexOfBar := range barIndexes {
					if coctail.Ingredients[i] == state.Bar[indexOfBar].Type {
						if err := state.Bar[indexOfBar].SubVolume(vol * 0.001); err != nil {
							alert.Error(err.Error())
							return
						}
						product := state.Bar[indexOfBar]
						alert.VolumeOfProductSpent(vol, coctail.Units[i], product.Name)
					}
				}

			}
			state.YourCoctail = *coctail
			state.CoctailReady = true
			alert.CoctailIsReady(name)
			return
		}
	}

	for _, index := range barIndexes {
		vol := float64(10+rand.Intn(20)) * 0.01
		if err := state.Bar[index].SubVolume(vol); err != nil {
			alert.Error(err.Error())
			return
		}
		product := state.Bar[index]
		alert.VolumeOfProductSpent(vol, product.TypeVolume(), product.Name)
	}
	alert.DontTheRecipies()
	state.Mix = false
}
