package products

import (
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/structs"
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
			count := 0
			for _, availableProduct := range MapsiAvailableProducts.Values {
				if product.Type == availableProduct.Type {
					count += 1
				}
			}
			if count < state.Stage {
				MapsiAvailableProducts.SetValue(name, product)
			}
		}
	}
}

var NotExitst = map[string]int{}

func AddAvailableCoctail(coctail structs.Coctail) bool {
	availableTypes := []string{}
	if coctail.GetPrice() > 3.0+float64(state.Stage)/2 || len(coctail.Ingredients) > 4+state.Stage {
		return false
	}

	for _, ingredient := range coctail.Ingredients {
		if !funcs.Contains(AvailableIngredients, ingredient) {
			NotExitst[ingredient] += 1
			return false
		} else {
			availableTypes = append(availableTypes, ingredient)
		}
	}

	AvailableTypes = append(AvailableTypes, availableTypes...)
	MapsiAvailableCoctail.SetValue(coctail.Name, coctail)
	return true
}
