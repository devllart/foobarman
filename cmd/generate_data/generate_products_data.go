package main

import (
	anames "devllart/foobarman/internal/a_names"
	"devllart/foobarman/internal/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func getProducts() []structs.ProductInfo {
	var jsonFileWithData = anames.JsonDataDir + "/products.json"
	products := []structs.ProductInfo{}
	file, err := ioutil.ReadFile(jsonFileWithData)
	if err != nil {
		panic(err.Error())
	}

	if err := json.Unmarshal([]byte(file), &products); err != nil {
		panic(err.Error())
	}

	return products
}

func generateProductsData() {
	var outFile = "internal/products/products_data.go"
	var goCode = `package products

import (
	"devllart/foobarman/internal/structs"
  %s
)

var AvailableProducts = map[string]structs.ProductInfo{%s
}`

	productsStruct := ``
	if env, _ := os.LookupEnv("GENERATE"); env == "empty" {
		ioutil.WriteFile(outFile, []byte(fmt.Sprintf(goCode, "", productsStruct)), 0644)
		return
	}

	for _, product := range getProducts() {
		name := replaceBitch(product.Name)
		typeProduct := replaceBitch(product.Type)
		alc := product.Alc
		// taste := product.Taste
		taste := "GetTaste(\"" + product.Type + "\")"
		// taste := product.GetTaste(product)
		volumes := ""
		prices := ""
		description := saveText("DescProduct", replaceBitch(*product.Description))

		for i, vol := range product.AviableVolume {
			if i > 0 {
				volumes += ", "
			}
			volumes += fmt.Sprintf("%.3f", vol)
		}

		for i, price := range product.Prices {
			if i > 0 {
				prices += ", "
			}
			prices += fmt.Sprintf("%.2f", price)
		}

		productsStruct += fmt.Sprintf(`
    "%s": {
      Name:          "%s",
      Type:          "%s",
      Alc:           %.2f,
      Taste:         %s,
      AviableVolume: []float64{%s},
      Prices:        []float64{%s},
      Description:   &%s,
    },`, strings.Title(name), name, typeProduct, alc, taste, volumes, prices, description)
	}

	importTexts := `"devllart/foobarman/internal/texts"`

	ioutil.WriteFile(outFile, []byte(fmt.Sprintf(goCode, importTexts, productsStruct)), 0644)
}
