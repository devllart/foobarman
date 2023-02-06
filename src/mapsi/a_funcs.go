package mapsi

func New[T any](maps map[string]T) Mapsi[T] {
	keys := []string{}
	values := []T{}

	for key, value := range maps {
		keys = append(keys, key)
		values = append(values, value)
	}

	mapsi := Mapsi[T]{
		Keys:   keys,
		Values: values,
	}

	return mapsi
}

func (mapsi *Mapsi[T]) SetValue(key string, value T) {
	for i, keyMapsi := range mapsi.Keys {
		if key == keyMapsi {
			mapsi.Values[i] = value
			return
		}
	}

	mapsi.Keys = append(mapsi.Keys, key)
	mapsi.Values = append(mapsi.Values, value)
}

func (mapsi *Mapsi[T]) GetValue(key string) *T {
	for i, keyMapsi := range mapsi.Keys {
		if key == keyMapsi {
			return &mapsi.Values[i]
		}
	}
	return nil
}

func (mapsi Mapsi[T]) Get(key string) *T {
	for i, keyMapsi := range mapsi.Keys {
		if key == keyMapsi {
			return &mapsi.Values[i]
		}
	}
	return nil
}

func (mapsi *Mapsi[T]) GetValueOfIndex(index int) *T {
	if index < len(mapsi.Values) {
		return &mapsi.Values[index]
	}

	return nil
}

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

func (mapsi *Mapsi[T]) Concat(secondMapsi Mapsi[T]) {
	for i, key := range secondMapsi.Keys {
		mapsi.SetValue(key, secondMapsi.Values[i])
	}
}

func (mapsi *Mapsi[T]) CopyElemntsOfMap(maps map[string]T, start, end int) {
	if start < 0 || end >= len(maps) {
		panic("Out of range map")
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
