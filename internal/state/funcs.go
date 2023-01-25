package state

import "fmt"

func AddInfo(args ...any) {
	Info += fmt.Sprint(args...)
}

func AddInfof(textf string, args ...any) {
	Info += fmt.Sprintf(textf, args...)
}

func ClearTemp() {
	Command = ""
	Arg1 = ""
	Arg2 = ""

	DrinksIds = []string{}
}
