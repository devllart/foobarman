package drinks

type DrinkInfo struct {
	Description   string
	AviableVolume []float64
}

var AviableDrinks = map[string]DrinkInfo{
	"Биттер": {
		Description:   "",
		AviableVolume: []float64{1.5},
	},
	"Апероль": {
		Description:   "Знаменитый итальянский аперитив производят с 1919 года по секретной рецептуре братьев Барбьери из Падуи.\nВ состав напитка входит более 30 компонентов, включая цедру горьких апельсинов, травы и ревень.",
		AviableVolume: []float64{0.7},
	},
}
