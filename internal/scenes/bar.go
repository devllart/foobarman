package scenes

import (
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"

	"fmt"
)

func Bar() {
	fmtc.Printf(texts.SceneBar, funcs.Indent(15))
	for i, drink := range state.Bar {
		fmt.Printf("%d. ", i+1)
		drink.Show()
	}
}
