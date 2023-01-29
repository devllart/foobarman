package fmtc

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/TwiN/go-color"
)

var Reader = bufio.NewReader(os.Stdin)
var Replacer = strings.NewReplacer("%R", color.Red, "%G", color.Green, "%B", color.Blue, "%Y", color.Yellow, "%C", color.Reset, "%f", "%.3f")

func Scan(sym byte) string {
	text, _ := Reader.ReadString(sym)
	return strings.Trim(text, "\n")
}

func Sprintf(text string, args ...any) string {
	for i, arg := range args {
		switch v := arg.(type) {
		case string:
			args[i] = ReplaceText(v)
		}
	}
	return fmt.Sprintf(ReplaceText(text), args...)
}

func Printf(text string, args ...any) {
	fmt.Print(Sprintf(text, args...))
}

func ReplaceText(text string) string {
	return Replacer.Replace(text)
}

func Errorf(text string, args ...any) error {
	return fmt.Errorf(Sprintf(text, args...))
}

func Perrorf(text string, args ...any) {
	err := Errorf(text, args...)
	panic(err)
}
