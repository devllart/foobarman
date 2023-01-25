package funcs

func Contains[T string | float64 | int](s []T, el T) bool {
	for _, v := range s {
		if v == el {
			return true
		}
	}

	return false
}

func LineContains[T string | float64 | int](el T, slice ...T) bool {
	for _, elFromSlice := range slice {
		if elFromSlice == el {
			return true
		}
	}

	return false
}
