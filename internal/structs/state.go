package structs

type State struct {
	// Count
	Stage, CountVisitorsServed int
	// Money
	Money float64
	// ** Strings
	// Name
	Name, RawName string
	// Scene
	Scene, LastScene string
	// Status
	Status string
	// Command
	Command string
	Args    []string
	// Bar
	CurrentHistory, ProductsIds []string
	Bar                         []Product
	// Coctail
	Order, YourCoctail Coctail
	// ** Flags (Bool)
	// Game
	Mix, RandomBuy, BarOpen, NotSaler, CoctailReady bool
	// Cheatcodes
	GodMod, InfiniteMoney bool

	// Temp
	TempBool bool
	TempStr  string
}
