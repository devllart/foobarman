package texts

// Scenes
var StoreHello string = "Добро пожаловать в магазин %RBrizly%C названый в честь Бри Ларсон.\n%s выбирай, что хочешь.\n\n"

// Commands
var AllowCommands string = "Доступные команды (регистр не важен):\n"

// Hint
var BarmanStatus string = "\nБармен %s   денег: %Y%.2f $%C\n\n"
var ShowDrinkInBar string = "%B%s%C %R%.2f процентов%C (%G%g .л%C) %Y%dX%C | в последней бутылке осталось %G%g .л%C\n"
var SelectIngredients string = "%YИз каких ингридиентов будет ваш коктель?%C\n"
var ClueStore string = "\nПодсказка: Напиши {Название напитка} {Его объём} {Количество (по умолчанию 1)}, чтобы купить.\nНапример \"%Y\\>%C Просекко 0.75 3\"\n"

// Buying error

var NotAvailableDrink string = "%R!%C Напитка %B%s%C нет продаже в продаже\n"
var NotAvailableIndexDrink string = "%R!%C Напитка под номером %B%d%C нет в продаже\n"
var NotEnoughFundsToBuy string = "%R!%C Недостаточно средст для покупки (общая сумма составила %Y%.2f $%C)\n"
var NotVolumeOfDrink string = "%R!%C $B%s%C с объёмом %G%.3f .л%C нет в продаже, возьмите другой объём\n"

var IncorrectAmountOfDrink string = "Неверно указанно количество напитка\n"
var IncorrectVolumeOfDrink string = "Неверно указан объём напитка\n"
