package state

var Status = "Norm" // Status barman

func HandlerStatus() {
	if Money < 10 {
		SetStatusBad()
	} else {
		SetStatusNorm()
	}
}

func SetStatusBad() {
	Status = "Bad"
	// structs.NewTastes.SetValue("Лондонский сухой джин", "Очень сухой :/ — но твоя девушка суше ;) ")
}

func SetStatusNorm() {
	Status = "Norm"
	// structs.NewTastes.SetValue("Лондонский сухой джин", "Очень сухой :/")
}
