package drinks

import (
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"
	"fmt"
	"strings"
)

func (drink *DrinkInfo) PrettyDescription() {
	if config.ShowDescription {
		if drink.Description == "" {
			fmtc.Printf("%sО %s ничего неизвестно\n", funcs.Indent(4), drink.Name)
			return
		}

		lines := strings.Split(drink.Description, "\n")

		for _, line := range lines {
			fmtc.Printf("%s%s\n", funcs.Indent(4), line)
		}
		fmt.Println()
	}
}

func (drink *DrinkInfo) TypeVolume() string {
	if typeVolume, exist := DrinksTypesVolume[drink.Type]; exist == true {
		return typeVolume
	}

	return ".л"
}

func (drink *DrinkInfo) PrintPrices() {
	for i := range drink.Prices {
		price := drink.Prices[i]
		fmtc.Printf(texts.StoreDrinkInfoPrice, price, drink.AviableVolume[i], drink.TypeVolume())
	}
	fmtc.Printf(" |\n\n")
}

func (drink *DrinkInfo) PrintInStore(index int) {
	fmtc.Printf(texts.StoreDrinkInfo, index, drink.Name, drink.Alc, drink.Type, drink.Taste)
}
