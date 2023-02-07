package commands

import (
	"devllart/foobarman/internal/products"
	"devllart/foobarman/internal/state"
)

func SexInBigCity() {
	cosmopolitan := products.MapsiAvailableCoctail.GetValue("Космополитен")
	if cosmopolitan == nil {
		state.AddInfof("%YУ вас нет такого рецепта%C")
		return
	}
	for _, ingredient := range cosmopolitan.Ingredients {
		for _, product := range products.MapsiAvailableProducts.Values {
			if product.Type == ingredient {
				money := state.Money
				state.Money = 999999
				buyProduct(product.Name, 0, 1)
				state.Money = money
			}
		}

	}
	state.Bar = append(state.Bar)
	state.AddInfof("%YЧит-код активирован%C\n")
}

func ManyMoney() {
	state.Money = 999_999_999.999
	state.AddInfof("%YТеперь ты крутой как devllart (можешь зайти не его в страницу в инсте и поставить лайк)%C\n")
}

func GetMoney() {
	state.Money += 0.01
	state.AddInfof("%YНа%C\n")
}

func GetAllAvailabelRecipes() {
	products.AddAvailablesCoctail(-1)
	state.AddInfof("%YЧит-код активирован | теперь у вас есть все доступные рецепты%C\n")
}
