package drinks

import "sort"

var AviableCoctail = map[string]Coctail{
	"Апероль Шприц": {
		Ingredients: []string{"Апероль", "Содовая", "Просекко"},
	},
}

func init() {
	for _, coctail := range AviableCoctail {
		sort.Sort(sort.StringSlice(coctail.Ingredients))
	}
}
