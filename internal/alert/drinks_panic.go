package alert

import (
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
)

/**
 * Aler for drinks with panic!
 */

// Aaaa this drink is not available! What is do it?
func PanicNotAvailableProduct(drinkName string) {
	state.AddInfof(texts.NotAvailableProduct, drinkName)
	panic(texts.IncorrectInput)
}

// Aaaa no such index! What is do it?
func PanicNotAvailableIndexProduct(drinkIndex int) {
	state.AddInfof(texts.NotAvailableIndexProduct, drinkIndex)
	panic(texts.IncorrectInput)
}

// Aaaa not enoght money! What is do it?
func NotEnoughtFundsToBuy(drinkName string, sumPrice float64) {
	// If user select buy random, then no alert about not enought money
	state.AddInfof(texts.NotEnoughFundsToBuy, drinkName, sumPrice)
	panic(texts.IncorrectInput)
}

// Aaaa such volume of drink not exitst! What is do it?
func PanicNotVolumeOfProduct(drinkName string, volume float64) {
	state.AddInfof(texts.NotVolumeOfProduct, drinkName, volume)
	panic(texts.IncorrectInput)
}

/**
 * So, nothing needs to do — all these panics are intercepted package dontpanic.
 * Just alerting user.
 */
