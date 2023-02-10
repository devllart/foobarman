package products

import (
	"devllart/foobarman/internal/data"
	"devllart/foobarman/internal/structs"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"
)

/**
 * Global funcs for package products
 */

// For added new product to bar
func New(name string, volume float64, count int) structs.Product {
	info, exitst := AvailableProducts[name]
	if exitst == false {
		fmtc.Perrorf(texts.ErrorNotExistProduct, name)
	}
	if funcs.Contains(info.AviableVolume, volume) == false {
		fmtc.Perrorf(texts.ErrorNotVolumeOfProduct, name, volume)
	}

	product := structs.Product{
		Name:        name,
		Volume:      volume,
		Count:       count,
		LastVolume:  volume,
		ProductInfo: &info,
	}

	return product
}

func GetTaste(productType string) *string {
	return data.NewTastes.GetValue(productType, texts.Unknown)
}
