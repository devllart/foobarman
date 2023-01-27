package drinks

import (
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"
)

func New(name string, volume float64, count int) Drink {
	info, exitst := AviableDrinks[name]
	if exitst == false {
		fmtc.Perrorf(texts.NotExistDrinkError, name)
	}
	if funcs.Contains(info.AviableVolume, volume) == false {
		fmtc.Perrorf(texts.NotVolumeOfDrinkError, name, volume)
	}

	return Drink{name, volume, count, volume, info}
}
