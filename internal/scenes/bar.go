package scenes

import (
	"devllart/foobarman/internal/bars"
	"fmt"
)

func Bar() {
	fmt.Println("           [ВАШ БАР]")
	for _, drink := range bars.Bar {
		drink.Show()
	}
}
