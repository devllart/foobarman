package commands

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/dontpanic"
	"devllart/foobarman/internal/products"
	"devllart/foobarman/internal/scenes"
	"devllart/foobarman/internal/state"
	"errors"
	"strconv"
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
		alert.Clue(answer)
		alert.Line()
	} else if commandFunc, exist := hideCommands[state.Command]; exist { // Cheat codes
		commandFunc()
		alert.CheatCodeActivate()
	} else if commandFunc, exist := ExecHandler(); exist { // Exec command
		commandFunc()
	} else if scenes.CurrentIs(scenes.Store) { // Context dependent commands start here
		buy()
	} else if scenes.CurrentIs(scenes.Bar) && CommandIs("mix") { // mix
		mix()
	} else if state.Command == strings.ToLower(state.RawName) { // If user input name barman
		alert.Text(" %G:)%C Что ?\n")
	} else { // otherwise say user "not command"
		alert.UnknownCommand(state.Command)
	}

	state.ClearTemp() // Cleaning
}

func FakeExec() error {
	command := state.Command

	if command == "" { // Empty command -> None action
		return nil
	} else if _, exist := answers[command]; exist { // Answers
		return nil
	} else if _, exist := hideCommands[command]; exist { // Cheat codes
		return nil
	} else if _, exist := ExecHandler(); exist { // Exec command
		return nil
	} else if command == strings.ToLower(state.RawName) { // If user input name barman
		return nil
	}

	index, err := strconv.Atoi(command)
	if err == nil {
		if index <= len(state.ProductsIds) {
			return nil
		} else {
			return errors.New("Нет продукта под индексом " + strconv.Itoa(index))
		}
	}

	for _, commands := range state.AvailableCommands() {
		for _, alias := range commands.Aliases {
			if command == alias {
				return nil
			}
		}
	}

	if _, exist := products.AvailableProducts[strings.Title(command)]; exist {
		return nil
	}

	return errors.New("Нет такой команды")
}
