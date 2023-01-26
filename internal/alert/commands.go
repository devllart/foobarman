package alert

import "devllart/foobarman/internal/state"

/**
 * Don't panic
 */

func CoctailIsReady(name string) {
	state.AddInfof("У вас получился %s — очень хорошо\n", name)
}
func DontTheRecipies() {
	state.AddInfo("Чтож жаль, но такого рецепта нет\n")
}
