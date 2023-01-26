package commands

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/state"
	"strconv"
	"strings"
)

/**
 * Get Name or Index drink's and return correct its name.
 * + alerting user if his index out of range drink's menu.
 */
func correctDrinkName(drinkNameOrIndex string) string {
	drinkIndex, err := strconv.Atoi(drinkNameOrIndex)
	if err == nil {
		if drinkIndex > len(state.DrinksIds) {
			alert.NotAvailableIndexDrink(drinkIndex)
		}
		return state.DrinksIds[drinkIndex-1]
	}

	return strings.Title(drinkNameOrIndex)
}

func correctCount(arg string) int {
	var count int = 1
	var err error
	if arg != "" {
		count, err = strconv.Atoi(arg)
	}
	if err != nil || count <= 0 {
		alert.IncorrectAmountOfDrink()
		return 0
	}

	return count
}

func correctVolume(arg string) float64 {
	var volume float64 = 0
	var err error
	if arg != "" {
		volume, err = strconv.ParseFloat(arg, 64)
	}
	if err != nil || volume <= 0 {
		alert.IncorrectVolumeOfDrink()
		return 0
	}

	return volume
}
