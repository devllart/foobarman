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

var GetCommandsForScenes = getCommandsForScenes()
