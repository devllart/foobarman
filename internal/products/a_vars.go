package products

import "devllart/foobarman/src/mapsi"

/**
 * Global varialbles for package products
 */

// Storage (Mapsi) for available products and coctails
var MapsiAvailableCoctail mapsi.Mapsi[Coctail]
var MapsiAvailableProducts mapsi.Mapsi[ProductInfo]

// Pre-settings
var ProductsStandartTypesPrice = map[string]float64{}                  // Append standart price for used products in here
var AvailableTypes = []string{"Пиво", "Кола", "Содовая", "Обуховская"} // Products types available default || append another in here

// For selected available coctails and products in store in start game
var AvailableIngredients = []string{}
var notExistEngredient = map[string]int{}

// For added recipes
var AllCoctailEndIndex = 0
