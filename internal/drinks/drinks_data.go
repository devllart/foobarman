package drinks

import (
	"devllart/foobarman/internal/texts"
)

var AviableDrinks = map[string]DrinkInfo {
	"Don Julio Reposado": {
		Type:          "Текила",
		Alc:           40,
		AviableVolume: []float64{0.750, 1.750},
		Prices:        []float64{59.99, 114.99},
		Description:   `
Don Julio Reposado Tequila — это идеальная текила, выдержанная в бочках. 
С ароматами спелых косточковых фруктов, ванили и корицы каждый глоток приводит к шелковистому, теплому послевкусию. 
Наша текила янтарного цвета, выдержанная не менее 6 месяцев в бочках из американского белого дуба, изготовлена из 100% голубой агавы Вебера и идеально подходит для любого торжества. 
Дон Хулио был назван одной из самых популярных текил на церемонии вручения наград Drinks International 2020 года. 
Просто смешайте с соком грейпфрута, свежевыжатым соком лайма и нектаром агавы и налейте на лед, чтобы получить освежающий вкус Don Julio Paloma. 
Включает в себя одну бутылку текилы Reposado Tequila на 80 градусов объемом 375 мл. Пожалуйста, пейте ответственно. 
Аромат: манящий аромат с нотками спелых лимонных цитрусов и слоев специй с оттенками спелых косточковых фруктов. 
Вкус: невероятно мягкие и элегантные оттенки темного шоколада, ванили и легкой корицы. Послевкусие: шелковистое,`,
	},
	"Roca Patrón Silver": {
		Type:          "Серебренная текила",
		Alc:           45,
		AviableVolume: []float64{0.375, 0.750},
		Prices:        []float64{36.99, 69.99},
		Description:   `
Roca Patrón Silver изготавливается вручную из лучшей 100% голубой агавы Weber на винокурне Hacienda Patrón. 
После выпекания в течение 79 часов в небольших кирпичных печах агава прессуется двухтонным колесом тахоны из вулканического камня,
затем ферментируется и перегоняется на собственных волокнах. 
Roca Patrón Silver специально доведен до крепости 90 для создания свежей, крепкой текилы с нотками сладкой агавы. 
Текила Patrón всегда была без добавок из-за нашей непоколебимой приверженности мастерству, аутентичности и честности.
Приготовленная агава, землистый, растительный, сушеные травы, сложный`,
	},
	"Torada Silver": {
		Type:          "Серебренная текила",
		Alc:           40,
		AviableVolume: []float64{1},
		Prices:        []float64{11.99},
		Description:   `
Производится из сока растения агавы, смешанного с предварительно ферментированным соком агавы.`,
	},
	"Соль": {
		Type:          "Соль",
		Alc:           15,
		AviableVolume: []float64{1},
		Prices:        []float64{0.29},
		Description:   texts.DescSolt,
	},
	"Martini & Rossi Extra Dry Vermouth": {
		Type:          "Сухой вермут",
		Alc:           15,
		AviableVolume: []float64{0.750},
		Prices:        []float64{10.99},
		Description:   `
Вермут Martini & Rossi Extra Dry дебютировал в первый день Нового года в 1900 году после десяти лет совершенствования. 
Этот исключительный вермут стал культовым напитком Martini & Rossi, который доминировал в столетии, 
коктейлем Dry Martini & Rossi. Martini & Rossi и Tonic всегда приветствуются на вечеринках, приемах и общественных мероприятиях. 
Martini & Rossi® Extra Dry Vermouth передает суть сухого вермута через корень флорентийского ириса. 
Вытащенный из земли после трех долгих лет роста, затем высушенный на солнце и отжатый, 
корень высвобождает свои уникальные терпкие ароматические масла с ароматом фиалки, придающие узнаваемый вкус вермуту Martini & Rossi® Extra Dry. 
Обладая острым цитрусовым ароматом и оттенком малины, это классика в своем роде, которая придает игривый вкус коктейлю с мартини.`,
	},
	"Barsmith Grenadine": {
		Type:          "Гренандин",
		Alc:           0,
		AviableVolume: []float64{0.354},
		Prices:        []float64{4.99},
		Description:   `
Сладкий и сложный вкус Barsmith Grenadine идеально дополняет любимые коктейли. 
Используйте с любимым спиртным или смешайте со льдом и газированной водой для идеального напитка Ширли Темпл.`,
	},
	"Rose's Grenadine": {
		Type:          "Гренандин",
		Alc:           0,
		AviableVolume: []float64{0.357, 0.740},
		Prices:        []float64{4.99, 6.99},
		Description:   `
Rose's делает лучший сироп гренадин. В нем есть то, что большинство других гренадин давно потеряло — богатый, неповторимый вкус граната. 
Гренадин из розы, используемый во многих рецептах классических напитков, не имеет себе равных.`,
	},
	"Liqueur Fruko Schulz Triple Sec": {
		Type:          "Трипл сек fruko schulz",
		Alc:           40,
		AviableVolume: []float64{0.700},
		Prices:        []float64{21.43},
		Description:   texts.DescSimplyLemonade,
	},
	"Simply Lemonade": {
		Type:          "Лимонад",
		Alc:           0,
		AviableVolume: []float64{0.473, 1.744},
		Prices:        []float64{2.98, 4.99},
		Description:   texts.DescSimplyLemonade,
	},
	"Tropicana Orange Juice": {
		Type:          "Апельсиновый сок",
		Alc:           0,
		AviableVolume: []float64{0.946},
		Prices:        []float64{3.99},
		Description:   texts.DescTropicanaOrangeJuice,
	},
	"Dole Pineapple Juice": {
		Type:          "Ананасовый сок",
		Alc:           0,
		AviableVolume: []float64{1.360},
		Prices:        []float64{4.99},
		Description:   texts.DescDolePineappleJuice,
	},
	"Ocean Spray Cranberry Juice Cocktail": {
		Type:          "Клюквенный сок",
		Alc:           0,
		AviableVolume: []float64{0.946},
		Prices:        []float64{3.99},
		Description:   texts.DescOceanSprayCranberryJuiceCocktail,
	},
	"Realime Lime Juice": {
		Type:          "Лаймовый сок",
		Alc:           0,
		AviableVolume: []float64{0.133, 0.236},
		Prices:        []float64{1.59, 2.99},
		Description:   texts.DescRealimeLimeJuice,
	},
	"Torani French Vanilla": {
		Type:          "Ванильный сироп",
		Alc:           0,
		AviableVolume: []float64{0.750},
		Prices:        []float64{11.99},
		Description:   texts.DescToraniFrenchVanilla,
	},
	"Torani Peppermint Syrup": {
		Type:          "Мятный сироп",
		Alc:           0,
		AviableVolume: []float64{0.375, 0.750},
		Prices:        []float64{6.99, 9.99},
		Description:   texts.DescToraniPeppermintSyrup,
	},
	"Torani Amer": {
		Type:          "Креплённое вино",
		Alc:           39,
		AviableVolume: []float64{0.750},
		Prices:        []float64{11.99},
		Description:   texts.DescSugarSyrup,
	},
	"Torani Orgeat Syrup": {
		Type:          "Сахарный сироп",
		Alc:           0,
		AviableVolume: []float64{0.750},
		Prices:        []float64{11.99},
		Description:   texts.DescSugarSyrup,
	},
	"Torani Sugar Free Coconut Syrup": {
		Type:          "Кокосовый сироп",
		Alc:           0,
		AviableVolume: []float64{0.750},
		Prices:        []float64{11.99},
		Description:   texts.DescCoconutSyrup,
	},
	"Monin Coconut Syrup": {
		Type:          "Кокосовый сироп",
		Alc:           0,
		AviableVolume: []float64{0.750, 1},
		Prices:        []float64{10.57, 12.31},
		Description:   texts.DescMoninCoconutSyrup,
	},
	"Обуховская": {
		Type:          "Обуховская",
		Alc:           0,
		AviableVolume: []float64{0.500, 1, 1.500, 2},
		Prices:        []float64{0.50, 0.70, 1, 1.20},
		Description:   texts.DescObukhovskaya,
	},
	"Deep Eddy Lemon Vodka": {
		Type:          "Цитрусовая водка",
		Alc:           35,
		AviableVolume: []float64{0.750, 1.750},
		Prices:        []float64{18.69, 27.99},
		Description:   texts.DescOrangeLiqueurTrpleSec,
	},
	"Cointreau Orange Liqueur Triple Sec": {
		Type:          "Трипл сек",
		Alc:           40,
		AviableVolume: []float64{0.750, 1},
		Prices:        []float64{40.65, 45.97},
		Description:   texts.DescOrangeLiqueurTrpleSec,
	},
	"Jameson Orange Irish Whiskey": {
		Type:          "Ирландский виски",
		Alc:           30,
		AviableVolume: []float64{0.750, 1.500},
		Prices:        []float64{31.99, 54.99},
		Description:   texts.DescJamesonOrangeIrishWhiskey,
	},
	"Белый Медведь": {
		Type:          "Пиво",
		Alc:           6.2,
		AviableVolume: []float64{1, 1.500},
		Prices:        []float64{1.30, 2},
		Description:   texts.DescBeer,
	},
	"Gold": {
		Type:          "Пиво",
		Alc:           6.5,
		AviableVolume: []float64{1, 1.500},
		Prices:        []float64{1.30, 2},
		Description:   texts.DescBeer,
	},
	"Три Медведя": {
		Type:          "Пиво",
		Alc:           6.1,
		AviableVolume: []float64{1, 1.500},
		Prices:        []float64{1.30, 2},
		Description:   texts.DescBeer,
	},
	"Bud Light": {
		Type:          "Пиво",
		Alc:           4.7,
		AviableVolume: []float64{0.500},
		Prices:        []float64{1.70},
		Description:   texts.DescBeerBudLight,
	},
	"Corona Extra Lager Mexican Beer": {
		Type:          "Пиво",
		Alc:           4.6,
		AviableVolume: []float64{0.500},
		Prices:        []float64{1.80},
		Description:   texts.DescBeerCoronaExtraLagerMexican,
	},
	"Bud": {
		Type:          "Пиво",
		Alc:           5.9,
		AviableVolume: []float64{0.500},
		Prices:        []float64{1.10},
		Description:   texts.DescBeer,
	},
	"Клинское": {
		Type:          "Пиво",
		Alc:           5.7,
		AviableVolume: []float64{0.500, 1, 1.500},
		Prices:        []float64{0.90, 1.60, 2},
		Description:   texts.DescBeer,
	},
	"Sprite": {
		Type:          "Спрайт",
		Alc:           0,
		AviableVolume: []float64{0.500, 1, 1.500, 2},
		Prices:        []float64{1, 2, 3, 4},
		Description:   texts.DescSprite,
	},
	"Coca-cola": {
		Type:          "Кола",
		Alc:           0,
		AviableVolume: []float64{0.500, 1, 1.500, 2},
		Prices:        []float64{1, 2, 3, 4},
		Description:   texts.DescCocaCola,
	},
	"New Amsterdam Gin": {
		Type:          "Амстердамский сухой джин",
		Alc:           40,
		AviableVolume: []float64{0.750, 1.750},
		Prices:        []float64{14.99, 23.10},
		Description:   texts.DescNewAmsterdamLondonDryGin,
	},
	"Gordons London Dry Gin": {
		Type:          "Лондонский сухой джин",
		Alc:           40,
		AviableVolume: []float64{1, 1.750},
		Prices:        []float64{17.99, 22.99},
		Description:   texts.DescGordonsLondonDryGin,
	},
	"Bombay Dry Gin": {
		Type:          "Лондонский сухой джин",
		Alc:           43,
		AviableVolume: []float64{0.750, 1.750},
		Prices:        []float64{22.99, 34.49},
		Description:   texts.DescBombayDryGin,
	},
	"Sipsmith London Dry Gin": {
		Type:          "Лондонский сухой джин",
		Alc:           41.6,
		AviableVolume: []float64{0.375, 0.750},
		Prices:        []float64{21.99, 35.99},
		Description:   texts.DescLondonDryGin,
	},
	"Оливки": {
		Type:          "Оливки",
		Alc:           0,
		AviableVolume: []float64{1, 50, 100},
		Prices:        []float64{0.01, 0.40, 0.70},
		Description:   texts.DescOlivki,
	},
	"Листья лайма": {
		Type:          "Листья лайма",
		Alc:           0,
		AviableVolume: []float64{1, 2, 5},
		Prices:        []float64{.2, .3, .5},
		Description:   texts.DescLaim,
	},
	"BACARDI Spiced Rum": {
		Type:          "Ром",
		Alc:           35,
		AviableVolume: []float64{0.750, 1.750},
		Prices:        []float64{15.89, 24.14},
		Description:   texts.DescRum,
	},
	"BACARDI Black Rum": {
		Type:          "Тёмный ром",
		Alc:           40,
		AviableVolume: []float64{0.200, 0.750},
		Prices:        []float64{5.49, 15.49},
		Description:   texts.DescRum,
	},
	"Havana Club Anejo Clasico": {
		Type:          "Ром",
		Alc:           40,
		AviableVolume: []float64{0.050, 0.750},
		Prices:        []float64{1.49, 22.99},
		Description:   texts.DescRum,
	},
	"Havana Club Rum": {
		Type:          "Белый ром",
		Alc:           40,
		AviableVolume: []float64{0.050, 0.750},
		Prices:        []float64{1.49, 23.09},
		Description:   texts.DescRum,
	},
	"Cherry Bomb Cocktail Syrup": {
		Type:          "Сахарный сироп",
		Alc:           0,
		AviableVolume: []float64{0.400},
		Prices:        []float64{19},
		Description:   texts.DescSugarSyrup,
	},
	"Petite Canne Sugar Cane Syrup": {
		Type:          "Сахарный сироп",
		Alc:           0,
		AviableVolume: []float64{0.500, 1},
		Prices:        []float64{12.00, 24.99},
		Description:   texts.DescSugarSyrup,
	},
	"Collins Simple Syrup Cocktail Mix": {
		Type:          "Сахарный сироп",
		Alc:           0,
		AviableVolume: []float64{0.360, 0.750},
		Prices:        []float64{5.99, 6.99},
		Description:   texts.DescSugarSyrup,
	},
	"Dailys Simple Syrup": {
		Type:          "Сахарный сироп",
		Alc:           0,
		AviableVolume: []float64{0.935},
		Prices:        []float64{6.22},
		Description:   texts.DescSugarSyrup,
	},
	"Barsmith Simple Syrup": {
		Type:          "Сахарный сироп",
		Alc:           0,
		AviableVolume: []float64{0.360},
		Prices:        []float64{4.99},
		Description:   texts.DescSugarSyrup,
	},
	"Мята": {
		Type:          "Мята",
		Alc:           0,
		AviableVolume: []float64{1, 2, 5, 10},
		Prices:        []float64{0.50, 0.80, 1.50, 2},
		Description:   texts.DescMint,
	},

	"Дроблённый Лёд": {
		Type:          "Дроблённый лёд",
		Alc:           0,
		AviableVolume: []float64{1, 2, 5, 10},
		Prices:        []float64{0.50, 0.80, 1.50, 2},
		Description:   texts.DescCrushedIce,
	},
	"Лёд В Кубиках": {
		Type:          "Лёд в кубиках",
		Alc:           0,
		AviableVolume: []float64{1, 2, 5, 10},
		Prices:        []float64{0.50, 0.80, 1.50, 2},
		Description:   texts.DescIceCubes,
	},
	"Лимон": {
		Type:          "Лимон",
		Alc:           0,
		AviableVolume: []float64{1, 2, 3, 5},
		Prices:        []float64{5, 8, 12, 14},
		Description:   texts.DescLimon,
	},
	"Лайм": {
		Type:          "Лайм",
		Alc:           0,
		AviableVolume: []float64{1, 2, 3, 5},
		Prices:        []float64{6, 10, 13, 16.50},
		Description:   texts.DescLaim,
	},
	"Содовая": {
		Type:          "Содовая",
		Alc:           0,
		AviableVolume: []float64{0.250, 0.500, 1, 1.500},
		Prices:        []float64{0.30, 0.50, 0.70, 1},
		Description:   texts.DescSoda,
	},
	"Эбботтс Биттер": {
		Type:          "Ароматический биттер",
		Alc:           41.5,
		AviableVolume: []float64{0.100},
		Prices:        []float64{20.50},
		Description:   texts.DescAbbotsBitter,
	},
	"Апероль": {
		Type:          "Апероль",
		Alc:           22,
		AviableVolume: []float64{0.750, 1},
		Prices:        []float64{27.99, 34.71},
		Description:   texts.DescAperrol,
	},
	"Просекко": {
		Type:          "Просекко",
		Alc:           11,
		AviableVolume: []float64{0.750},
		Prices:        []float64{14.94},
		Description:   texts.DescProsecco,
	},
	"Cruzan Aged Light Rum": {
		Type:          "Белый ром",
		Alc:           40,
		AviableVolume: []float64{0.375, 0.750},
		Prices:        []float64{7.99, 13.19},
		Description:   texts.DescWhiteRum,
	},
	"Rum Mocambo 20 Year Old": {
		Type:          "Ром",
		Alc:           40,
		AviableVolume: []float64{0.750},
		Prices:        []float64{42.99},
		Description:   texts.DescWhiteRum,
	},
	"Titos Handmade Vodka": {
		Type:          "Водка",
		Alc:           40,
		AviableVolume: []float64{1, 1.750},
		Prices:        []float64{28.40, 36.29},
		Description:   texts.DescVodka,
	},
	"Belsazar Vermouth": {
		Type:          "Красный вермут",
		Alc:           40,
		AviableVolume: []float64{0.375, 0.750},
		Prices:        []float64{7.99, 13.19},
		Description:   texts.DescRedVermouth,
	},
}
