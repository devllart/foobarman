package commands

import (
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/dontpanic"
	"devllart/foobarman/internal/drinks"
	"devllart/foobarman/internal/scenes"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
	"strings"
)

func Exec() {
	defer dontpanic.RecoverAll()
	state.Command = strings.ToLower(state.Command)

	if state.Command == "" { // Empty command -> None action
		return
	} else if state.Command == "sexinbigcity" { // Cheat code
		for _, ingredient := range drinks.MapsiAvailableCoctail.GetValue("Космополитен").Ingredients {
			for _, drink := range drinks.MapsiAvailableDrinks.Values {
				if drink.Type == ingredient {
					// buyTransaction(drink.Name)
					buyDrink(drink.Name, 0, 1)
				}
			}

		}
		state.Bar = append(state.Bar)
	} else if CommandIs("exit") { // command exit -> close app, exit from app, .... you undestened
		state.Run = false
	} else if CommandIs("start") {
		state.BarOpen = true
	} else if CommandIs("skip") {
		state.NotSaler = true
	} else if CommandIs("restart") {
		Restart()
	} else if CommandIs("restartrand") {
		RestartAndRandBuy()
	} else if CommandIs("hideall") { // Standart commands start here
		config.HideAll()
	} else if CommandIs("showall") {
		config.ShowAll()
	} else if CommandIs("cmds") {
		config.TurnShowCommands()
	} else if CommandIs("description") {
		config.TurnShowDescription()
	} else if CommandIs("instruction") {
		config.TurnShowInstruction()
	} else if CommandIs("bar") { // Change scenes context commands start here
		state.Scene = scenes.Bar
	} else if CommandIs("recipes") { // Change scenes context commands start here
		state.Scene = scenes.Recipes
	} else if CommandIs("store") {
		state.Scene = scenes.Store
	} else if scenes.CurrentIs(scenes.Store) { // Context dependent commands start here
		buy()
	} else if scenes.CurrentIs(scenes.Bar) && CommandIs("mix") {
		mix()
	} else { // otherwise say user "not command"
		state.AddInfof(texts.UnknownCommand, "%B"+state.Command+"%C")
	}

	state.ClearTemp() // Cleaning
}
