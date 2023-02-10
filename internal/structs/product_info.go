package structs

import (
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/data"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"
	"fmt"
	"strings"
)

/**
 * Funcs for struct ProductInfo
 */

func (product *ProductInfo) PrettyDescription() {
	if config.ShowDescription {
		if *product.Description == "" {
			fmtc.Printf("%sО %s ничего неизвестно\n", funcs.Indent(15), product.Name)
			return
		}
		lines := strings.Split(*product.Description, "\n")

		for _, line := range lines {
			fmtc.Printf("%s%s\n", funcs.Indent(15), line)
		}
		fmt.Println()
	}
}

// Calculate type of volume of ProductInfo
func (product *ProductInfo) TypeVolume() string {
	if typeVolume, exist := data.ProductsTypesVolume[product.Type]; exist == true {
		return typeVolume
	}

	return ".л"
}

// Print available prices (depending of volume) of ProductInfo
func (product *ProductInfo) PrintPrices() {
	for i := range product.Prices {
		price := product.Prices[i]
		fmtc.Printf(texts.StoreProductInfoPrice, price, product.AviableVolume[i], product.TypeVolume())
	}
	fmtc.Printf(" |\n\n")
}

func (product *ProductInfo) PrintInStore(index int) {
	fmtc.Printf(texts.StoreProductInfo, index, product.Name, product.Alc, product.Type, *product.Taste)
}
