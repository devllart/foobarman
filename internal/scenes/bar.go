package scenes

import (
	"devllart/foobarman/internal/state"
	"devllart/foobarman/src/funcs"

	"fmt"

	"github.com/TwiN/go-color"
)

func Bar() {
	fmt.Printf("%s%s[%sБАР%s]%s\n\n", funcs.Indent(15), color.Red, color.Yellow, color.Red, color.Reset)
	for i, drink := range state.Bar {
		fmt.Printf("%d.", i+1)
		drink.Show()
	}
}
