package mapsi

func MapsiStringsToSliceKeysFirst(mapsi *Mapsi[string]) []string {
	slice := []string{}

	for i, key := range mapsi.Keys {
		slice = append(slice, key, mapsi.Values[i])
	}

	return slice
}

func MapsiStringsToSliceValuesFirst(mapsi *Mapsi[string]) []string {
	slice := []string{}

	for i, key := range mapsi.Keys {
		slice = append(slice, mapsi.Values[i], key)
	}

	return slice
}
