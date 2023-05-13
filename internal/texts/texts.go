package texts

const (
	// Typics

	Unknown = "Неизвестен"
	UnknownCommand = " %R!%C Неизвестная комманда: %B%s%C"
   
   // Scenes
   
   // ** Store
	StoreHello = "Добро пожаловать в магазин %RBrizly%C названый в честь Бри Ларсон.\n%s выбирай, что хочешь.\n\n"
   
	StoreProductInfo = "%d. %B%s%C %R%.2f %%%C | %Y%s%C (вкус %G%s%C)\n"
	StoreProductInfoPrice = " | %Y%.2f$%C за %G%.3f%s%C"
   
   // ** Hello
	WhatIsName = "%CНу и как тебя зовут юный бармен: %B"
	NameIsOk = "%CОтлично %s, ну что пошли в магазин закупаться? (Нажми Enter): "
   
   // ** Bar
   
	SceneBar = "%s%R[%YБАР %B\"Монталь и Молись\"%R]%C\n\n"
   
   // ** Recipes
   
	SceneRecipes = "%s%YДоступные рецепты%C\n\n"
   
   // Commands
   
	AllowCommands = "Доступные команды (регистр не важен):\n"
   
   // Hint
   
	BarmanStatus = "\nБармен %s   денег: %Y%.2f $%C         Количество довольных посетителей: %Y%d%C\n\n"
	ShowProductInBar = "%B%s%C %R%.2f %%%C  (%G%.3f%s%C) %Y%dX%C | %s %G%.3f%s%C\n"
	SelectIngredients = "%YИз каких ингредиентов будет твой коктель?%C\n"
   
   // Products
   
	TotalLeftVolume = "всего осталось"
	LeftVolumeInLastBottle = "в последней бутылке осталось"
   
   // ** Coctails
   
	CoctailIsReady = "У тебя получился %B%s%C — очень хорошо"
	DontTheRecipies = "Чтож жаль, но такого %Bрецепта нет%C"
   
	NotSayAboutCoctail = "%sО коктейле \"%B%s%C\" нечего сказать\n\n"
   
   // ** Recipes
   
	UnknownRecipes = "\n%s%R):%C Тебе пока неизветен рецепт для коктейля \"%R%s%C\"\n"
	RecipesCoctail = "%sРецепт коктейля %R%s%C\n\n"
)
