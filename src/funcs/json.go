package funcs

import (
	"encoding/json"
	"io/ioutil"
)

func ParseJsonToStruct[T any](pathJson string, sliceStruct *T) {
	file, err := ioutil.ReadFile(pathJson)
	if err != nil {
		panic(err.Error())
	}
	if err := json.Unmarshal([]byte(file), sliceStruct); err != nil {
		panic(err.Error())
	}
}
