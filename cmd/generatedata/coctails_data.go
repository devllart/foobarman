package main

import (
	"devllart/foobarman/internal/structs"
	"fmt"
	"strconv"
	"strings"
)

func coctailsData() string {
	coctailsData := []structs.Coctail{}
	getJsonData("coctails", &coctailsData)

	coctailsStruct := ``
	for _, coctail := range coctailsData {
		grammars := ""

		name := replaceBitch(strings.Title(coctail.Name))
		ingredients := joinBitch(coctail.Ingredients)
		units := joinBitch(coctail.Units)
		price := coctail.GetPrice()
		description := saveText("DescCoctail", replaceBitch(coctail.Description))
		instruction := saveText("InstCoctail", replaceBitch(coctail.Instruction))

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
      Description: %s,
      Instruction: %s,
      Price: %.2f,
    },`, name, ingredients, grammars, units, description, instruction, price)
	}
	return coctailsStruct
}
