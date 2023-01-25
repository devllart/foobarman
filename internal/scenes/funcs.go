package scenes

import (
	"devllart/foobarman/internal/funcs"
	"devllart/foobarman/internal/state"
	"fmt"
)

func Show(scene func()) {
	funcs.CliClear()
	scene()

	fmt.Printf("\n\nБармен %s   денег: %.2f $\n\n", state.Name, state.Money)
	if state.Info != "" {
		fmt.Printf("%s\n", state.Info)
	}
	fmt.Print("\nДоступные команды (регистр не важен):\n\n")
	fmt.Print("exit или выйти — выйти из игры\n")
	fmt.Print("desc, description или описание — спрятать/показать описание\n")
	if CurrentIs(Store) {
		fmt.Print("ok или ок — закончить покупку алкоголя\n")
	} else {
		fmt.Print("store или магазин — пойти в магазин за ингридиентами\n")
	}
}

func CurrentIs(scene func()) bool {
	return fmt.Sprintf("%v", state.Scene) == fmt.Sprintf("%v", scene) // Совсем немного говнокода
}
