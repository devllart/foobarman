package alert

/**
 * Alert for products don't panic
 */

// Bought product
func IncorrectAmountOfProduct() {
	Text(" %R!%C Неверно указанно количество напитка")
}

func IncorrectVolumeOfProduct() {
	Text(" %R!%C Неверно указан объём напитка")
	Line()
}

func ProductBought(productName, typeVolume string, volume float64, sumPrice float64, count int) {
	Text(" %Y+%C %B%s%C %G%.3f%s%C %Y%dX%C куплено (%Y-%.2f $%C)", productName, volume, typeVolume, count, sumPrice)
	Line()
}

func ProductBoughtYet(productName, typeVolume string, volume, sumPrice float64, count, countSum int) {
	Text(" %Y+%C %B%s%C %G%.3f%s%C %Y%dX%C куплено ещё (общее количество: %G%d%C) (%Y-%.2f $%C)", productName, volume, typeVolume, count, countSum, sumPrice)
	Line()
}
