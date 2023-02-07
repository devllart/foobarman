package main

import (
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
