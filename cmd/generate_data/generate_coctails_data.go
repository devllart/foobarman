package main

import (
	"devllart/foobarman/internal/products"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func getCoctails() []products.Coctail {
	var jsonFileWithData = "data/coctails.json"

	coctails := []products.Coctail{}
	file, err := ioutil.ReadFile(jsonFileWithData)
	if err != nil {
		panic(err.Error())
	}

	if err := json.Unmarshal([]byte(file), &coctails); err != nil {
		panic(err.Error())
	}

	return coctails
}

func generateCoctailsData() {
	var outFile = "internal/products/coctails_data.go"
	var goCode = `package products

var AllCoctail = []Coctail{%s
}`

	coctailsStruct := ``
	if env, _ := os.LookupEnv("GENERATE"); env == "empty" {
		ioutil.WriteFile(outFile, []byte(fmt.Sprintf(goCode, coctailsStruct)), 0644)
		return
	}

	for _, coctail := range getCoctails() {
		grammars := ""

		name := replaceBitch(strings.Title(coctail.Name))
		ingredients := joinBitch(coctail.Ingredients)
		units := joinBitch(coctail.Units)
		price := coctail.GetPrice()
		description := replaceBitch(coctail.Description)
		instruction := replaceBitch(coctail.Instruction)

		for i, gram := range coctail.Grammar {
			if i > 0 {
				grammars += ", "
			}
			grammars += strconv.Itoa(int(gram))
		}

		coctailsStruct += fmt.Sprintf(`
    {
      Name: "%s",
      Ingredients: []string{"%s"},
      Grammar: []float64{%s},
      Units: []string{"%s"},
      Description: "%s",
      Instruction: "%s",
      Price: %.2f,
    },`, name, ingredients, grammars, units, description, instruction, price)
	}

	ioutil.WriteFile(outFile, []byte(fmt.Sprintf(goCode, coctailsStruct)), 0644)
}
