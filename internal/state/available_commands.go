package state

import (
	"devllart/foobarman/src/funcs"

	"golang.org/x/exp/maps"
)

func AvailableCommands() map[string]CommandStruct {
	cmds := map[string]CommandStruct{}
	maps.Copy(cmds, StandartCommands)

	if funcs.IsFunc(Scene, "Store") {
		maps.Copy(cmds, ShopCommands)
	} else if funcs.IsFunc(Scene, "Bar") {
		maps.Copy(cmds, BarCommands)
	}

	return cmds
}
