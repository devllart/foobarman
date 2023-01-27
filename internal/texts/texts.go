package texts

// Scenes
// ** Hello
var WhatIsName = "%CНу и как тебя зовут юный бармен: %B"
var NameIsOk = "%CОтлично %s, ну что пошлите в магазин закупаться? (Нажми Enter): "

// ** Bar

var SceneBar = "%s%R[%YБАР %B\"Монатль и Молись\"%R]%C\n\n"

// ** Store
var StoreHello = "Добро пожаловать в магазин %RBrizly%C названый в честь Бри Ларсон.\n%s выбирай, что хочешь.\n\n"

var StoreDrinkInfo = "%d. %B%s%C (%R%.2f%C%%) | %Y%s%C (вкус %G%s%C)\n"
var StoreDrinkInfoPrice = "| %Y%.3f$%C за %G%.3f .л%C"

// Commands

var AllowCommands = "Доступные команды (регистр не важен):\n"

// Hint

var BarmanStatus = "\nБармен %s   денег: %Y%.2f $%C\n\n"
var ShowDrinkInBar = "%B%s%C %R%.2f%%%C (%G%g .л%C) %Y%dX%C | в последней бутылке осталось %G%g .л%C\n"
var SelectIngredients = "%YИз каких ингридиентов будет ваш коктель?%C\n"
var ClueStore = "\nПодсказка: Напиши {Название напитка} {Его объём} {Количество (по умолчанию 1)}, чтобы купить.\nНапример \"%Y\\>%C Просекко 0.75 3\"\n"

// Buying error

var NotAvailableDrink = "%R!%C Напитка %B%s%C нет продаже в продаже\n"
var NotAvailableIndexDrink = "%R!%C Напитка под номером %B%d%C нет в продаже\n"
var NotEnoughFundsToBuy = "%R!%C Недостаточно средст для покупки (общая сумма составила %Y%.2f $%C)\n"
var NotVolumeOfDrink = "%R!%C $B%s%C с объёмом %G%.3f .л%C нет в продаже, возьмите другой объём\n"

var IncorrectAmountOfDrink = "Неверно указанно количество напитка\n"
var IncorrectVolumeOfDrink = "Неверно указан объём напитка\n"

// Coctails

var CoctailIsReady = "У вас получился %B%s%C — очень хорошо\n"
var DontTheRecipies = "Чтож жаль, но такого рецепта нет\n"
