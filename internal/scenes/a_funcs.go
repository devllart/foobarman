package scenes

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"
)

/**
 * Func print scene, barman status and hints
 * ( Depending on config flags )
 */

func Show(scene string) {
	funcs.CliClear()      // Clear Console
	FuncFromName(scene)() // Print scene to console

	// If show barman is on, then show barman status (name and money)
	if config.ShowBarman {
		fmtc.Printf(texts.BarmanStatus, state.Name, state.Money, state.CountVisitorsServed)
	}

	// If show commands is on, then show commands
	if config.ShowCommands {
		for _, command := range state.AvailableCommands() {
			command.ShowClue()
		}
	}

	// If show hints is on, then show hints
	if config.ShowHits {
		alert.PrintInfo()
	}
}

/**
 * Check current context scenes and compare with gets scene
 * PS: All scenes its functions
 */

// func CurrentIs(scene interface{}) bool {
// 	return funcs.IsFunc(state.Scene, funcs.FuncName(scene))
// }

func FuncFromName(name string) (scene func()) {
	switch name {
	case "Hello":
		scene = Hello
	case "Store":
		scene = Store
	case "Bar":
		scene = Bar
	case "Recipes":
		scene = Recipes
	}
	return
}
