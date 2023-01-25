package scenes

import (
	"devllart/foobarman/internal/funcs"
	"devllart/foobarman/internal/state"
	"fmt"
)

func Show(scene func()) {
	funcs.CliClear()
	scene()
	if state.Info != "" {
		fmt.Println()
		fmt.Println(state.Info)
	}
	fmt.Println()
}
