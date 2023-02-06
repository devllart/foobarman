package drinks

import (
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
)

func (drink *Drink) Show() {
	fmtc.Printf(texts.ShowDrinkInBar, drink.Name, drink.Alc, drink.Volume, drink.TypeVolume(), drink.Count, drink.LeftVolumeText(), drink.GetLastVolume(), drink.TypeVolume())
	drink.PrettyDescription()
}

func (drink *Drink) StandartFlow() float64 {
	if flow, exist := DrinksStandartFlow[drink.Type]; exist == true {
		return flow
	}

	return drink.AviableVolume[0] / 25
}

func (drink *Drink) GetLastVolume() float64 {
	if drink.TypeVolume() != ".л" {
		return drink.Volume*float64(drink.Count-1) + drink.LastVolume
	}

	return drink.LastVolume
}

func (drink *Drink) LeftVolumeText() string {
	if drink.TypeVolume() != ".л" {
		return texts.TotalLeftVolume
	}

	return texts.LeftVolumeInLastBottle
}

func (drink *Drink) SubVolume(vol float64) error {
	newVol := drink.LastVolume - vol
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
				return fmtc.Errorf(texts.ErrorNotEnoughtVolume)
			}
		}
	}

	return nil
}
