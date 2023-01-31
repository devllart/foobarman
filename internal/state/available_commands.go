package state

import (
	"devllart/foobarman/src/funcs"
)

func AvailableCommands() map[string]CommandStruct {
	sceneName := funcs.FuncName(Scene)
	if LastScene == sceneName {
		return cmds
	}

	LastScene = sceneName
	cmds = map[string]CommandStruct{}

	mapsCopy(cmds, GetCommandsForScenes(sceneName))
	mapsCopy(cmds, StandartCommands)
	return cmds
}

func mapsCopy[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2) {
	for k, v := range src {
		dst[k] = v
	}
}

