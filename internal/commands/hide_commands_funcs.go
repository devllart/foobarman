package commands

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/products"
	"devllart/foobarman/internal/state"
)

func SexInBigCity() {
	cosmopolitan := products.MapsiAvailableCoctail.GetValue("Космополитен")
	if cosmopolitan == nil {
		alert.DontTheRecipies()
		return
	}

	state.InfiniteMoney = true
	for _, ingredient := range cosmopolitan.Ingredients {
		for _, product := range products.MapsiAvailableProducts.Values {
			if product.Type == ingredient {
				buyProduct(product.Name, 0, 1)
			}
		}

	}
	state.InfiniteMoney = false
	state.Bar = append(state.Bar)
}

func ManyMoney() {
	state.Money = 999_999_999.999
}

func GetMoney() {
	state.Money += 0.01
}

func GetAllAvailabelRecipes() {
	products.AddAvailablesCoctail(-1)
}

func TurnInfiniteMoney() {
	state.InfiniteMoney = !state.InfiniteMoney
}
