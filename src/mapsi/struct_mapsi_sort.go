package mapsi

import "sort"

func (mapsi *Mapsi[T]) SortLongestKeysFirst() Mapsi[T] {
	pairs := mapsi.ToPairs()

	sort.SliceStable(pairs, func(i, j int) bool {
		return len(pairs[i].Key) > len(pairs[j].Key)
	})
	return PairsToMapsi(pairs)
}
