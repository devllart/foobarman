package alert

import "devllart/foobarman/internal/state"

func Show() {
	for _, text := range state.Alerts {
		state.Info += text
	}
}
