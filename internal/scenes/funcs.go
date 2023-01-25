package scenes

import (
	"devllart/foobarman/internal/state"
	"devllart/foobarman/src/funcs"
	"fmt"
)

func Show(scene func()) {
	funcs.CliClear()
	scene()

	fmt.Printf("\nБармен %s   денег: %.2f $\n\n", state.Name, state.Money)

	fmt.Print("Доступные команды (регистр не важен):\n")
	fmt.Print("exit или выйти — выйти из игры\n")
	fmt.Print("desc, description или описание — спрятать/показать описание\n")
	if CurrentIs(Store) {
		fmt.Print("ok или ок — закончить покупку алкоголя\n")
	} else {
		fmt.Print("store или магазин — пойти в магазин за ингридиентами\n\n")
	}

	if state.Info != "" {
		fmt.Printf("\n%s", state.Info)
	}
}

func CurrentIs(scene func()) bool {
	return fmt.Sprintf("%v", state.Scene) == fmt.Sprintf("%v", scene) // Совсем немного говнокода
}
