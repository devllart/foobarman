package funcs

func Contains[T string | float32 | float64 | int](s []T, el T) bool {
	for _, v := range s {
		if v == el {
			return true
		}
	}

	return false
}

func LineContains[T string | float32 | float64 | int](el T, slice ...T) bool {
	for _, elFromSlice := range slice {
		if elFromSlice == el {
			return true
		}
	}

	return false
}

func IndexOf[T string | float32 | float64 | int](val T, sliceMap []T) int {
	for i, v := range sliceMap {
		if v == val {
			return i
		}
	}
	return -1
}

func IndexOfOrNull[T string | float32 | float64 | int](val T, sliceMap []T) int {
	index := IndexOf(val, sliceMap)
	if index == -1 {
		return 0
	}
	return index
}

func RemoveElementSliceOfIndex[T any](s []T, index int) []T {
	ret := make([]T, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func SlicesEqual(firstSlice, secondSlice []string) bool {
	if len(firstSlice) != len(secondSlice) {
		return false
	}
	for i, val := range firstSlice {
		if secondSlice[i] != val {
			return false
		}
	}
	return true
}

func SlicesSortAndEqual(firstSlice, secondSlice []string) bool {
	first := append([]string{}, firstSlice...)
	second := append([]string{}, secondSlice...)
	return SlicesEqual(first, second)
}
