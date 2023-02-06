package drinks

import (
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"
	"fmt"
	"strings"
)

func (drink *ProductInfo) PrettyDescription() {
	if config.ShowDescription {
		if drink.Description == "" {
			fmtc.Printf("%sО %s ничего неизвестно\n", funcs.Indent(15), drink.Name)
			return
		}

		lines := strings.Split(drink.Description, "\n")

		for _, line := range lines {
			fmtc.Printf("%s%s\n", funcs.Indent(15), line)
		}
		fmt.Println()
	}
}

func (drink *ProductInfo) TypeVolume() string {
	if typeVolume, exist := ProductsTypesVolume[drink.Type]; exist == true {
		return typeVolume
	}

	return ".л"
}

func (drink *ProductInfo) PrintPrices() {
	for i := range drink.Prices {
		price := drink.Prices[i]
		fmtc.Printf(texts.StoreProductInfoPrice, price, drink.AviableVolume[i], drink.TypeVolume())
	}
	fmtc.Printf(" |\n\n")
}

func (drink *ProductInfo) PrintInStore(index int) {
	fmtc.Printf(texts.StoreProductInfo, index, drink.Name, drink.Alc, drink.Type, *drink.Taste)
}
