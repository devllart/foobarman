package locale

import (
	"devllart/foobarman/config"

	"golang.org/x/text/message"
)

func GetText(text string, args ...any) string {
	return message.NewPrinter(config.LangTag).Sprintf(text, args...)
}


