package commands

import (
	"devllart/foobarman/internal/drinks"
	"devllart/foobarman/internal/state"
	"fmt"
	"sort"

	"golang.org/x/exp/slices"
)

func Mix() {
	recipes := []string{}
	state.Command = ""
	fmt.Print("Из каких ингридиентов будет ваш коктель?\n")
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
			state.Info += fmt.Sprintf("У вас поличился %s\n", name)
			return
		}
	}
	state.Info += fmt.Sprint("Чтож жаль, но такого рецепта нет\n")
}
