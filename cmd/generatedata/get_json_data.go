package main

import (
	anames "devllart/foobarman/internal/a_names"
	"encoding/json"
	"io/ioutil"
)

func getJsonData(name string, storageData any) {
	var jsonFileWithData = anames.JsonDataDir + "/" + name + ".json"
	file, err := ioutil.ReadFile(jsonFileWithData)
	if err != nil {
		panic(err.Error())
	}

	if err := json.Unmarshal([]byte(file), storageData); err != nil {
		panic(err.Error())
	}
}
