package drinks

import (
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/errors"
	"devllart/foobarman/internal/ptf"
	"devllart/foobarman/src/funcs"
	"fmt"
	"strings"

	"github.com/TwiN/go-color"
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
	// fmt.Printf(texts.ShowDrinkInBar, color.Blue, drink.Name, color.Reset, color.Red, drink.Info.Alc, color.Reset, color.Green, drink.Volume, color.Reset, color.Yellow, drink.Count, color.Reset, color.Green, drink.LastVolume, color.Reset)
	ptf.DrinkShow(drink.Name, drink.Count, drink.Info.Alc, drink.Volume, drink.LastVolume)
	drink.Info.PrettyDescription()
}

func (drink DrinkInfo) Valid() {

	if drink.Name != "Содовая" {
		for _, price := range drink.Prices {
			if price < 1 {
				panic(fmt.Sprintf("Для %s%s%s %s%.3f$%s это слишком дешёво", color.Red, drink.Name, color.Reset, color.Yellow, price, color.Reset))
			}
		}
	}
}
