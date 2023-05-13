package state

import "devllart/foobarman/internal/structs"

var State = structs.State{}

func SaveState() {
	State = structs.State{
		Name:                Name,
		RawName:             RawName,
		Scene:               Scene,
		LastScene:           LastScene,
		Money:               Money,
		Stage:               Stage,
		CountVisitorsServed: CountVisitorsServed,
		// Bar
		Bar:            Bar,
		ProductsIds:    ProductsIds,
		CurrentHistory: CurrentHistory,
		Status:         Status,
		// Coctail
		Order:       Order,
		YourCoctail: YourCoctail,

		// Command
		Command: Command,
		Args:    Args,
		// ** Flags
		// Game
		Mix:          Mix,
		RandomBuy:    RandomBuy,
		BarOpen:      BarOpen,
		CoctailReady: CoctailReady,
		NotSaler:     NotSaler,
		// Cheatcodes
		GodMod:        GodMod,
		InfiniteMoney: InfiniteMoney,

		// Temps
		TempBool: TempBool,
		TempStr:  TempStr,
	}
}

func UpdateStateFromSave() {
	Name = State.Name
	RawName = State.RawName
	Scene = State.Scene
	LastScene = State.LastScene
	Money = State.Money
	Stage = State.Stage
	CountVisitorsServed = State.CountVisitorsServed
	// Bar
	Bar = State.Bar
	ProductsIds = State.ProductsIds
	CurrentHistory = State.CurrentHistory
	Status = State.Status
	// Coctail
	Order = State.Order
	YourCoctail = State.YourCoctail

	// Command
	Command = State.Command
	Args = State.Args
	// ** Flags
	// Game
	Mix = State.Mix
	RandomBuy = State.RandomBuy
	BarOpen = State.BarOpen
	CoctailReady = State.CoctailReady
	NotSaler = State.NotSaler
	// Cheatcodes
	GodMod = State.GodMod
	InfiniteMoney = State.InfiniteMoney

	// Temps
	TempBool = State.TempBool
	TempStr = State.TempStr
}
