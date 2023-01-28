package drinks

import (
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
)

func (drink Drink) Show() {
	fmtc.Printf(texts.ShowDrinkInBar, drink.Name, drink.Info.Alc, drink.Volume, drink.Info.GetTypeVolume(), drink.Count, "в последней бутылке осталось", drink.LastVolume, drink.Info.GetTypeVolume())
	drink.Info.PrettyDescription()
}

func (drink Drink) StandartFlow() float64 {
	if flow, exist := DrinksStandartFlow[drink.Info.Type]; exist == true {
		return flow
	}

	return 0.1
}

func (drink Drink) TypeVolume() string {
	return drink.Info.GetTypeVolume()
}

func (drink *Drink) SubVolume() error {
	newVol := drink.LastVolume - drink.StandartFlow()
	newCount := drink.Count

	for true {
		if newVol > 0 {
			drink.Count = newCount
			drink.LastVolume = newVol
			return nil
		} else {
			newCount -= 1
			newVol += drink.Volume
			if newCount < 1 {
				return fmtc.Errorf("%RНе хватает объёма%C")
			}
		}
	}

	return nil
}
