package ptf

import (
	"devllart/foobarman/internal/data"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
)

func StandartCommands() {
	fmtc.Printf(texts.AllowCommands)
	for i := range data.StandartCommands {
		data.StandartCommands[i].ShowClue()
	}
}

func FinishShoopingCommand() {
	for i := range data.ShopCommands {
		data.ShopCommands[i].ShowClue()
	}
}

func StartShoopingCommand() {
	for i := range data.BarCommands {
		data.BarCommands[i].ShowClue()
	}
}
