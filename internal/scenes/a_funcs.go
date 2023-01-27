package scenes

import (
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/ptf"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"
	"fmt"
)

/**
 * Func print scene, barman status and hints
 * ( Depending on config flags )
 */
func Show(scene func()) {
	funcs.CliClear() // Clear Console
	scene()          // Print scene to console

	// If show barman is on, then show barman status (name and money)
	if config.ShowBarman {
		fmtc.Printf(texts.BarmanStatus, state.Name, state.Money)
	}

	// If show commands is on, then show commands
	if config.ShowCommands {
		ptf.StandartCommands() // Show standart commands

		if CurrentIs(Store) {
			// If context scene "Store" then show command finis shooping.
			ptf.FinishShoopingCommand()
		} else {
			// Otherwise show command go to store
			ptf.StartShoopingCommand()
		}
	}

	// If show hints is on, then show hints
	if config.ShowHits {
		if state.Info != "" {
			fmt.Printf("\n%s", state.Info)
		}
	}
}

/**
 * Check current context scenes and compare with gets scene
 * PS: All scenes its functions
 */
func CurrentIs(scene func()) bool {
	return fmt.Sprintf("%v", state.Scene) == fmt.Sprintf("%v", scene) // Just a little bit of codesshit
}
