package commands

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/products"
	"devllart/foobarman/internal/state"
)

func buyTransaction(productName string, count int, volume, price float64, rand bool) {
	if state.TempBool {
		return
	}

	sumPrice := float64(count) * price

	if !state.InfiniteMoney && state.Money-sumPrice < 0 {
		if !rand {
			alert.PanicNotEnoughtFundsToBuy(productName, sumPrice)
		}
		state.TempBool = true
		return
	}

	state.SubMoney(sumPrice)

	newProduct := products.New(productName, volume, count)

	// Buy the not first product (exitst in the bar yet)
	if product, i := ProductExistYet(newProduct.Name, newProduct.Volume); product != nil {
		product.Count += count
		state.Bar[i] = *product
		alert.ProductBoughtYet(productName, newProduct.TypeVolume(), volume, sumPrice, count, product.Count)
		return
	}

	// Buy the first product (not in the bar yet)
	state.Bar = append(state.Bar, newProduct)
	alert.ProductBought(productName, newProduct.TypeVolume(), volume, sumPrice, count)
}
