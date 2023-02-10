package structs

import (
	"devllart/foobarman/internal/data"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
)

/**
 * Funcs for struct Producct
 */

func (product *Product) Show() {
	fmtc.Printf(texts.ShowProductInBar, product.Name, product.Alc, product.Volume, product.TypeVolume(), product.Count, product.LeftVolumeText(), product.GetLastVolume(), product.TypeVolume())
	product.PrettyDescription()
}

func (product *Product) StandartFlow() float64 {
	if flow, exist := data.ProductsStandartFlow[product.Type]; exist == true {
		return flow
	}

	return product.AviableVolume[0] / 25
}

func (product *Product) GetLastVolume() float64 {
	if product.TypeVolume() != ".л" {
		// fmt.Println(product)
		// panic("AAA")
		return product.Volume*float64(product.Count-1) + product.LastVolume
	}

	return product.LastVolume
}

func (product *Product) LeftVolumeText() string {
	if product.TypeVolume() != ".л" {
		return texts.TotalLeftVolume
	}

	return texts.LeftVolumeInLastBottle
}

func (product *Product) SubVolume(vol float64) error {
	newVol := product.LastVolume - vol
	newCount := product.Count

	for true {
		if newVol > 0 {
			product.Count = newCount
			product.LastVolume = newVol

			return nil
		} else {
			newCount -= 1
			newVol += product.Volume
			if newCount < 1 {
				return fmtc.Errorf(texts.ErrorNotEnoughtVolume, product.Name)
			}
		}
	}

	return nil
}

func (product *Product) AddCount(count int) {
	product.Count += count
}
