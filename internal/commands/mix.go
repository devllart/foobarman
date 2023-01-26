package commands

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/drinks"
	"devllart/foobarman/internal/ptf"
	"devllart/foobarman/internal/state"
	"fmt"
	"sort"

	"golang.org/x/exp/slices"
)

func Mix() {
	recipes := []string{}
	state.Command = ""
	ptf.SelectIngredients()
	for !CommandIs("mix") {
		fmt.Print(" | ")
		fmt.Scanf("%s", &state.Command)
		_, exist := drinks.AviableDrinks[state.Command]
		if exist {
			recipes = append(recipes, state.Command)
		}
	}

	for name, coctail := range drinks.AviableCoctail {
		sort.Sort(sort.StringSlice(recipes))
		if slices.Equal(coctail.Ingredients, recipes) {
			alert.CoctailIsReady(name)
			return
		}
		alert.DontTheRecipies()
	}

}
