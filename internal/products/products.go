package products

import (
	"devllart/foobarman/internal/config"
	"devllart/foobarman/src/funcs"
)

func init() {
	if config.Env == "generate" {
		return
	} else if config.Env != "production" {
		funcs.ParseJsonToStruct("data/coctails.json", &AllCoctail)
		funcs.ParseJsonToStruct("data/products.json", &AvailableProducts)
	}

	for _, product := range AvailableProducts {
		// product.Taste = getTaste(product)
		if _, exist := ProductsStandartTypesPrice[product.Type]; exist {
			ProductsStandartTypesPrice[product.Type] += product.Prices[0] * product.AviableVolume[0] * 0.5
		} else {
			ProductsStandartTypesPrice[product.Type] = product.Prices[0] * product.AviableVolume[0]
		}

		AvailableIngredients = append(AvailableIngredients, product.Type)
		// AvailableProducts[name] = product
	}

	AddAvailablesCoctail(5)

}
