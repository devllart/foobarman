package structs

/**
 * Gloabal struct for packages products
 */

type Command struct {
	Aliases     []string
	Description string
}

type ProductInfo struct {
	Name          string
	Type          string
	Alc           float64
	AviableVolume []float64
	Prices        []float64
	Description   *string
	Taste         *string
}

type Product struct {
	Name       string
	Volume     float64
	Count      int
	LastVolume float64
	*ProductInfo
}

type Coctail struct {
	NecessaryTings        []string
	NecessaryTingsAmounts []string
	NecessaryTingsUnits   []string
	Ingredients           []string
	Grammar               []float64
	Description           string
	Instruction           string
	Price                 float64
	Name                  string
	Units                 []string
}
