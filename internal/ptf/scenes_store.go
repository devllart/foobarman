package ptf

import (
	"devllart/foobarman/internal/texts"
	"fmt"

	"github.com/TwiN/go-color"
)

func StoreHello(name string) {
	fmt.Printf(texts.StoreHello, color.Red, color.Reset, name)
}
