package drinks

import (
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"
)

func New(name string, volume float64, count int) Product {
	info, exitst := AvailableProducts[name]
	if exitst == false {
		fmtc.Perrorf(texts.ErrorNotExistProduct, name)
	}
	if funcs.Contains(info.AviableVolume, volume) == false {
		fmtc.Perrorf(texts.ErrorNotVolumeOfProduct, name, volume)
	}

	return Product{name, volume, count, GetLastVolume(info.TypeVolume(), count, volume), &info}
}

func GetLastVolume(typeVolume string, count int, volume float64) float64 {
	if typeVolume != ".л" {
		return volume * float64(count)
	}

	return volume
}