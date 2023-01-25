package scenes

import (
	"devllart/foobarman/internal/state"
	"fmt"
)

func Bar() {
	fmt.Print("           [ВАШ БАР]\n\n")
	for _, drink := range state.Bar {
		drink.Show()
	}
}
