package funcs

import (
	"unicode"

	"golang.org/x/exp/slices"
)

func IsVolwels(char rune) bool {
	vowels := []rune("AEIOUYАОУЭЫЯЕЁЮИ")
	return slices.Contains(vowels, unicode.ToUpper(char))
}
