package texts

// Typics

const Unknown = "Неизвестен"
const UnknownCommand = "%R!%C Неизвестная комманда: %s\n"

// Scenes

// ** Store
const StoreHello = "Добро пожаловать в магазин %RBrizly%C названый в честь Бри Ларсон.\n%s выбирай, что хочешь.\n\n"

const StoreProductInfo = "%d. %B%s%C %R%.2f %%%C | %Y%s%C (вкус %G%s%C)\n"
const StoreProductInfoPrice = " | %Y%.2f$%C за %G%.3f%s%C"

// ** Hello
const WhatIsName = "%CНу и как тебя зовут юный бармен: %B"
const NameIsOk = "%CОтлично %s, ну что пошлите в магазин закупаться? (Нажми Enter): "

// ** Bar

const SceneBar = "%s%R[%YБАР %B\"Монталь и Молись\"%R]%C\n\n"

// ** Recipes

const SceneRecipes = "%s%YДоступные рецепты%C\n\n"

// Commands

const AllowCommands = "Доступные команды (регистр не важен):\n"

// Hint

const BarmanStatus = "\nБармен %s   денег: %Y%.2f $%C         Количство довольных посетителей: %Y%d%C\n\n"
const ShowProductInBar = "%B%s%C %R%.2f %%%C  (%G%.3f%s%C) %Y%dX%C | %s %G%.3f%s%C\n"
const SelectIngredients = "%YИз каких ингредиентов будет твой коктель?%C\n"

// Products

const TotalLeftVolume = "всего осталось"
const LeftVolumeInLastBottle = "в последней бутылке осталось"

// ** Products buying error

const NotAvailableProduct = "%R!%C Напитка %B%s%C нет в продаже\n"
const NotAvailableIndexProduct = "%R!%C Напитка под номером %B%d%C нет в продаже\n"
const NotEnoughFundsToBuy = "%R!%C Недостаточно средств для покупки %B%s%C (общая сумма составила %Y%.2f $%C)\n"
const NotVolumeOfProduct = "%R!%C %B%s%C с объёмом %G%.3f .л%C нет в продаже, возьмите другой объём\n"

const IncorrectAmountOfProduct = "%R!%C Неверно указанно количество напитка\n"
const IncorrectVolumeOfProduct = "%R!%C Неверно указан объём напитка\n"

// ** Products buy

const ProductBought = "%Y+%C %B%s%C %G%.3f%s%C %Y%dX%C куплено (%Y-%.2f $%C)\n"
const ProductBoughtYet = "%Y+%C %B%s%C %G%.3f%s%C %Y%dX%C куплено ещё (общее количество: %G%d%C) (%Y-%.2f $%C)\n"

// ** Coctails

const CoctailIsReady = "У вас получился %B%s%C — очень хорошо\n"
const DontTheRecipies = "Чтож жаль, но такого %Bрецепта нет%C\n"

const NotSayAboutCoctail = "%sО коктейле \"%B%s%C\" нечего сказать\n\n"

// ** Recipes

const UnknownRecipes = "\n%s%R):%C Тебе пока неизветен рецепт для коктейля \"%R%s%C\"\n"
const RecipesCoctail = "%sРецепт коктейля %R%s%C\n\n"
