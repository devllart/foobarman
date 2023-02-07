package main

import (
	"io/ioutil"
	"strconv"
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

type TextsFile struct {
	Index  int
	GoCode string
}

var textsFiles = map[string]TextsFile{}

func saveText(fileName, text string) string {
	var textsFile TextsFile
	var exist bool

	if textsFile, exist = textsFiles[fileName]; exist {
		textsFile.Index += 1
	} else {
		textsFile.Index = 1
		textsFile.GoCode = "package texts\n\n"
	}

	nameVar := strings.Title(fileName) + strconv.Itoa(textsFile.Index)
	textsFile.GoCode += "var " + nameVar + " = \"" + text + "\""

	ioutil.WriteFile("internal/texts/"+fileName+".go", []byte(textsFile.GoCode), 0644)
	textsFiles[fileName] = textsFile

	return "texts." + nameVar
}
