package scenes

import (
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/ptf"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/src/funcs"
	"fmt"
)

func Show(scene func()) {
	funcs.CliClear()
	scene()

	if config.ShowBarman {
		ptf.ShowBarmanStatus(state.Name, state.Money)
	}

	if config.ShowCommands {
		ptf.StandartCommands()
		if CurrentIs(Store) {
			ptf.FinishShoopingCommand()
		} else {
			ptf.StartShoopingCommand()
		}
	}

	if config.ShowHits {
		if state.Info != "" {
			fmt.Printf("\n%s", state.Info)
		}
	}
}

func CurrentIs(scene func()) bool {
	return fmt.Sprintf("%v", state.Scene) == fmt.Sprintf("%v", scene) // Just a little bit of codesshit
}
