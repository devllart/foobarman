package data

import (
	"devllart/foobarman/internal/structs"

	"golang.org/x/exp/maps"
)

var AvailableCommands = map[string]structs.Command{}

var StandartCommands = map[string]structs.Command{
	"exit": {
		Name:        "exit",
		Aliases:     []string{"exit", "quit", "выйти", "пока"},
		Description: "Выйти из игры",
	},
	"description": {
		Name:        "description",
		Aliases:     []string{"description", "desc", "описание"},
		Description: "Показать/спрятать описание",
	},
}

var ShopCommands = map[string]structs.Command{
	"ok": {
		Name:        "ok",
		Aliases:     []string{"ok", "ок", "всё", "закончить"},
		Description: "Закончить покупку ингридиентов",
	},
	"rand": {
		Name:        "rand",
		Aliases:     []string{"rand", "random", "рандомно"},
		Description: "Закупиться рандомно",
	},
}

var BarCommands = map[string]structs.Command{
	"store": {
		Name:        "store",
		Aliases:     []string{"store", "магазин"},
		Description: "Пойти в магазин",
	},
	"mix": {
		Name:        "mix",
		Aliases:     []string{"mix", "смешать"},
		Description: "Смешать напитки",
	},
}

func init() {
	maps.Copy(AvailableCommands, StandartCommands)
	maps.Copy(AvailableCommands, ShopCommands)
	maps.Copy(AvailableCommands, BarCommands)
}
