package mapsi

func New[T string | int | float32 | float64](maps map[string]T) Mapsi[T] {
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

	panic("Mapsi: KeyNotFound")
}

func (mapsi *Mapsi[T]) GetValue(key string) *T {
	for i, keyMapsi := range mapsi.Keys {
		if key == keyMapsi {
			return &mapsi.Values[i]
		}
	}
	return nil
}
