package ptf

import (
	"devllart/foobarman/internal/texts"
	"fmt"

	"github.com/TwiN/go-color"
)

func DrinkShow(name string, count int, alc, volume, lastVolume float64) {
	fmt.Printf(texts.ShowDrinkInBar, color.Blue, name, color.Reset, color.Red, alc, color.Reset, color.Green, volume, color.Reset, color.Yellow, count, color.Reset, color.Green, lastVolume, color.Reset)
}
