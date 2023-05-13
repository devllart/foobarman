package state

import (
	"devllart/foobarman/internal/structs"
)

var (
	// ** Flags
	Run = true        // State game loop
	Mix = false       // Is command mix?
	RandomBuy = false // For seperate logics rand buy and buy of user
   
   // ** For manipulation with views
	ProductsIds = []string{} // Save output printers products
   
   // ** State of player
	Scene string = "Hello"    // Context of scenes
	LastScene string          // Last scene
	Name = ""                 // Barman name
	RawName = ""              // Barman name in raw format
	Bar = []structs.Product{} // State of bar
	CountVisitorsServed = 0   // Number of visitors served
	Stage = 0                 // STAGE
   
   // ** State of bar
   
	BarOpen = false
	NotSaler = true
	CoctailReady = false
	Order structs.Coctail
	YourCoctail structs.Coctail
	CurrentHistory = []string{}
)

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
