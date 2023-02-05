package mapsi

type Mapsi[T any] struct {
	Keys   []string
	Values []T
}
