package alert

import (
	"devllart/foobarman/internal/state"
)

// Add line
func Line() {
	state.AddInfo("\n")
}

// Just Text
func Text(text string, args ...any) {
	state.AddInfof(text, args...)
}

// Colorized alert
func TextYellow(text string, args ...any) {
	state.AddInfof("%Y"+text+"%C", args...)
}

func TextRed(text string, args ...any) {
	state.AddInfof("%R"+text+"%C", args...)
}

func TextBlue(text string, args ...any) {
	state.AddInfof("%B"+text+"%C", args...)
}

func TextGreen(text string, args ...any) {
	state.AddInfof("%G"+text+"%C", args...)
}

// Aliases
var Clue = TextYellow
var Error = TextRed
