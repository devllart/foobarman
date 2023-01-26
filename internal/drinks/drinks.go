package drinks

import (
	"encoding/json"
	"io/ioutil"
)

var AviableDrinks = map[string]DrinkInfo{}

func init() {
	var err error
	var file []byte

	data := []DrinkInfo{}

	if file, err = ioutil.ReadFile("data/drinks.json"); err != nil {
		panic(err)
	}
	if err = json.Unmarshal([]byte(file), &data); err != nil {
		panic(err)
	}

	for i := range data {
		drink := data[i]
		drink.Valid()
		AviableDrinks[drink.Name] = drink
	}
}
