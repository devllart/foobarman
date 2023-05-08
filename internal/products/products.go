package products

import (
	anames "devllart/foobarman/internal/a_names"
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/data"
	"devllart/foobarman/internal/structs"
	"devllart/foobarman/src/funcs"
	"strings"
)

func init() {
	if config.Env == "generate" {
		return
	} else if config.Env != "production" {
		useJsonData()
	}

	for _, product := range AvailableProducts {
		var coefficient float32 = 1

		if _, exist := data.ProductsStandartTypesPrice[product.Type]; exist {
			coefficient = 0.5
		}

		data.ProductsStandartTypesPrice[product.Type] = product.Prices[0] * product.AviableVolume[0] * float64(coefficient)
		AvailableIngredients = append(AvailableIngredients, product.Type)
	}

	AddAvailablesCoctail(3)

	// if state.LoadSave {
	// panic(state.Stage)
	// for i := 2; i <= state.Stage; i++ {
	// 	AddAvailablesCoctail(i)
	// }
	// }
}

func useJsonData() {
	products := []structs.ProductInfo{}

	funcs.ParseJsonToStruct(anames.JsonDataDir+"/coctails.json", &AllCoctail)
	funcs.ParseJsonToStruct(anames.JsonDataDir+"/products.json", &products)

	for _, product := range products {
		AvailableProducts[strings.Title(product.Name)] = product
	}
}
