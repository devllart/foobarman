package sale

import (
	"devllart/foobarman/internal/products"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/src/fmtc"
	"devllart/foobarman/src/funcs"
	"math/rand"
)

func NewClient() {
	state.CurrentHistory = []string{}
	state.NotSaler = false

	index := rand.Intn(products.MapsiAvailableCoctail.Len())
	coctail := *products.MapsiAvailableCoctail.GetValueOfIndex(index)
	state.Order = coctail
	if scripts, exist := Scripts[coctail.Name]; exist && len(scripts) > 0 && rand.Intn(len(scripts)) == 0 {
		index := rand.Intn(len(scripts))
		Scripts[coctail.Name] = funcs.RemoveElementSliceOfIndex(scripts, index)
		history := scripts[index]

		state.CurrentHistory = append(state.CurrentHistory, history.StartHistory, history.Leave)
		fmtc.Printf(history.Acquaintance)
	} else {
		fmtc.Printf("%YУ вас новый покупатель и он заказал %B%s%C\n", coctail.Name)
	}
}

func CoctailReady() {
	state.NotSaler = true
	state.CoctailReady = false

	if state.YourCoctail.Name == state.Order.Name {
		state.Money += state.Order.Price
		if len(state.CurrentHistory) == 2 {
			fmtc.Printf(state.CurrentHistory[0])
		} else {
			fmtc.Printf("%YОтлично клиент доволен.")
		}

		fmtc.Printf("%Y Вы получили %.2f$ %C\n", state.Order.Price)

		state.CountVisitorsServed += 1
		if state.CountVisitorsServed >= state.Stage*state.Stage {
			state.Stage += 1
			countNewProducts := products.MapsiAvailableProducts.Len()
			products.AddAvailablesCoctail(state.Stage)
			countNewProducts = products.MapsiAvailableProducts.Len() - countNewProducts
			fmtc.Printf("\n%Y ! Этап пройден.\n%Y рецепты: %R+%d%C\n%Y продукты: %R+%d%C\n", state.Stage, countNewProducts)
		}
	} else {
		if len(state.CurrentHistory) == 2 {
			fmtc.Printf(state.CurrentHistory[1])
		} else {
			fmtc.Printf("%RНо клиент заказал не это, он ушёл накричав на вас%C\n")
			text := state.SubtractFromSalary(state.Money / 97)
			fmtc.Printf(text)
		}
	}
}
