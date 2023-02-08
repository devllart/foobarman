package texts

// Typics

const Unknown = "Неизвестен"
const UnknownCommand = " %R!%C Неизвестная комманда: %B%s%C"

// Scenes

// ** Store
const StoreHello = "Добро пожаловать в магазин %RBrizly%C названый в честь Бри Ларсон.\n%s выбирай, что хочешь.\n\n"

const StoreProductInfo = "%d. %B%s%C %R%.2f %%%C | %Y%s%C (вкус %G%s%C)\n"
const StoreProductInfoPrice = " | %Y%.2f$%C за %G%.3f%s%C"

// ** Hello
const WhatIsName = "%CНу и как тебя зовут юный бармен: %B"
const NameIsOk = "%CОтлично %s, ну что пошли в магазин закупаться? (Нажми Enter): "

// ** Bar

const SceneBar = "%s%R[%YБАР %B\"Монталь и Молись\"%R]%C\n\n"

// ** Recipes

const SceneRecipes = "%s%YДоступные рецепты%C\n\n"

// Commands

const AllowCommands = "Доступные команды (регистр не важен):\n"

// Hint

const BarmanStatus = "\nБармен %s   денег: %Y%.2f $%C         Количество довольных посетителей: %Y%d%C\n\n"
const ShowProductInBar = "%B%s%C %R%.2f %%%C  (%G%.3f%s%C) %Y%dX%C | %s %G%.3f%s%C\n"
const SelectIngredients = "%YИз каких ингредиентов будет твой коктель?%C\n"

// Products

const TotalLeftVolume = "всего осталось"
const LeftVolumeInLastBottle = "в последней бутылке осталось"

// ** Coctails

const CoctailIsReady = "У вас получился %B%s%C — очень хорошо"
const DontTheRecipies = "Чтож жаль, но такого %Bрецепта нет%C"

const NotSayAboutCoctail = "%sО коктейле \"%B%s%C\" нечего сказать\n\n"

// ** Recipes

const UnknownRecipes = "\n%s%R):%C Тебе пока неизветен рецепт для коктейля \"%R%s%C\"\n"
const RecipesCoctail = "%sРецепт коктейля %R%s%C\n\n"
