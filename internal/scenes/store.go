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

	for name, drink := range drinks.AviableDrinks {
		fmt.Printf("%s (%.2f процентов алкоголя) | %s (вкус %s)\n", name, drink.Alc, drink.Type, drink.Taste)
		drink.PrettyDescription()
		for i, price := range drink.Prices {
			fmt.Printf("| %.3f$ за %.3f .л ", price, drink.AviableVolume[i])
		}
		fmt.Print(" |\n\n")
	}
}
