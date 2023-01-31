package mapsi

type Mapsi[T string | int | float32 | float64] struct {
	Keys   []string
	Values []T
}
