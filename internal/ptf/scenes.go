package ptf

import (
	"devllart/foobarman/internal/data"
	"devllart/foobarman/internal/texts"
	"fmt"

	"github.com/TwiN/go-color"
)

func StandartCommands() {
	fmt.Print(texts.AllowCommands)
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

func ShowBarmanStatus(name string, money float64) {
	fmt.Printf(texts.BarmanStatus, name, color.Yellow, money, color.Reset)
}
