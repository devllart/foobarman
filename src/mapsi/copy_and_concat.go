package mapsi

func (mapsi *Mapsi[T]) Concat(secondMapsi Mapsi[T]) {
	for i, key := range secondMapsi.Keys {
		mapsi.SetValue(key, secondMapsi.Values[i])
	}
}

func (mapsi *Mapsi[T]) CopyElemntsOfMap(maps map[string]T, start, end int) {
	if start < 0 || end >= len(maps) {
		return
	}

	i := 0
	for name, val := range maps {
		if i >= end {
			return
		}
		if i >= start {
			mapsi.SetValue(name, val)
		}

		i += 1
	}
}
