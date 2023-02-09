package alert

/**
 * Don't panic
 */

func UnknownCommand(command string) {
	Text(" %R!%C Неизвестная комманда: %B%s%C", command)
	Line()
}

func CoctailIsReady(name string) {
	Line()
	Text("У вас получился %B%s%C — очень хорошо", name)
	Line()
}

func DontTheRecipies() {
	Clue("Чтож жаль, но такого %Bрецепта нет%C")
	Line()
}

func CheatCodeActivate() {
	Clue("Чит-код активирован")
	Line()
}

func RestartGame() {
	TextBlue(" ! Игра перезапущена")
	Line()
}

func VolumeOfProductSpent(volume float64, typeVolume, name string) {
	TextRed(" -%.3f%s %B%s%C", volume, typeVolume, name)
	Line()
}

func MissClient() {
	Clue("Ты упустил клиента")
	Line()
}

func MissHistory() {
	Line()
	Clue("И ещё похоже ты упустил интересную историю")
	Line()
}
