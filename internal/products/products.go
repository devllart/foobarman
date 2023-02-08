package products

import (
	"devllart/foobarman/internal/config"
	"devllart/foobarman/src/funcs"
	"strings"
)

func init() {
	if config.Env == "generate" {
		return
	} else if config.Env != "production" {
		funcs.ParseJsonToStruct("data/coctails.json", &AllCoctail)
		products := []ProductInfo{}
		funcs.ParseJsonToStruct("data/products.json", &products)
		for _, product := range products {
			AvailableProducts[strings.Title(product.Name)] = product
		}
	}

	for _, product := range AvailableProducts {
		if _, exist := ProductsStandartTypesPrice[product.Type]; exist {
			ProductsStandartTypesPrice[product.Type] += product.Prices[0] * product.AviableVolume[0] * 0.5
		} else {
			ProductsStandartTypesPrice[product.Type] = product.Prices[0] * product.AviableVolume[0]
		}

		AvailableIngredients = append(AvailableIngredients, product.Type)
	}

	AddAvailablesCoctail(3)
}
