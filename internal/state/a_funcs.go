package state

import (
	"devllart/foobarman/src/fmtc"
	"fmt"
)

func AddInfo(args ...any) {
	Info += fmt.Sprint(args...)
}

func AddInfof(textf string, args ...any) {
	Info += fmtc.Sprintf(textf, args...)
}

func ClearTemp() {
	Command = ""
	Arg1 = ""
	Arg2 = ""

	DrinksIds = []string{}
}
