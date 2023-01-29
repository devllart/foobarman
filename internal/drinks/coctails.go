package drinks

import "sort"

func init() {
	for _, coctail := range AviableCoctail {
		sort.Sort(sort.StringSlice(coctail.Ingredients))
	}
}
