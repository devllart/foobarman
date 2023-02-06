package drinks

import (
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/mapsi"
	"strings"
)

var ProductsStandartTypesPrice = map[string]float64{}

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
		AvailableProducts[name] = drink
	}

	for name, coctail := range AviableCoctail {
		delete(AviableCoctail, name)
		name = strings.Title(name)
		coctail.Name = name
		coctail.Price = coctail.GetPrice()
		AviableCoctail[name] = coctail
	}

	MapsiAvailableCoctail = mapsi.New(AviableCoctail)
	MapsiAvailableProducts = mapsi.New(AvailableProducts)
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
