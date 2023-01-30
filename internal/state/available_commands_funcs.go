package state

func getCommandsForScenes() func(sceneName string) map[string]CommandStruct {
	barCommands := BarCommands
	shopCommands := ShopCommands
	recipesCommands := RecipesCommands
	standartCommands := StandartCommands

	return func(sceneName string) map[string]CommandStruct {
		if sceneName == "Bar" {
			return barCommands
		} else if sceneName == "Store" {
			return shopCommands
		} else if sceneName == "Recipes" {
			return recipesCommands
		} else {
			return standartCommands
		}
	}
}

func getCommands(keys ...string) map[string]CommandStruct {
	cmds := map[string]CommandStruct{}
	for _, key := range keys {
		cmds[key] = AllCommands[key]
	}

	return cmds
}

var GetCommandsForScenes = getCommandsForScenes()
