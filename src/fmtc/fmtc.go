package fmtc

import (
	"fmt"
	"strings"

	"github.com/TwiN/go-color"
)

var Replacer = strings.NewReplacer("%R", color.Red, "%G", color.Green, "%B", color.Blue, "%Y", color.Yellow, "%C", color.Reset)

func Sprintf(text string, args ...any) string {
	return fmt.Sprintf(ReplaceText(text), args...)
}

func Printf(text string, args ...any) {
	fmt.Print(Sprintf(text, args...))
}

func ReplaceText(text string) string {
	return Replacer.Replace(text)
}
