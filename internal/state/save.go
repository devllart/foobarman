package state

import (
	"encoding/json"
	"os"
	"regexp"
)

func Save() {
	SaveState()
	if stateJson, err := json.MarshalIndent(State, "", "  "); err == nil {
		if file, err := os.Create(RawName + "_save.json"); err == nil {
			file.Write(stateJson)
		} else {
			panic(err.Error())
		}
	} else {
		panic(err.Error())
	}
}

func Load() {
	files := filesWithSaves()
	if len(files) > 0 {
		if data, err := os.ReadFile(files[0]); err == nil {
			json.Unmarshal(data, &State)
		} else {
			panic(err.Error())
		}
		UpdateStateFromSave()
		// Money = state.Money
		// RawName = state.RawName
		// Name = state.Name
		// Scene = state.Scene
	}
}

func filesWithSaves() []string {
	listFiles := []string{}
	entries, err := os.ReadDir("./")
	if err != nil {
		panic(err.Error)
	}

	for _, e := range entries {
		if regexp.MustCompile(".*_save.json").Match([]byte(e.Name())) {
			listFiles = append(listFiles, e.Name())
		}
	}

	return listFiles
}
