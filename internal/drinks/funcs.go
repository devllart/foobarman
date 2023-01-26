package drinks

import (
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/errors"
	"devllart/foobarman/internal/ptf"
	"devllart/foobarman/src/funcs"
	"fmt"
	"strings"
)

func New(name string, volume float64, count int) Drink {
	info, exitst := AviableDrinks[name]
	if exitst == false {
		errors.NotExistDrink(name)
	}
	if funcs.Contains(info.AviableVolume, volume) == false {
		errors.NoVolumeOfDrink(name, volume)
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
	ptf.DrinkShow(drink.Name, drink.Volume, drink.Count, drink.LastVolume)
	drink.Info.PrettyDescription()
}
