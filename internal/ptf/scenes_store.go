package ptf

import (
	"devllart/foobarman/internal/texts"
	"fmt"
)

func StoreHello(name string) {
	fmt.Printf(texts.StoreHello, name)
}
