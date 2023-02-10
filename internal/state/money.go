package state

import "devllart/foobarman/src/fmtc"

var Money float64 = 300.33 // Player money

func SubMoney(amount float64) {
	if !InfiniteMoney {
		Money -= amount
	}
}

func SubtractFromSalary(money float64) string {
	if money < 0.01 {
		return fmtc.Sprintf("%YТебе хотели дать штраф, но похоже ты почти %Rбанкрот%Y так что тебя пожелели%C\n")
	} else {
		Money -= money
		return fmtc.Sprintf("%YЗа это проступок будешь платить из своего кармана: %R-%.2f$%C\n", money)
	}
}
