package alert

// Add line
func Line() {
	AddInfo("\n")
}

// Just Text
func Text(text string, args ...any) {
	AddInfof(text, args...)
}

func Textf(text string, args ...any) {
	AddInfof(text, args...)
}

// Colorized alert
func TextYellow(text string, args ...any) {
	Text("%Y"+text+"%C", args...)
}

func TextRed(text string, args ...any) {
	Text("%R"+text+"%C", args...)
}

func TextBlue(text string, args ...any) {
	Text("%B"+text+"%C", args...)
}

func TextGreen(text string, args ...any) {
	Text("%G"+text+"%C", args...)
}

// Aliases
var Clue = TextYellow
var Error = TextRed
