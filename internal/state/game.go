package state

import (
	"devllart/foobarman/internal/structs"
)

// ** Flags

var Run = true        // State game loop
var Mix = false       // Is command mix?
var RandomBuy = false // For seperate logics rand buy and buy of user

// ** For manipulation with views
var ProductsIds = []string{} // Save output printers products

// ** State of player
var Scene string = "Hello"    // Context of scenes
var LastScene string          // Last scene
var Name = ""                 // Barman name
var RawName = ""              // Barman name in raw format
var Bar = []structs.Product{} // State of bar
var CountVisitorsServed = 0   // Number of visitors served
var Stage = 0                 // STAGE

// ** State of bar

var BarOpen = false
var NotSaler = true
var CoctailReady = false
var Order structs.Coctail
var YourCoctail structs.Coctail
var CurrentHistory = []string{}

func ToOpenBar() {
	BarOpen = true
}

func CloseBar() {
	BarOpen = false
}

func Restart() {
	ClearTemp()
	Bar = []structs.Product{}
	Money = 300.33
	CloseBar()
	NotSaler = true
	TempBool = false
}

func ExitGame() {
	Run = false
}
