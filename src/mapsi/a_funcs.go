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
