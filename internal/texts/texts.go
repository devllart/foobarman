package texts

// Typics

const Unknown = "Неизвестен"

// Scenes
// ** Hello
const WhatIsName = "%CНу и как тебя зовут юный бармен: %B"
const NameIsOk = "%CОтлично %s, ну что пошлите в магазин закупаться? (Нажми Enter): "

// ** Bar

const SceneBar = "%s%R[%YБАР %B\"Монатль и Молись\"%R]%C\n\n"

// ** Store
const StoreHello = "Добро пожаловать в магазин %RBrizly%C названый в честь Бри Ларсон.\n%s выбирай, что хочешь.\n\n"

const StoreDrinkInfo = "%d. %B%s%C (%R%.2f%C%%) | %Y%s%C (вкус %G%s%C)\n"
const StoreDrinkInfoPrice = " | %Y%.3f$%C за %G%.3f %s%C"

// Commands

const AllowCommands = "Доступные команды (регистр не важен):\n"

// Hint

const BarmanStatus = "\nБармен %s   денег: %Y%.2f $%C\n\n"
const ShowDrinkInBar = "%B%s%C %R%.2f%%%C (%G%.3f %s%C) %Y%dX%C | %s %G%.3f %s%C\n"
const SelectIngredients = "%YИз каких ингридиентов будет ваш коктель?%C\n"
const ClueStore = "\n%YПодсказка:\n%CНапиши %R{%BНазвание%C или %Bномер%C напитка%R} {%GЕго объём %R(%Cпо умолчанию наименьший — можно выбрать наименьший написав %B0%R)} {%YКоличество%C (по умолчанию %B1)%R}%C, чтобы купить.\nНапример \"%Y\\>%B Просекко 0.75 3%C\" или  \"%Y\\>%B 1 0 3%C\", также можно написать %Brand%C чтобы закупиться рандомно \n"

// Buying error

const NotAvailableDrink = "%R!%C Напитка %B%s%C нет продаже в продаже\n"
const NotAvailableIndexDrink = "%R!%C Напитка под номером %B%d%C нет в продаже\n"
const NotEnoughFundsToBuy = "%R!%C Недостаточно средств для покупки (общая сумма составила %Y%.2f $%C)\n"
const NotVolumeOfDrink = "%R!%C %B%s%C с объёмом %G%.3f .л%C нет в продаже, возьмите другой объём\n"

const IncorrectAmountOfDrink = "%R!%C Неверно указанно количество напитка\n"
const IncorrectVolumeOfDrink = "%R!%C Неверно указан объём напитка\n"

// Coctails

const CoctailIsReady = "У вас получился %B%s%C — очень хорошо\n"
const DontTheRecipies = "Чтож жаль, но такого %Bрецепта нет%C\n"
