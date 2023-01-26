package scenes

import (
	"devllart/foobarman/internal/drinks"
	"devllart/foobarman/internal/state"
	"fmt"
)

func Store() {
	fmt.Printf("Добро пожаловаь в магазин Brizly названый в честь Бри Ларсон.\n%s выбирай, что хочешь.\n\n", state.Name)
	state.Info += `
Подсказка: Напиши {Название напитка} {Его объём} {Количество (по умолчанию 1)}, чтобы купить.
Например "> Просекко 0.75 3"
`

	i := 1
	for name, drink := range drinks.AviableDrinks {
		fmt.Printf("%d. %s (%.2f процентов алкоголя) | %s (вкус %s)\n", i, name, drink.Alc, drink.Type, drink.Taste)
		drink.PrettyDescription()

		state.DrinksIds = append(state.DrinksIds, name)
		for i := range drink.Prices {
			price := drink.Prices[i]
			fmt.Printf("| %.3f$ за %.3f .л ", price, drink.AviableVolume[i])
		}
		fmt.Print(" |\n\n")
		i += 1
	}
}
