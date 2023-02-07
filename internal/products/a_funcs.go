package products

import (
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"
)

/**
 * Global funcs for package products
 */

// For added new product to bar
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

// For show volume in last (or all) volume of product in the bar
func GetLastVolume(typeVolume string, count int, volume float64) float64 {
	if typeVolume != ".л" {
		return volume * float64(count)
	}

	return volume
}
