package drinks

type Drink struct {
	Name       string
	Volume     float64
	Count      int
	LastVolume float64
	*DrinkInfo
}

type DrinkInfo struct {
	Name          string
	Type          string
	Alc           float64
	AviableVolume []float64
	Prices        []float64
	Description   string
	Taste         string
}

type Coctail struct {
	Drink
	NecessaryTings []string
	Ingredients    []string
	Grammar        []float64
	Description    string
	Instruction    string
}
