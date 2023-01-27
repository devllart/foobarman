package drinks

import (
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
)

func (drink Drink) Show() {
	fmtc.Printf(texts.ShowDrinkInBar, drink.Name, drink.Info.Alc, drink.Volume, drink.Count, drink.LastVolume)
	drink.Info.PrettyDescription()
}
