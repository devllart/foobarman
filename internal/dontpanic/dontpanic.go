package dontpanic

import (
	"devllart/foobarman/internal/state"
	"devllart/foobarman/internal/texts"
	"log"
)

func RecoverAll() {
	state.ClearTemp()

	if err := recover(); err != nil {
		log.Println(err)
		if err != texts.IncorrectInput {
			panic(err)
		}
	}
}
