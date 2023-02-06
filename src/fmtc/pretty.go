package fmtc

import "strings"

func PrettyPercent(text string) string {
	text = strings.ReplaceAll(text, ".000", "")
	text = strings.ReplaceAll(text, ".00 %", "%")
	text = strings.ReplaceAll(text, " %", "%")
	text = strings.ReplaceAll(text, " "+Red+"0%"+Reset, " ")
	text = strings.ReplaceAll(text, "  ", " ")
	return text
}
