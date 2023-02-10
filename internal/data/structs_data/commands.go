package structsdata

import "devllart/foobarman/internal/structs"

var AllCommands = map[string]structs.Command{
	// Standart commands
	"hideall": {
		Aliases:     []string{"hideall", "hide", "спрятать всё", "спрятать"},
		Description: "Спрятать все подсказки/команды и описание",
	},
	"showall": {
		Aliases:     []string{"showall", "show", "показть всё", "показать"},
		Description: "Показать все подсказки/команды и описание",
	},
	"description": {
		Aliases:     []string{"description", "desc", "описание"},
		Description: "Показать/спрятать описание",
	},
	"cmds": {
		Aliases:     []string{"commands", "cmds", "команды"},
		Description: "Показать/спрятать команды",
	},
	"restart": {
		Aliases:     []string{"restart", "рестарт"},
		Description: "Начать игру заново не меняя имя",
	},
	"restartrand": {
		Aliases:     []string{"restartrand"},
		Description: "Начать игру заново с тем же иминем и закупится рандомными продуктами",
	},
	"exit": {
		Aliases:     []string{"exit", "quit", "выйти", "пока"},
		Description: "Выйти из игры",
	},
	// Main commands
	"start": {
		Aliases:     []string{"go", "start", "open", "open bar", "старт", "поехали", "открыться", "открыть бар"},
		Description: "Начать игру (открыть бар и начать принимать поситителей)",
	},
	"skip": {
		Aliases:     []string{"skip", "пропустить"},
		Description: "Пропустить клиента",
	},
	"ok": {
		Aliases:     []string{"ok", "ок", "всё", "закончить"},
		Description: "Закончить покупку продуктов",
	},
	"rand": {
		Aliases:     []string{"rand", "random", "рандомно"},
		Description: "Закупиться рандомно и перейти в бар",
	},
	"mix": {
		Aliases:     []string{"mix", "смешать"},
		Description: "Смешать напитки",
	},

	// Turn show
	"instruction": {
		Aliases:     []string{"instruction", "inst", "инструкция"},
		Description: "Показать, скрыть инструкцию по приготовлению напитка",
	},

	// Change context scene
	"bar": {
		Aliases:     []string{"bar", "бар"},
		Description: "Пойти в бар",
	},
	"store": {
		Aliases:     []string{"store", "магазин"},
		Description: "Пойти в магазин",
	},
	"recipes": {
		Aliases:     []string{"recipes", "res", "рецепты"},
		Description: "Посмотреть рецепты приготовления коктелей",
	},
}
