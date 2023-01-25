package drinks

import (
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/funcs"
	"fmt"
	"strings"
)

type Drink struct {
	Name       string
	Volume     float64
	Count      int
	LastVolume float64
	Info       DrinkInfo
}

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

func (drink Drink) Show() {
	fmt.Printf("   %s (%g .л) %dX | в последней бутылке осталось %g .л\n", drink.Name, drink.Volume, drink.Count, drink.LastVolume)
	if config.ShowDescription {
		drink.PrettyDescription()
	}
}

func (drink Drink) PrettyDescription() {
	lines := strings.Split(drink.Info.Description, "\n")

	for _, line := range lines {
		fmt.Printf("      | %s\n", line)
	}
}
