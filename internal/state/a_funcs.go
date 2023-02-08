package state

import (
	"devllart/foobarman/internal/products"
	"devllart/foobarman/src/fmtc"
	"fmt"
	"strconv"
	"strings"
)

func AddInfo(args ...any) {
	Info += fmt.Sprint(args...)
}

func AddInfof(textf string, args ...any) {
	Info += fmtc.Sprintf(textf, args...)
}

func ClearTemp() {
	Command = ""
	Args = []string{}
	Alerts = []string{}
	TempBool = false

	ProductsIds = []string{}
}

func Restart() {
	ClearTemp()
	Bar = []products.Product{}
	Money = 300.33
	CloseBar()
	Info = ""
	NotSaler = true
	TempBool = false
}

func SetCommand(text string) {
	Command = ""
	for i, word := range strings.Split(text, " ") {
		if i == 0 {
			Command += word + " "
		} else if _, err := strconv.ParseFloat(word, 64); err == nil {
			Args = append(Args, word)
		} else {
			Command += word + " "
		}
	}

	for len(Args) < 3 {
		Args = append(Args, "")
	}

	Command = Command[:len(Command)-1]

	if len(Command) > 0 && Command[len(Command)-1] == 13 {
		Command = Command[:len(Command)-1]
	}
}

func HandlerStatus() {
	if Money < 10 {
		SetStatusBad()
	} else {
		SetStatusNorm()
	}
}

func SetStatusBad() {
	Status = "Bad"
	products.NewTastes.SetValue("Лондонский сухой джин", "Очень сухой :/ — но твоя девушка суше ;) ")
}

func SetStatusNorm() {
	Status = "Norm"
	products.NewTastes.SetValue("Лондонский сухой джин", "Очень сухой :/")
}

func SubtractFromSalary(money float64) string {
	if money < 0.01 {
		return fmtc.Sprintf("%YТебе хотели дать штраф, но похоже ты почти %Rбанкрот%Y так что тебя пожелели%C\n")
	} else {
		Money -= money
		return fmtc.Sprintf("%YЗа это проступок будешь платить из своего кармана: %R-%.2f$%C\n", money)
	}
}

func ExitGame() {
	Run = false
}

func ToOpenBar() {
	BarOpen = true
}

func CloseBar() {
	BarOpen = false
}

func SubMoney(amount float64) {
	if !InfiniteMoney {
		Money -= amount
	}
}
