package products

import (
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/funcs"
	"fmt"
	"sort"
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

	for name, product := range AvailableProducts {
		if funcs.Contains(AvailableTypes, product.Type) {
			MapsiAvailableProducts.SetValue(name, product)
		}
	}
}

func GetTaste(product ProductInfo) *string {
	return NewTastes.GetValue(product.Name, product.Type, texts.Unknown)
}

func AddAvailablesCoctail(count int) {
	for i, coctail := range AllCoctail[AllCoctailEndIndex:] {
		if AddAvailableCoctail(coctail) {
			count -= 1
		}
		if count == 0 {
			AllCoctailEndIndex = i
			return
		}
	}
}

func AddAvailableCoctail(coctail Coctail) bool {
	availableTypes := []string{}
	for _, ingredient := range coctail.Ingredients {
		if !funcs.Contains(AvailableIngredients, ingredient) {
			// if _, exist := notExistEngredient[ingredient]; exist {
			// 	notExistEngredient[ingredient] += 1
			// } else {
			// 	notExistEngredient[ingredient] = 1
			// }
			// panic(ingredient)
			return false
		} else {
			availableTypes = append(availableTypes, ingredient)
		}
	}

	AvailableTypes = append(AvailableTypes, availableTypes...)
	MapsiAvailableCoctail.SetValue(coctail.Name, coctail)
	return true
}

func sortIntOfMapAndPrint(m map[string]int) {
	// m := map[string]int{"hello": 10, "foo": 20, "bar": 20}
	n := map[int][]string{}
	var a []int
	for k, v := range m {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for _, k := range a {
		for _, s := range n[k] {
			fmt.Printf("%s, %d\n", s, k)
		}
	}
}
