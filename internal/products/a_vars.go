package products

import (
	"devllart/foobarman/internal/structs"
	"devllart/foobarman/src/mapsi"
)

/**
 * Global varialbles for package products
 */

// Storage (Mapsi) for available products and coctails
var MapsiAvailableCoctail mapsi.Mapsi[structs.Coctail]
var MapsiAvailableProducts mapsi.Mapsi[structs.ProductInfo]

// Pre-settings
var AvailableTypes = []string{"Пиво", "Кола", "Содовая", "Обуховская"} // Products types available default || append another in here

// For selected available coctails and products in store in start game
var AvailableIngredients = []string{}
var notExistEngredient = map[string]int{}

// For added recipes
var AllCoctailEndIndex = 0
