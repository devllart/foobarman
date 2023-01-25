package drinks

type Drink struct {
	Name       string
	Volume     float64
	Count      int
	LastVolume float64
	Info       DrinkInfo
}

type DrinkInfo struct {
	Type          string
	Taste         string
	Alc           float64
	AviableVolume []float64
	Prices        []float64
	Description   string
}
