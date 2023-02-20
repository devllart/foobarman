package sale

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/state"
	"fmt"
)

func Sale() {
	if !state.BarOpen {
		return
	}

	fmt.Println()
	if state.NotSaler {
		NewClient()
	} else if state.CoctailReady {
		CoctailReady()
	} else {
		alert.Clue("%YВаш покупатель до сих пор ждёт %B%s%C\n", state.Order.Name)
	}
}
