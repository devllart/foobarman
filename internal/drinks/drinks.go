package drinks

import (
	"encoding/json"
	"io/ioutil"
)

var AviableDrinks = map[string]DrinkInfo{}

func init() {
	file, err := ioutil.ReadFile("data/drinks.json")
	if err != nil {
		panic(err)
	}

	data := []DrinkInfo{}

	err = json.Unmarshal([]byte(file), &data)

	if err != nil {
		panic(err)
	}
	for i := range data {
		drink := data[i]
		AviableDrinks[drink.Name] = drink
	}
}
