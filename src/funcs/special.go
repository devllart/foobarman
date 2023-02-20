package funcs

func If[T any](condition bool, ret, ret2 T) T {
	if condition {
		return ret
	} else {
		return ret2
	}
}
