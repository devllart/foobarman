package alert

import (
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
)

func ClueStore() {
	state.AddInfof(texts.ClueStore)
}
