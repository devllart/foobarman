package alert

import (
	"devllart/foobarman/internal/texts"
)

/**
 * Alert for products with panic!
 */

// Aaaa this product is not available! What is do it?
func PanicNotAvailableProduct(productName string) {
	Text(" %R!%C Напитка %B%s%C нет в продаже", productName)
	Line()
	panic(texts.IncorrectInput)
}

// Aaaa no such index! What is do it?
func PanicNotAvailableIndexProduct(productIndex int) {
	Text("%R!%C Напитка под номером %B%d%C нет в продаже", productIndex)
	Line()
	panic(texts.IncorrectInput)
}

// Aaaa not enoght money! What is do it?
func PanicNotEnoughtFundsToBuy(productName string, sumPrice float64) {
	Text(" %R!%C Недостаточно средств для покупки %B%s%C (общая сумма составила %Y%.2f $%C)", productName, sumPrice)
	Line()
	panic(texts.IncorrectInput)
}

// Aaaa such volume of product not exitst! What is do it?
func PanicNotVolumeOfProduct(productName string, volume float64) {
	Text("%R!%C %B%s%C с объёмом %G%.3f .л%C нет в продаже, возьмите другой объём", productName, volume)
	Line()
	panic(texts.IncorrectInput)
}

/**
 * So, nothing needs to do — all these panics are intercepted package dontpanic.
 * Just alerting user.
 */
