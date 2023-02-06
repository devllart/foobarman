package commands

import (
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/state"
)

func ExecHandler() (func(), bool) {
	for command, commandFunc := range CommandsForHandler {
		if CommandIs(command) {
			return commandFunc, true

		}
	}
	return nil, false
}

var CommandsForHandler = map[string]func(){
	"skip":        SkipClient,
	"restart":     Restart,
	"restartrand": RestartAndRandBuy,
	"hideall":     config.HideAll,
	"showall":     config.ShowAll,
	"cmds":        config.TurnShowCommands,
	"description": config.TurnShowDescription,
	"instruction": config.TurnShowInstruction,
	"bar":         SetScenesBar,
	"recipes":     SetScenesRecipes,
	"store":       SetScenesStore,
	"exit":        state.ExitGame,
	"start":       state.ToOpenBar,
}
