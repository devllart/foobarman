package alert

import (
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
)

/**
 * Alert for scenes don't panic
 */

func ClueStore() {
	state.AddInfof(texts.ClueStore)
}
