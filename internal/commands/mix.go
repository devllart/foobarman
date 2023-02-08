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
		fmt.Println(coctail.Ingredients)
		fmt.Println(coctail.Grammar)

		sort.Sort(sort.StringSlice(recipes))
		ingredients := []string{}
		ingredients = append(ingredients, coctail.Ingredients...)
		sort.Sort(sort.StringSlice(ingredients))
		if funcs.SlicesEqual(ingredients, recipes) {
			for i, vol := range coctail.Grammar {

				for _, indexOfBar := range barIndexes {
					if coctail.Ingredients[i] == state.Bar[indexOfBar].Type {
						if err := state.Bar[indexOfBar].SubVolume(vol * 0.001); err != nil {
							state.AddInfof(err.Error())
							return
						}
						product := state.Bar[indexOfBar]
						state.AddInfof("%R - %.3f%s %B%s%C (для %B%s%C) %C\n", vol, coctail.Units[i], product.Name, coctail.Name)
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
			state.AddInfof(err.Error())
			return
		}
		product := state.Bar[index]
		state.AddInfof("%R - %.3f%s %B%s %R(в пустую) %C\n", vol, product.TypeVolume(), product.Name)
	}
	alert.DontTheRecipies()
	state.Mix = false
}
