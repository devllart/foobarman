package alert

import (
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
)

/**
 * Don't panic
 */

func CoctailIsReady(name string) {
	state.AddInfof(texts.CoctailIsReady, name)
}

func DontTheRecipies() {
	state.AddInfof(texts.DontTheRecipies)
}
