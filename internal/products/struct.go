package drinks

type Product struct {
	Name       string
	Volume     float64
	Count      int
	LastVolume float64
	*ProductInfo
}

type ProductInfo struct {
	Name          string
	Type          string
	Alc           float64
	AviableVolume []float64
	Prices        []float64
	Description   string
	Taste         *string
}

type Coctail struct {
	NecessaryTings []string
	Ingredients    []string
	Grammar        []float64
	Description    string
	Instruction    string
	Price          float64
	Name           string
}
