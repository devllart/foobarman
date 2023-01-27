package main

import (
	"devllart/foobarman/internal/drinks"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type CatlogNodes struct {
	CatlogNodes []Catlog
}

type Catlog struct {
	Name          string
	Type          string
	Taste         string
	Alc           int
	AviableVolume []float64
	Prices        []float64
	Description   string
}

func JsonRun() {
	file, err := ioutil.ReadFile("data/drinks.json")

	fmt.Println(err)

	data := []drinks.DrinkInfo{}

	err = json.Unmarshal([]byte(file), &data)

	fmt.Println(err)

	fmt.Println(data)
	// for i := 0; i < len(data.CatlogNodes); i++ {
	// 	fmt.Println("Product Id: ", data.CatlogNodes[i].Product_id)
	// 	fmt.Println("Quantity: ", data.CatlogNodes[i].Quantity)
	// }

}
