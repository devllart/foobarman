package drinks

import (
	"devllart/foobarman/internal/config"
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"
	"fmt"
	"strings"
)

func (coctail *Coctail) PrettyDescription() {
	if config.ShowDescription {
		if coctail.Description == "" {
			fmtc.Printf(texts.NotSayAboutCoctail, funcs.Indent(4), coctail.Name)
		} else {
			lines := strings.Split(coctail.Description, "\n")
			for _, line := range lines {
				fmtc.Printf("%s%s\n", funcs.Indent(4), line)
			}
			fmt.Println()
		}
	}
}

func (coctail *Coctail) PrettyInstruction() {
	if config.ShowInstruction {
		if coctail.Instruction == "" {
			fmtc.Printf(texts.UnknownRecipes, funcs.Indent(4), coctail.Name)
		} else {
			fmtc.Printf(texts.RecipesCoctail, funcs.Indent(4), coctail.Name)
			lines := strings.Split(coctail.Instruction, "\n")
			for i, line := range lines {
				if i == 0 {
					continue
				}
				fmtc.Printf("%s%B%d.%C %s\n", funcs.Indent(4), i, line)
			}
		}
		fmt.Println()
	}
}

func (coctail *Coctail) Show() {
	fmtc.Printf("%R%s%C — ингредиенты: ", coctail.Name)

	for i, ingredient := range coctail.Ingredients {
		if i > 0 {
			fmtc.Printf(", ")
		}
		if i < len(coctail.Grammar) {
			typeVolume := ".л"
			if DrinksTypesVolume[ingredient] != "" {
				typeVolume = DrinksTypesVolume[ingredient]
			}

			fmtc.Printf("%B%s%C (%.3f%s)", ingredient, coctail.Grammar[i], typeVolume)
		} else {
			fmtc.Printf("%B%s%C", ingredient)
		}
	}
	fmt.Println()
}
