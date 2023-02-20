package state

import (
	structsdata "devllart/foobarman/internal/data/structs_data"
	"devllart/foobarman/internal/structs"
	"devllart/foobarman/src/funcs"
)

// Local
var cmds = map[string]structs.Command{}

// Context commands

var StandartCommands = getCommands("hideall", "showall", "description", "cmds", "restart", "restartrand", "exit")
var ShopCommands = getCommands("rand", "ok", "bar", "recipes")
var BarCommands = getCommands("start", "mix", "store", "recipes")
var BarWithClientCommands = getCommands("skip")
var RecipesCommands = getCommands("instruction", "bar", "store")

func getCommands(keys ...string) map[string]structs.Command {
	cmds := map[string]structs.Command{}
	for _, key := range keys {
		cmds[key] = structsdata.AllCommands[key]
	}

	return cmds
}

func AvailableCommands() map[string]structs.Command {
	if State.LastScene == State.Scene {
		return cmds
	}
	State.LastScene = State.Scene
	cmds = map[string]structs.Command{}

	funcs.MapsCopy(cmds, *GetCommandsForScenes(State.Scene))
	funcs.MapsCopy(cmds, StandartCommands)

	if !NotSaler {
		funcs.MapsCopy(cmds, BarWithClientCommands)
	}

	return cmds
}

func GetCommandsForScenes(sceneName string) *map[string]structs.Command {
	if sceneName == "Bar" {
		return &BarCommands
	} else if sceneName == "Store" {
		return &ShopCommands
	} else if sceneName == "Recipes" {
		return &RecipesCommands
	} else {
		return &StandartCommands
	}
}
