package sale

import (
	"devllart/foobarman/internal/state"
	"devllart/foobarman/src/fmtc"
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
		fmtc.Printf("%YВаш покупатель до сих пор ждёт %B%s%C\n", state.Order.Name)
	}
}
