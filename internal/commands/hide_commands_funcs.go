package commands

import (
	"devllart/foobarman/internal/products"
	"devllart/foobarman/internal/state"
)

func SexInBigCity() {
	for _, ingredient := range products.MapsiAvailableCoctail.GetValue("Космополитен").Ingredients {
		for _, drink := range products.MapsiAvailableProducts.Values {
			if drink.Type == ingredient {
				money := state.Money
				state.Money = 999999
				buyProduct(drink.Name, 0, 1)
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
