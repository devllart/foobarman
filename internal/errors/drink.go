package errors

import "fmt"

func NotExistDrink(name string) {
	err := fmt.Errorf("Напитка %s не существует", name)
	panic(err)
}

func NoVolumeOfDrink(name string, volume float64) {
	err := fmt.Errorf("%s с объёмом %g не существует", name, volume)
	panic(err)
}
