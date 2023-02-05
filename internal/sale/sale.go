package sale

import (
	"devllart/foobarman/internal/drinks"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/src/fmtc"
	"math/rand"
	"fmt"
)

func Sale() {
	if !state.BarOpen {
		return
	}

	fmt.Println()
	if state.NotSaler {
		state.NotSaler = false
		i := rand.Intn(len(drinks.AviableCoctail))
		coctail := *drinks.MapsiAviableCoctail.GetValueOfIndex(i)
		state.Order = coctail
		fmtc.Printf("%YУ вас новый покупатель и он заказал %B%s%C\n", coctail.Name)
	} else if state.CoctailReady {
		state.NotSaler = true
		state.CoctailReady = false

		if state.YourCoctail.Name == state.Order.Name {	
			state.Money += state.Order.Price
			fmtc.Printf("%YОтлично клиент доволен. %Y Вы получили %.2f$ %C\n", state.Order.Price)
		} else {
			fmtc.Printf("%RКлиент заказал не это, он ушёл накричав на вас%C\n")
		}
	} else {
		fmtc.Printf("%YВаш покупатель до сих пор ждёт %B%s%C\n", state.Order.Name)
	}
}