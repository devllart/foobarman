package products

import (
	"devllart/foobarman/src/funcs"
)

func AddAvailablesCoctail(count int) {
	for i, coctail := range AllCoctail[AllCoctailEndIndex:] {
		if AddAvailableCoctail(coctail) {
			count -= 1
		}
		if count == 0 {
			AllCoctailEndIndex = i
			break
		}
	}

	for name, product := range AvailableProducts {
		if funcs.Contains(AvailableTypes, product.Type) && MapsiAvailableProducts.GetValue(product.Name) == nil {
			MapsiAvailableProducts.SetValue(name, product)
		}
	}
}

func AddAvailableCoctail(coctail Coctail) bool {
	availableTypes := []string{}
	for _, ingredient := range coctail.Ingredients {
		if !funcs.Contains(AvailableIngredients, ingredient) {
			return false
		} else {
			availableTypes = append(availableTypes, ingredient)
		}
	}

	AvailableTypes = append(AvailableTypes, availableTypes...)
	MapsiAvailableCoctail.SetValue(coctail.Name, coctail)
	return true
}
