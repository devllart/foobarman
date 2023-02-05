package drinks

import (
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"
)

func New(name string, volume float64, count int) Drink {
	info, exitst := AviableDrinks[name]
	if exitst == false {
		fmtc.Perrorf(texts.ErrorNotExistDrink, name)
	}
	if funcs.Contains(info.AviableVolume, volume) == false {
		fmtc.Perrorf(texts.ErrorNotVolumeOfDrink, name, volume)
	}

	return Drink{name, volume, count, GetLastVolume(info.TypeVolume(), count, volume), &info}
}

func GetLastVolume(typeVolume string, count int, volume float64) float64 {
	if typeVolume != ".л" {
		return volume * float64(count)
	}

	return volume
}