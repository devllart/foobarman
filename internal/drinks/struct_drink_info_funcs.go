package drinks

import (
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
	"strings"
)

func (drink DrinkInfo) PrettyDescription() {
	if config.ShowDescription {
		if drink.Description == "" {
			fmtc.Printf("      | О %s ничего неизвестно\n", drink.Name)
			return
		}

		lines := strings.Split(drink.Description, "\n")

		for _, line := range lines {
			fmtc.Printf("      | %s\n", line)
		}
	}
}

func (drink DrinkInfo) GetTaste() string {

	if taste, exist := Tastes[drink.Name]; exist == true {
		return taste
	}
	if taste, exist := Tastes[drink.Type]; exist == true {
		return taste
	}

	return texts.Unknown
}

func (drink DrinkInfo) GetTypeVolume() string {
	if typeVolume, exist := DrinksTypesVolume[drink.Type]; exist == true {
		return typeVolume
	}

	return ".л"
}

func (drink DrinkInfo) PrintPrices() {
	for i := range drink.Prices {
		price := drink.Prices[i]
		fmtc.Printf(texts.StoreDrinkInfoPrice, price, drink.AviableVolume[i], drink.GetTypeVolume())
	}
	fmtc.Printf(" |\n\n")
}

func (drink DrinkInfo) PrintInStore(index int) {
	fmtc.Printf(texts.StoreDrinkInfo, index, drink.Name, drink.Alc, drink.Type, drink.GetTaste())
}
