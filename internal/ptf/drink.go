package ptf

import "fmt"

func DrinkShow(name string, volume float64, count int, lastVolume float64) {
	fmt.Printf("   %s (%g .л) %dX | в последней бутылке осталось %g .л\n", name, volume, count, lastVolume)
}
