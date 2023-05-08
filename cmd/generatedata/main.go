package main

import (
	anames "devllart/foobarman/internal/a_names"
	"devllart/foobarman/src/funcs"
	"fmt"
	"io/ioutil"
	"os"
)

var goCode = `package generateddata

import (
	"devllart/foobarman/internal/structs"
  %s
)

var %s = %s{%s
}`

func goCodeGenerate(dep, name, typename, data string) {
	var outFile = anames.GameDataDir + "/" + funcs.ToSnakeCase(name) + ".go"
	ioutil.WriteFile(outFile, []byte(fmt.Sprintf(goCode, dep, name, typename, data)), 0644)
}

func generateData(name, typename string, getData func() string) {
	if env, _ := os.LookupEnv("GENERATE"); env == "empty" {
		goCodeGenerate("", name, typename, "")
		return
	}

	importTexts := `"devllart/foobarman/internal/texts"`
	goCodeGenerate(importTexts, name, typename, getData())
}

func main() {
	generateData("AvailableProducts", "map[string]structs.ProductInfo", productsData)
	generateData("AllCoctail", "[]structs.Coctail", coctailsData)
}
