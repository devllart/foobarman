package main

import (
	"github.com/gopherjs/gopherjs/js"
)

type Drink struct {
	Name string
}

func main() {
	js.Global.Get("myButton").Call("addEventListener", "click", func() {
		go func() {
			cmd := js.Global.Get("myInput").Get("value")
			drink := Drink{Name: cmd.String()}
			js.Global.Get("myContent").Set("innerText", drink.Name)
		}()
	})
}
