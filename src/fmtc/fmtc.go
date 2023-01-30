package fmtc

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/TwiN/go-color"
)

var Reader = bufio.NewReader(os.Stdin)
var PreReplacer = strings.NewReplacer("Оливки", "%GОливки%C", "Coca-Cola", "%RCoca-Cola%C", "Листья", "%GЛистья%C", "красный", "%Rкрасный%C", "синий", "%Bсиний%C", "зелёный", "%Gзелёный%С", "мята", "%Gмята%C", "Мята", "%GМята%C", "лаймовая цедра", "%Gламовая цедра%C", "Лаймовая цедра", "лаймовая", "%Gлаймовая%C", "Лаймовая", "%GЛаймовая%C", "%GЛаймовая цедра%C", "лаймовый сок", "%Gлаймовый сок%C", "Лаймовый сок", "%GЛаймовый сок%C", "лаймовый", "%G%лаймовыйC", "Лаймовый", "%GЛаймовый%C", "лайма", "%Gлайма%C", "лайм", "%Gлайм%C", "Лайма", "%GЛайма%C", "Лайм", "%GЛайм%C", "лимон", "%Yлимон%C", "Лимон", "%YЛимон%C")
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
	text = fmt.Sprintf(ReplaceText(text), args...)
	regex := regexp.MustCompile("\\.0+\\.")
	text = regex.ReplaceAllString(text, ".")
	text = strings.ReplaceAll(text, " "+color.Red+"0.00 %", "")
	text = strings.ReplaceAll(text, ".00 %", "%")
	text = strings.ReplaceAll(text, " %", "%")
	return text
}

func Printf(text string, args ...any) {
	fmt.Print(Sprintf(text, args...))
}

func ReplaceText(text string) string {
	text = PreReplacer.Replace(text) // Too long
	return Replacer.Replace(text)
}

func Errorf(text string, args ...any) error {
	return fmt.Errorf(Sprintf(text, args...))
}

func Perrorf(text string, args ...any) {
	err := Errorf(text, args...)
	panic(err)
}
