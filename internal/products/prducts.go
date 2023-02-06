package drinks

import (
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/mapsi"
	"encoding/json"
	"io/ioutil"
	"strings"
)

var ProductsStandartTypesPrice = map[string]float64{}

var MapsiAvailableCoctail mapsi.Mapsi[Coctail]
var MapsiAvailableProducts mapsi.Mapsi[ProductInfo]
var AllCoctail = map[string]Coctail{}
var CountAvailableCoctail = 20

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

	coctails := []Coctail{}
	file, err := ioutil.ReadFile("data/coctails.json")
	if err != nil {
		panic(err.Error())
	}

	if err := json.Unmarshal([]byte(file), &coctails); err != nil {
		panic(err.Error())
	}

	for _, coctail := range coctails {
		coctail.Name = strings.Title(coctail.Name)
		coctail.Price = coctail.GetPrice()
		AllCoctail[coctail.Name] = coctail
		// if i < CountAvailableCoctail {
		// 	AviableCoctail[coctail.Name] = coctail
		// }
	}

	// for name, coctail := range AviableCoctail {
	// 	delete(AviableCoctail, name)
	// 	name = strings.Title(name)
	// 	coctail.Name = name
	// 	coctail.Price = coctail.GetPrice()
	// 	AviableCoctail[name] = coctail
	// }

	MapsiAvailableCoctail = mapsi.New(map[string]Coctail{})
	MapsiAvailableCoctail.CopyElemntsOfMap(AllCoctail, 0, CountAvailableCoctail)
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
