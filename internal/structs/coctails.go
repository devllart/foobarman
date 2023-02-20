package structs

import (
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/data"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"
	"fmt"
	"strings"
)

/**
 * Funcs for struct Coctail
 */

func (coctail *Coctail) Show() {
	fmtc.Printf("%R%s%C\n%sингредиенты: ", coctail.Name, funcs.Indent(4))

	for i, ingredient := range coctail.Ingredients {
		if i > 0 {
			fmtc.Printf(", ")
		}
		if i < len(coctail.Grammar) {
			fmtc.Printf("\n%s - %B%s%C (%.3f%s)", funcs.Indent(10), ingredient, coctail.Grammar[i], coctail.Units[i])
		} else {
			fmtc.Printf("\n%s - %B%s%C", funcs.Indent(10), ingredient)
		}

	}

	var price float64 = 0
	if config.Env == "production" {
		price = coctail.Price
	} else {
		price = coctail.GetPrice()
	}

	fmtc.Printf("\n%sцена: %Y%.2f$%C\n\n", funcs.Indent(4), price)
}

func (coctail *Coctail) PrettyDescription() {
	if config.ShowDescription {
		if coctail.Description == "" {
			fmtc.Printf(texts.NotSayAboutCoctail, funcs.Indent(15), coctail.Name)
		} else {
			lines := strings.Split(coctail.Description, "\n")
			for _, line := range lines {
				fmtc.Printf("%s%s\n", funcs.Indent(15), line)
			}
			fmt.Println()
		}
	}
}

func (coctail *Coctail) PrettyInstruction() {
	if config.ShowInstruction {
		if coctail.Instruction == "" {
			fmtc.Printf(texts.UnknownRecipes, funcs.Indent(15), coctail.Name)
		} else {
			fmtc.Printf(texts.RecipesCoctail, funcs.Indent(15), coctail.Name)
			lines := strings.Split(coctail.Instruction, "\n")
			for i, line := range lines {
				if i == 0 {
					continue
				}
				fmtc.Printf("%s%B%d.%C %s\n", funcs.Indent(15), i, line)
			}
		}
		fmt.Println()
	}
}

func (coctail *Coctail) GetPrice() float64 {
	var sumPrice float64 = 0.5
	for i, ingredient := range coctail.Ingredients {
		if price, exist := data.ProductsStandartTypesPrice[ingredient]; exist {
			var grammar float64 = 0.1
			if i < len(coctail.Grammar) {
				grammar = coctail.Grammar[i] * 0.001
			}

			sumPrice += price * grammar * (1.1 - GetStandartFlow(ingredient))
		} else {
			sumPrice += 0.5
		}
	}

	return sumPrice * float64(len(coctail.Ingredients)) * 0.3
}

func GetStandartFlow(productType string) float64 {
	if flow, exist := data.ProductsStandartFlow[productType]; exist == true {
		return flow
	}

	return 0.1
}
