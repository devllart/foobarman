package products

/**
 * My standart flow
 */

var ProductsStandartFlow = map[string]float64{
	// "Водка":                     0.1,
	// "Ром":                       0.1,
	// "Тёмный ром":                0.1,
	// "Белый ром":                 0.1,
	// "Содовая":                   0.25,
	// "Фрукт":                     0.25,
	// "Цитрус":                    0.25,
	// "Кола":                      0.3,
	"Оливки":             0.001,
	"Апельсиновая цедра": 0.001,
	// "Листья лайма":              1,
	// "Лед в кубиках":             1,
	// "Лёд в кубиках":             1,
	// "Дроблённый лёд":            1,
	// "Дроблённый лед":            1,
	// "Коктельная вишня крассная": 1,
}

func GetStandartFlow(drinkType string) float64 {
	if flow, exist := ProductsStandartFlow[drinkType]; exist == true {
		return flow
	}

	return 0.1
}
