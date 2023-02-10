package state

import (
	"devllart/foobarman/internal/structs"
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
)

type State struct {
	Name    string
	RawName string
	Money   float64
	Bar     []structs.Product
	Scene   string
}

var state State

func Save() {
	if stateJson, err := json.Marshal(State{
		Name:    Name,
		RawName: RawName,
		Money:   Money,
		Bar:     Bar,
		Scene:   Scene,
	}); err == nil {
		if file, err := os.Create(RawName + "_save.json"); err == nil {
			file.Write(stateJson)
		} else {
			panic(err.Error())
		}

		// ioutil.WriteFile(RawName+"_save.json", stateJson, 666)
	} else {
		panic(err.Error())
	}
}

func Load() {
	files := filesWithSaves()
	if len(files) > 0 {
		if data, err := ioutil.ReadFile(files[0]); err == nil {
			json.Unmarshal(data, &state)
		} else {
			panic(err.Error())
		}

		Money = state.Money
		RawName = state.RawName
		Name = state.Name
		Scene = state.Scene
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
