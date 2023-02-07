package main

import (
	"devllart/foobarman/src/funcs"
	"io/ioutil"
	"strconv"
)

type TextsFile struct {
	Index  int
	GoCode string
}

var textsFiles = map[string]TextsFile{}

func saveText(context, text string) string {
	var textsFile TextsFile
	var exist bool

	if textsFile, exist = textsFiles[context]; exist {
		textsFile.Index += 1
	} else {
		textsFile.Index = 1
		textsFile.GoCode = "package texts\n\n"
	}

	nameVar := context + strconv.Itoa(textsFile.Index)
	textsFile.GoCode += "var " + nameVar + " = \"" + text + "\"\n\n"

	ioutil.WriteFile("internal/texts/"+funcs.ToSnakeCase(context)+".go", []byte(textsFile.GoCode), 0644)
	textsFiles[context] = textsFile

	return "texts." + nameVar
}
