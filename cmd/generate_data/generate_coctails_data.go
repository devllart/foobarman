package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var jsonFileWithData = "data/coctails.json"
var outFile = "internal/products/coctails_data.go"

var goCode = `package products

var AllCoctail = map[string]Coctail{%s
}`

func generateCoctailsData() {
	coctailsStruct := ``

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
    "%s": {
      Name: "%s",
      Ingredients: []string{"%s"},
      Grammar: []float64{%s},
      Units: []string{"%s"},
      Description: "%s",
      Instruction: "%s",
      Price: %.2f,
    },`, name, name, ingredients, grammars, units, description, instruction, price)
	}

	ioutil.WriteFile(outFile, []byte(fmt.Sprintf(goCode, coctailsStruct)), 0644)
}
