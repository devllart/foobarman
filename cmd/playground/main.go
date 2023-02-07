package main

import (
	"devllart/foobarman/internal/products"
	"encoding/json"
	"io/ioutil"
)

func main() {
	data := products.AvailableProducts
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		panic(err.Error())
	}

	if err = ioutil.WriteFile("data/products.json", file, 0644); err != nil {
		panic(err.Error())
	}
}
