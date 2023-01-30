package state

import (
	"devllart/foobarman/src/funcs"

	"golang.org/x/exp/maps"
)

func AvailableCommands() map[string]CommandStruct {
	sceneName := funcs.FuncName(Scene)
	if LastScene == sceneName {
		return cmds
	}

	LastScene = sceneName
	cmds = map[string]CommandStruct{}

	maps.Copy(cmds, GetCommandsForScenes(sceneName))
	maps.Copy(cmds, StandartCommands)
	return cmds
}
