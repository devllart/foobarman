package state

import (
	"devllart/foobarman/internal/drinks"
	"devllart/foobarman/src/fmtc"
	"fmt"
	"strconv"
	"strings"
)

func AddInfo(args ...any) {
	Info += fmt.Sprint(args...)
}

func AddInfof(textf string, args ...any) {
	Info += fmtc.Sprintf(textf, args...)
}

func ClearTemp() {
	Command = ""
	Args = []string{}
	Alerts = []string{}
	TempBool = false

	DrinksIds = []string{}
}

func Restart() {
	ClearTemp()
	Bar = []drinks.Drink{}
	Money = 300.33
}

func SetCommand(text string) {
	for i, word := range strings.Split(text, " ") {
		if i == 0 {
			Command += word + " "
		} else if _, err := strconv.ParseFloat(word, 64); err == nil {
			Args = append(Args, word)
		} else {
			Command += word + " "
		}
	}

	if len(Args) == 0 {
		Args = make([]string, 3)
	}
	Command = Command[:len(Command)-1]
}
