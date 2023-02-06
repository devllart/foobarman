package main

import (
	"devllart/foobarman/internal/products"
	"encoding/json"
	"io/ioutil"
	"strings"
)

func replaceBitch(str string) string {
	str = strings.ReplaceAll(str, "\"", "\\\"")
	str = strings.ReplaceAll(str, "\n", "\\n")
	return str
}

func joinBitch(strSlice []string) string {
	resStr := ""
	for i, str := range strSlice {
		if i > 0 {
			resStr += "\", \""
		}
		resStr += replaceBitch(str)
	}
	return resStr
}

func getCoctails() []products.Coctail {
	coctails := []products.Coctail{}
	file, err := ioutil.ReadFile(jsonFileWithData)
	if err != nil {
		panic(err.Error())
	}

	if err := json.Unmarshal([]byte(file), &coctails); err != nil {
		panic(err.Error())
	}

	return coctails
}
