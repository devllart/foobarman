package dontpanic

import (
	"devllart/foobarman/internal/state"
	"log"
)

func RecoverAll() {
	state.ClearTemp()

	if err := recover(); err != nil {
		log.Println(err)
	}
}
