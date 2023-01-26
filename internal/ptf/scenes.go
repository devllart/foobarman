package ptf

import "fmt"

func StandartCommands() {
	fmt.Print("Доступные команды (регистр не важен):\n")
	fmt.Print("exit или выйти — выйти из игры\n")
	fmt.Print("desc, description или описание — спрятать/показать описание\n")
}

func FinishShoopingCommand() {
	fmt.Print("ok или ок — закончить покупку алкоголя\n")
}

func StartShoopingCommand() {
	fmt.Print("store или магазин — пойти в магазин за ингридиентами\n\n")
}

func ShowBarmanStatus(name string, money float64) {
	fmt.Printf("\nБармен %s   денег: %.2f $\n\n", name, money)

}
