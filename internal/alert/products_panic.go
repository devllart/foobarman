package alert

import (
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
)

/**
 * Aler for products with panic!
 */

// Aaaa this product is not available! What is do it?
func PanicNotAvailableProduct(productName string) {
	state.AddInfof(texts.NotAvailableProduct, productName)
	panic(texts.IncorrectInput)
}

// Aaaa no such index! What is do it?
func PanicNotAvailableIndexProduct(productIndex int) {
	state.AddInfof(texts.NotAvailableIndexProduct, productIndex)
	panic(texts.IncorrectInput)
}

// Aaaa not enoght money! What is do it?
func NotEnoughtFundsToBuy(productName string, sumPrice float64) {
	// If user select buy random, then no alert about not enought money
	state.AddInfof(texts.NotEnoughFundsToBuy, productName, sumPrice)
	panic(texts.IncorrectInput)
}

// Aaaa such volume of product not exitst! What is do it?
func PanicNotVolumeOfProduct(productName string, volume float64) {
	state.AddInfof(texts.NotVolumeOfProduct, productName, volume)
	panic(texts.IncorrectInput)
}

/**
 * So, nothing needs to do — all these panics are intercepted package dontpanic.
 * Just alerting user.
 */
