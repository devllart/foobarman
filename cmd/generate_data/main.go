package main

import (
	"devllart/foobarman/internal/products"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var goCode = `package products

var AllCoctail = map[string]Coctail{%s
}`

func main() {
	coctails := []products.Coctail{}
	coctailsStruct := ``
	file, err := ioutil.ReadFile("data/coctails.json")
	if err != nil {
		panic(err.Error())
	}

	if err := json.Unmarshal([]byte(file), &coctails); err != nil {
		panic(err.Error())
	}

	for _, coctail := range coctails {
		grammars := ""
		coctail.Name = strings.Title(coctail.Name)
		coctail.Price = coctail.GetPrice()

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
    },`, replaceBitch(coctail.Name), replaceBitch(coctail.Name), joinBitch(coctail.Ingredients), grammars, joinBitch(coctail.Units), replaceBitch(coctail.Description), replaceBitch(coctail.Instruction))
	}

	ioutil.WriteFile("internal/products/coctails_data.go", []byte(fmt.Sprintf(goCode, coctailsStruct)), 0644)
}

func replaceBitch(str string) string {
	str = strings.ReplaceAll(str, "\"", "\\\"")
	str = strings.ReplaceAll(str, "\n", "\\n")
	return str
}

func joinBitch(strSlice []string) string {
	resStr := ""
	for i, str := range strSlice {
		if i > 0 {
			resStr += "\", \""
		}
		resStr += replaceBitch(str)
	}
	return resStr
}
