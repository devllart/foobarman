package drinks

import (
	"devllart/foobarman/internal/config"
	"devllart/foobarman/src/funcs"
	"fmt"
	"strings"
)

func New(name string, volume float64, count int) Drink {
	info, exitst := AviableDrinks[name]
	if exitst == false {
		err := fmt.Errorf("Напитка %s не существует", name)
		panic(err)
	}
	if funcs.Contains(info.AviableVolume, volume) == false {
		err := fmt.Errorf("%s с объёмом %g не существует", name, volume)
		panic(err)

	}

	return Drink{name, volume, count, volume, info}
}

func (drink DrinkInfo) PrettyDescription() {
	if config.ShowDescription {
		lines := strings.Split(drink.Description, "\n")

		for _, line := range lines {
			fmt.Printf("      | %s\n", line)
		}
	}
}

func (drink Drink) Show() {
	fmt.Printf("   %s (%g .л) %dX | в последней бутылке осталось %g .л\n", drink.Name, drink.Volume, drink.Count, drink.LastVolume)
	drink.Info.PrettyDescription()
}
