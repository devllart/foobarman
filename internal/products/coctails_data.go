package drinks

import "devllart/foobarman/internal/texts"

var AviableCoctailOld = map[string]Coctail{
	"Текила санрайз": {
		NecessaryTings: []string{"Хайбол", "Коктельная ложка", "Джиггер", "Трубочки"},
		Ingredients:    []string{"Серебренная текила", "Гренандин", "Апельсиновый сок", "Апельсин", "Лёд в кубиках"},
		Grammar:        []float64{0.050, 0.010, 0.150, 0.030, 0.180},
		Description:    texts.DescCoctailSexOnTheBeach,
		Instruction:    texts.InstCoctailSexOnTheBeach,
	},
	"Космополитен": {
		NecessaryTings: []string{"Коктельный бокал", "Шейкер", "Стрейнер", "Джиггер", "Горелка", "Нож для цедры", "Пресс для цитрусовых"},
		Ingredients:    []string{"Цитрусовая водка", "Трипл сек Fruko Schulz", "Клюквенный сок", "Лаймовый сок", "Апельсиновый сок", "Лёд в кубиках"},
		Grammar:        []float64{0.040, 0.020, 0.050, 0.010, 0.001, 0.200},
		Description:    texts.DescCoctailSexOnTheBeach,
		Instruction:    texts.InstCoctailSexOnTheBeach,
	},
	"Секс На Пляже": {
		NecessaryTings: []string{"Слинг", "Шейкер", "Стрейнер", "Джиггер", "Трубочки", "Коктельная шпажка"},
		Ingredients:    []string{"Водка", "Персиковый ликёр Fruko Schulz", "Клюквенный сок", "Ананасовый сок", "Коктейльная вишня крассная", "Лёд в кубиках"},
		Grammar:        []float64{0.050, 0.025, 0.040, 0.040, 0.015, 0.005, 0.350},
		Description:    texts.DescCoctailSexOnTheBeach,
		Instruction:    texts.InstCoctailSexOnTheBeach,
	},
	"Голубая лагуна": {
		NecessaryTings: []string{"Харрикейн", "Коктельная ложка", "Джиггер", "Трубочки"},
		Ingredients:    []string{"Водка", "Ликер блю кюрасао Fruko Schulz", "Спрайт", "Ананас", "Лёд в кубиках"},
		Grammar:        []float64{0.050, 0.020, 0.150, 0.030, 0.200},
		Description:    texts.DescCoctailBlueLagoon,
		Instruction:    texts.InstCoctailBlueLagoon,
	},
	"Дайкири": {
		NecessaryTings: []string{"Шампанское блюдце", "Шейкер", "Стрейнер", "Джиггер", "Пресс для цитрусовых"},
		Ingredients:    []string{"Белый ром", "Сахарный сироп", "Лаймовый сок", "Лёд в кубиках"},
		Grammar:        []float64{0.060, 0.015, 0.030, 0.200},
		Description:    texts.DescCoctailDaiquiri,
		Instruction:    texts.InstCoctailDaiquiri,
	},
	"Джин Тоник": {
		NecessaryTings: []string{"Хайбол", "Джиггер", "Коктейльная ложка", "Трубочки"},
		Ingredients:    []string{"Лондонский сухой джин", "Тоник", "Лайм", "Лёд в кубиках"},
		Grammar:        []float64{0.050, 0.150, 0.020, 0.180},
		Description:    texts.DescCoctailGinTonic,
		Instruction:    texts.InstCoctailGinTonic,
	},
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
		Ingredients: []string{"Лондонский сухой джин", "Красный вермут", "Красный биттер", "Апельсиновая цедра", "Лёд в кубиках"},
		Grammar:     []float64{0.030, 0.030, 0.030, 1, 0.120},
		Description: texts.DescCoctailNegroni,
		Instruction: texts.InstCoctailNegroni,
	},
	"Драй Мартини": {
		Ingredients: []string{"Лондонский сухой джин", "Сухой вермут", "Оливки", "Лёд в кубиках"},
		Grammar:     []float64{0.075, 0.015, 3, 0.300},
		Description: texts.DescCoctailDryMartini,
		Instruction: texts.InstCoctailDryMartini,
	},
}

var AviableCoctail = map[string]Coctail{}
