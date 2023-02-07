package mapsi

func (mapsi *Mapsi[T]) ToPairs() []Pair[T] {
	pairs := []Pair[T]{}

	for i, key := range mapsi.Keys {
		pairs = append(pairs, Pair[T]{key, mapsi.Values[i]})
	}

	return pairs
}
