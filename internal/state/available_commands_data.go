package state

var AllCommands = map[string]CommandStruct{
	// Standart commands
	"hideall": {
		Name:        "hediall",
		Aliases:     []string{"hideall", "hide", "спрятать всё", "спрятать"},
		Description: "Спрятать все подсказки/команды и описание",
	},
	"showall": {
		Name:        "showall",
		Aliases:     []string{"showeall", "show", "показть всё", "показать"},
		Description: "Показать все подсказки/команды и описание",
	},
	"description": {
		Name:        "description",
		Aliases:     []string{"description", "desc", "описание"},
		Description: "Показать/спрятать описание",
	},
	"cmds": {
		Name:        "Спрятать/показать команды",
		Aliases:     []string{"commands", "cmds", "команды"},
		Description: "Показать/спрятать команды",
	},
	"restart": {
		Name:        "restart",
		Aliases:     []string{"restart", "рестарт"},
		Description: "Начать игру заново не меняя имя",
	},
	"restartrand": {
		Name:        "restartrand",
		Aliases:     []string{"restartrand"},
		Description: "Начать игру заново с тем же иминем и закупится рандомными продуктами",
	},
	"exit": {
		Name:        "exit",
		Aliases:     []string{"exit", "quit", "выйти", "пока"},
		Description: "Выйти из игры",
	},
	// Main commands
	"ok": {
		Name:        "ok",
		Aliases:     []string{"ok", "ок", "всё", "закончить"},
		Description: "Закончить покупку ингридиентов",
	},
	"rand": {
		Name:        "rand",
		Aliases:     []string{"rand", "random", "рандомно"},
		Description: "Закупиться рандомно и перейти в бар",
	},
	"mix": {
		Name:        "mix",
		Aliases:     []string{"mix", "смешать"},
		Description: "Смешать напитки",
	},

	// Turn show
	"instruction": {
		Name:        "instruction",
		Aliases:     []string{"instruction", "inst", "инструкция"},
		Description: "Показать, скрыть инструкцию по приготовлению напитка",
	},

	// Change context scene
	"bar": {
		Name:        "bar",
		Aliases:     []string{"bar", "бар"},
		Description: "Пойти в бар",
	},
	"store": {
		Name:        "store",
		Aliases:     []string{"store", "магазин"},
		Description: "Пойти в магазин",
	},
	"recipes": {
		Name:        "recipes",
		Aliases:     []string{"recipes", "res", "рецепты"},
		Description: "Посмотреть рецепты приготовления коктелей",
	},
}

var StandartCommands = map[string]CommandStruct{
	"hideall":     AllCommands["hideall"],
	"showall":     AllCommands["showall"],
	"description": AllCommands["description"],
	"cmds":        AllCommands["cmds"],
	"restart":     AllCommands["restart"],
	"restartrand": AllCommands["restartrand"],
	"exit":        AllCommands["exit"],
}

var ShopCommands = map[string]CommandStruct{

	"rand":    AllCommands["rand"],
	"ok":      AllCommands["ok"],
	"bar":     AllCommands["bar"],
	"recipes": AllCommands["recipes"],
}

var BarCommands = map[string]CommandStruct{
	"mix":     AllCommands["mix"],
	"store":   AllCommands["store"],
	"recipes": AllCommands["recipes"],
}

var RecipesCommands = map[string]CommandStruct{
	"instruction": AllCommands["instruction"],
	"bar":         AllCommands["bar"],
	"store":       AllCommands["store"],
}
