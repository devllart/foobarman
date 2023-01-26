package alert

import (
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"

	"github.com/TwiN/go-color"
)

func ClueStore() {
	state.AddInfof(texts.ClueStore, color.Yellow, color.Reset)
}
