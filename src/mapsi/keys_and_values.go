package mapsi

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
