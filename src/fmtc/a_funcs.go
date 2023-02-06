package fmtc

import (
	"fmt"
	"strings"
)

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
	text = fmt.Sprintf(ReplaceText(text), args...)
	text = PrettyPercent(text)
	return text
}

func Printf(text string, args ...any) {
	fmt.Print(Sprintf(text, args...))
}

func ReplaceText(text string) string {
	text = PreReplacer.Replace(text) // Too long
	return ColorReplacer.Replace(text)
}

func Errorf(text string, args ...any) error {
	return fmt.Errorf(Sprintf(text, args...))
}

func Perrorf(text string, args ...any) {
	err := Errorf(text, args...)
	panic(err)
}
