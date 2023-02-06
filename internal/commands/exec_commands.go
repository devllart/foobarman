package commands

import (
	"devllart/foobarman/internal/dontpanic"
	"devllart/foobarman/internal/scenes"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
	"strings"
)

/**
 * Execute command
 */

func Exec() {
	defer dontpanic.RecoverAll()
	state.Command = strings.ToLower(state.Command)

	if state.Command == "" { // Empty command -> None action
		return
	} else if answer, exist := answers[state.Command]; exist { // Answers
		state.AddInfof("%Y" + answer + "%C\n")
	} else if commandFunc, exist := hideCommands[state.Command]; exist { // Cheat codes
		commandFunc()
	} else if commandFunc, exist := ExecHandler(); exist { // Exec command
		commandFunc()
	} else if scenes.CurrentIs(scenes.Store) { // Context dependent commands start here
		buy()
	} else if scenes.CurrentIs(scenes.Bar) && CommandIs("mix") { // mix
		mix()
	} else if state.Command == strings.ToLower(state.RawName) { // If user input name barman
		state.AddInfof(" %G:)%C Что ?\n")
	} else { // otherwise say user "not command"
		state.AddInfof(texts.UnknownCommand, "%B"+state.Command+"%C")
	}

	state.ClearTemp() // Cleaning
}
