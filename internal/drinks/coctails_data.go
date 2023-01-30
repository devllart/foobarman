package drinks

import "devllart/foobarman/internal/texts"

var AviableCoctail = map[string]Coctail{
	"Апероль Шприц": {
		NecessaryTings: []string{"Бокал для вина", "Джиггер", "Коктейльная ложка", "Трубочки"},
		Ingredients:    []string{"Апероль", "Просекко", "Содовая", "Апельсин", "Лёд в кубиках"},
		Grammar:        []float64{0.100, 0.100, 0.020, 0.040, 0.060},
		Description:    texts.DescCoctailAperrolShprits,
		Instruction:    texts.InstCoctailAperrolShprits,
	},

	"Пина Колада": {
		NecessaryTings: []string{"Харрикейн", "Джиггер", "Слайсер для ананаса", "Пресс для цитруссовых", "Совок для льда", "Трубочки", "Коктейльная шпажка", "Соковыжималка", "Блендер"},
		Ingredients:    []string{"Белый ром", "Тёмный ром", "Кокосовый сироп", "Лаймовый сок", "Ананас", "Ананасовые листья", "Коктельная вишня крассная", "Дроблённый лёд"},
		Grammar:        []float64{0.060, 0.015, 0.040, 0.020, 1, 1, 0.005, 0.060},
		Instruction:    texts.InstCoctailPinaColada,
	},

	"Зомби": {
		Ingredients: []string{"Тёмный ром", "Сверхкрепкий ром", "Сироп корицы", "Грейпфрутовый сок", "Лаймовый сок", "Грейпфрут", "Мята", "Дроблёный лёд", "Лёд в кубиках"},
	},

	"Май Тай": {
		Ingredients: []string{"Белый ром", "Тёмный ром", "Золотой ром", "Трипл сек", "Миндальный сироп", "Гренандин", "Ананасовый сок", "Лаймовый сок", "Апельсиновый сок", "Ананасовое пюре", "Ананасовые листья", "Тростниковый сахар в кубиках", "Лаймовая цедра", "Лёд в кубиках"},
	},

	"Ураган": {
		NecessaryTings: []string{"Харрикейн", "Шейкер", "Стрейнер", "Джиггер", "Пресс для цитрусовых", "Трубочки"},
		Ingredients:    []string{"Белый ром", "Тёмный ром", "Сироп маракуйи", "Гренандин", "Анансовый сок", "Апельсиновый сок", "Лаймовый сок", "Лимон", "Лёд в кубиках"},
		Grammar:        []float64{0.030, 0.030, 0.015, 0.005, 0.060, 0.060, 0.030, 0.020, 0.400},
		Description:    texts.DescCoctailHurricane,
		Instruction:    texts.InstCoctailHurricane,
	},

	"Ром кола": {
		NecessaryTings: []string{"Хайбол", "Пресс для цитрусовых", "Коктейльная ложка", "Джиггер", "Трубочки"},
		Ingredients:    []string{"Белый ром", "Лаймовый сок", "Кола", "Лайм", "Лёд в кубиках"},
		Grammar:        []float64{0.050, 0.010, 0.140, 0.020, 0.180},
		Description:    texts.DescCoctailRumCola,
		Instruction:    texts.InstCoctailRumCola,
	},

	"Тёмный ром с колой": {
		NecessaryTings: []string{"Хайбол", "Пресс для цитрусовых", "Коктейльная ложка", "Джиггер", "Трубочки"},
		Ingredients:    []string{"Тёмный ром", "Лаймовый сок", "Кола", "Лайм", "Лёд в кубиках"},
		Grammar:        []float64{0.050, 0.010, 0.140, 0.020, 0.180},
		Description:    texts.DescCoctailDarkRumWithCola,
		Instruction:    texts.InstCoctailDarkRumWithCola,
	},

	"Мохито": {
		NecessaryTings: []string{"Хайлбол", "Мадлер", "Джиггер", "Коктейльная ложка", "Трубочка"},
		Ingredients:    []string{"Белый ром", "Сахарный сироп", "Содовая", "Лайм", "Мята", "Дроблённый лёд"},
		Grammar:        []float64{0.050, 0.015, 0.100, 0.080, 0.003, 0.200},
		Description:    texts.DescCoctailMohito,
		Instruction:    texts.InstCoctailMohito,
	},

	"Маргарита": {
		NecessaryTings: []string{"Бокал маргарита", "Шейкер", "Стрейнер", "Джиггер", "Пресс для цитрусовых"},
		Ingredients:    []string{"Серебренная текила", "Трипл сек Fruko Schulz", "Сахарный сироп", "Лаймовый сок", "Лайм", "Соль", "Лёд в кубиках"},
		Grammar:        []float64{0.050, 0.025, 0.010, 0.030, 0.010, 0.002, 0.200},
		Description:    texts.DescCoctailMargarita,
		Instruction:    texts.InstCoctailMargarita,
	},

	"Негрони": {
		Ingredients: []string{"Лондонский сухой джин", "Красный вермут", "Красный биттер", "Апельсиновая цедра", "Лед в кубиках"},
		Grammar:     []float64{0.030, 0.030, 0.030, 1, 0.120},
		Description: texts.DescCoctailNegroni,
		Instruction: texts.InstCoctailNegroni,
	},

	"Драй Мартини": {
		Ingredients: []string{"Лондонский сухой джин", "Сухой вермут", "Оливки", "Лед в кубиках"},
		Grammar:     []float64{0.075, 0.015, 3, 0.300},
		Description: texts.DescCoctailDryMartini,
		Instruction: texts.InstCoctailDryMartini,
	},
}
