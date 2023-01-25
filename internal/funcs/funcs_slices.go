package funcs

func Contains[T string | float64 | int](s []T, str T) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
