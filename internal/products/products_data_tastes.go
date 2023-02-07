package products

import (
	"devllart/foobarman/internal/texts"
	"devllart/foobarman/src/mapsi"
)

var Tastes = map[string]string{
	// Alco drinks
	"Пиво": "Хмеля и пятницы",
	// ** Much degeress
	"Водка":            "Очень горький",
	"Ирландский виски": "Ирландии",
	"Цитрусовая водка": "Очень горький и весьма цитрусовый",

	// Alco drinks | Rum
	"Ром":        "Сладко-горький, нотки высушенной травы переплетается с оттенками пряностей, табака, хорошо и выделанной кожи",
	"Белый ром":  "Сладко-горький, имеет приятный сладковатый привкус, после которого остаётся ненавязчивое послевкусие",
	"Тёмный ром": "Сладко-горький, можно почувствовать намёки на специи вместе с сильными нотами патоки или карамели",

	// Alco drinks | Vermouth
	"Красный вермут": "Cладкий и завораживающий, с яркими апельсиновыми акцентами, плавно сменяющимися легкой горчинкой.",
	"Вермут":         "Приятный, ненавязчивый, и одновременно выразительный сладковатый вкус",
	"Апероль":        "Горьковатый, содержит ноты апельсинов, трав и ароматических приправ",

	// Alco drinks | Wine
	"Вино":       "Cочное, шелковистое, насыщенное, мягкое",
	"Сухое вино": "Терпкое, танинное и довольно агрессивное на вкус",
	"Проссеко":   "Да что это вобще такое!",

	// Alco drinks | Bitters
	"Биттер":               "Как горький эль",
	"Эботтс Биттер":        "Горьковатый",
	"Красный биттер":       "Горький, с нотками шафрана, ангостуру и коломбо",
	"Ароматический биттер": "Ароматный",

	// Alco drinks | Gins
	"Лондонский сухой джин":    "Очень сухой :/",
	"Cухой джин":               "Ну... такой, немного сухой",
	"Амстердамский сухой джин": "Ну не такой сухой как лондонский",

	// No alco drinks
	// No alco drinks | likers and syrups
	"Гренандин":        "сладкий и терпкий",
	"Ликер":            "Сладкий",
	"Шоколадный ликер": "Шоколадный — кто бы мог подумать?",
	"Вишнёвый ликер":   "Приторный, вишнёвый",
	"Сахарный сироп":   "Приторный",
	"Ванильный сироп":  "Ванильного мороженного",
	"Мятный сироп":     "Приторный, мятный",
	"Кокосовый сироп":  "Приторный и кокосовый",

	"Кола":             "%RПраздник к нам приходит, праздник к нам приходит...%C",
	"Спрайт":           "Лаймовый",
	"Лимонад":          "Детства",
	"Апельсиновый сок": "Апельсина",
	"Ананасовый сок":   "Ананаса",
	"Вишнёвый сок":     "Вишни",
	"Клюквенный сок":   "Клюквы",
	"Лаймовый сок":     "Лайма",
	"Креплённое вино":  "По сровнению с ним сухой вино нифига не агрессивное)",

	"Вода":              "Безвкусный",
	"Газированная вода": "Газированной воды",
	"Обуховская":        "Обуховской",
	"Содовая":           "Содовой",

	// Addeds
	"Оливки": "Здоровой пищи — он тебе не знаком",

	// Fruits
	"Фрукт": "Фруктозы",
	"Лайм":  "Зелёного цитруса",
	"Лимон": "Цитруса",

	// Others
	"Листья лайма":  "Ну какой вкус может быть у листьев? Вкус листового чая, может, я не знаю",
	"Мята":          "Мятный — очень нравится кискам ;) ",
	"Лед":           "Безвкусный",
	"Дроблёный лед": "Безвкусный",
	"Лед в кубиках": "Безвкусный",
	"Соль":          "Соль",

	// Triple sec
	"Трипл сек fruko schulz": "Ликер с экзотическим, богатым и сложным вкусом с нотками цитрусовых.",

	// Invdividual
	"Водка Царская": "Царско-горький",

	// Unknow
	texts.Unknown: texts.Unknown,
}

var NewTastes = mapsi.New(Tastes)
