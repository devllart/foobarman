package drinks

import (
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"
	"strings"
)

func New(name string, volume float64, count int) Drink {
	info, exitst := AviableDrinks[name]
	if exitst == false {
		fmtc.Perrorf(texts.NotExistDrinkError, name)
	}
	if funcs.Contains(info.AviableVolume, volume) == false {
		fmtc.Perrorf(texts.NotVolumeOfDrinkError, name, volume)
	}

	return Drink{name, volume, count, volume, info}
}

func (drink DrinkInfo) PrettyDescription() {
	if config.ShowDescription {
		lines := strings.Split(drink.Description, "\n")

		for _, line := range lines {
			fmtc.Printf("      | %s\n", line)
		}
	}
}

func (drink Drink) Show() {
	fmtc.Printf(texts.ShowDrinkInBar, drink.Name, drink.Info.Alc, drink.Volume, drink.Count, drink.LastVolume)
	drink.Info.PrettyDescription()
}

func (drink DrinkInfo) Valid() {
	if drink.Name != "Содовая" {
		for _, price := range drink.Prices {
			if price < 1 {
				panic(fmtc.Sprintf("Для %R%s%C %Y%.3f$%C это слишком дешёво", drink.Name, price))
			}
		}
	}
}
