package main

import (
	drinks "devllart/foobarman/internal/products"
	"encoding/json"
	"io/ioutil"
	"strings"
)

func main() {
	// 	var goCode = `package data

	// var AllCoctail = map[string]d
	// `
	coctails := []drinks.Coctail{}
	file, err := ioutil.ReadFile("data/coctails.json")
	if err != nil {
		panic(err.Error())
	}

	if err := json.Unmarshal([]byte(file), &coctails); err != nil {
		panic(err.Error())
	}

	for _, coctail := range coctails {
		coctail.Name = strings.Title(coctail.Name)
		coctail.Price = coctail.GetPrice()

		// AllCoctail[coctail.Name] = coctail
		// if i < CountAvailableCoctail {
		// 	AviableCoctail[coctail.Name] = coctail
		// }
	}
}
