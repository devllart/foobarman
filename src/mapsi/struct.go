package mapsi

type Mapsi[T any] struct {
	Keys   []string
	Values []T
}

type Pair[T any] struct {
	Key   string
	value T
}
