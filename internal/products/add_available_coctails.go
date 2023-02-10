package products

import (
	"devllart/foobarman/internal/config"
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
			if count < config.Stage {
				MapsiAvailableProducts.SetValue(name, product)
			}
		}
	}
}

func AddAvailableCoctail(coctail structs.Coctail) bool {
	availableTypes := []string{}
	if coctail.GetPrice() > 2.0+float64(config.Stage)/2 || len(coctail.Ingredients) > 4+config.Stage {
		return false
	}
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
