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
		textsFile.Index = 0
		textsFile.GoCode = "package texts\n\nvar " + context + " = []string{\n"
	}

	// nameVar := context + strconv.Itoa(textsFile.Index)
	textsFile.GoCode += "\"" + text + "\",\n"

	ioutil.WriteFile("internal/texts/a_generate_"+funcs.ToSnakeCase(context)+".go", []byte(textsFile.GoCode+"}"), 0644)
	textsFiles[context] = textsFile

	return "texts." + context + "[" + strconv.Itoa(textsFile.Index) + "]"
}
