package drinks

import (
	"devllart/foobarman/src/fmtc"
	"fmt"
)

func (coctail Coctail) Show() {
	fmtc.Printf("%R%s%C: ", coctail.Name)

	for i, ingredient := range coctail.Ingredients {
		if i > 0 {
			fmtc.Printf(", ")
		}
		fmtc.Printf("%B%s%C", ingredient)
	}
	fmt.Println()
}
