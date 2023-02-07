package products

import (
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/funcs"
	"devllart/foobarman/src/mapsi"
	"strings"
)

var ProductsStandartTypesPrice = map[string]float64{}

var AvailableIngredients = []string{}
var MapsiAvailableCoctail mapsi.Mapsi[Coctail]
var MapsiAvailableProducts mapsi.Mapsi[ProductInfo]

func init() {
	for name, drink := range AvailableProducts {
		delete(AvailableProducts, name)
		name = strings.Title(name)
		drink.Name = name
		drink.Taste = getTaste(drink)
		if _, exist := ProductsStandartTypesPrice[drink.Type]; exist {
			ProductsStandartTypesPrice[drink.Type] += drink.Prices[0] * drink.AviableVolume[0] * 0.5
		} else {
			ProductsStandartTypesPrice[drink.Type] = drink.Prices[0] * drink.AviableVolume[0]
		}

		AvailableIngredients = append(AvailableIngredients, drink.Type)
		AvailableProducts[name] = drink
	}

	MapsiAvailableProducts = mapsi.New(AvailableProducts)
	AddAvailablesCoctail(50)
}

func getTaste(drink ProductInfo) *string {
	if taste := NewTastes.GetValue(drink.Name); taste != nil {
		return taste
	}

	if taste := NewTastes.GetValue(drink.Type); taste != nil {
		return taste
	}

	return NewTastes.GetValue(texts.Unknown)
}

func AddAvailablesCoctail(count int) {
	for _, coctail := range AllCoctail {
		if AddAvailableCoctail(coctail) {
			count -= 1
		}
		if count <= 0 {
			return
		}
	}
}

func AddAvailableCoctail(coctail Coctail) bool {
	for _, ingredient := range coctail.Ingredients {
		if !funcs.Contains(AvailableIngredients, ingredient) {
			return false
		}
	}

	MapsiAvailableCoctail.SetValue(coctail.Name, coctail)
	return true
}
