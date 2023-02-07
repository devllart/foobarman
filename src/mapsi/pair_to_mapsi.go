package mapsi

func PairsToMapsi[T any](pairs []Pair[T]) Mapsi[T] {
	mapsi := Mapsi[T]{}
	for _, pair := range pairs {
		mapsi.SetValue(pair.Key, pair.value)
	}

	return mapsi
}
