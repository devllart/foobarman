package alert

import (
	"devllart/foobarman/src/fmtc"
	"fmt"
)

var Info = ""

func PrintInfo() {
	if Info != "" {
		fmtc.Printf("\n%s", Info)
	}
}

func AddInfo(args ...any) {
	Info += fmt.Sprint(args...)
}

func AddInfof(textf string, args ...any) {
	Info += fmtc.Sprintf(textf, args...)
}

func ClearInfo() {
	Info = ""
}
