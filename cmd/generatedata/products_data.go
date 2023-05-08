package main

import (
	"devllart/foobarman/internal/structs"
	"fmt"
	"strings"
)

func productsData() string {
	productsData := []structs.ProductInfo{}
	getJsonData("products", &productsData)
	productsStruct := ``

	for _, product := range productsData {
		name := replaceBitch(product.Name)
		typeProduct := replaceBitch(product.Type)
		alc := product.Alc
		taste := saveText("Tastes", replaceBitch(*product.Taste))
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
      Taste:         &%s,
      AviableVolume: []float64{%s},
      Prices:        []float64{%s},
      Description:   &%s,
    },`, strings.Title(name), name, typeProduct, alc, taste, volumes, prices, description)
	}

	return productsStruct
}
