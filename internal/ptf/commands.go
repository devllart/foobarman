package ptf

import (
	"devllart/foobarman/internal/texts"
	"fmt"

	"github.com/TwiN/go-color"
)

func SelectIngredients() {
	fmt.Printf(texts.SelectIngredients, color.Yellow, color.Reset)
}
