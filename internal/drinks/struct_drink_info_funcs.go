package drinks

import (
	"devllart/foobarman/config"
	"devllart/foobarman/src/fmtc"
	"strings"
)

func (drink DrinkInfo) PrettyDescription() {
	if config.ShowDescription {
		if drink.Description == "" {
			fmtc.Printf("      | О %s ничего неизвестно\n", drink.Name)
			return
		}

		lines := strings.Split(drink.Description, "\n")

		for _, line := range lines {
			fmtc.Printf("      | %s\n", line)
		}
	}
}

func (drink DrinkInfo) GetTaste() string {

	if taste, exist := Tastes[drink.Name]; exist == true {
		return taste
	}
	if taste, exist := Tastes[drink.Type]; exist == true {
		return taste
	}

	return "Неизвестен"
}
