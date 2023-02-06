package mapsi

func (mapsi *Mapsi[T]) Data() map[string]*T {
	data := map[string]*T{}

	for _, key := range mapsi.Keys {
		data[key] = mapsi.GetValue(key)
	}

	return data
}

func (mapsi *Mapsi[T]) Len() int {
	return len(mapsi.Keys)
}
