package config

func ShowAll() {
	ShowBarman = true
	ShowCommands = true
	ShowHits = true
	ShowDescription = true
	ShowInstruction = true
}

func HideAll() {
	ShowBarman = false
	ShowCommands = false
	ShowHits = false
	ShowDescription = false
	ShowInstruction = false
}

func TurnShowBarman() {
	ShowBarman = !ShowBarman
}
func TurnShowCommands() {
	ShowCommands = !ShowCommands
}
func TurnShowHits() {
	ShowHits = !ShowHits
}
func TurnShowDescription() {
	ShowDescription = !ShowDescription
}
func TurnShowInstruction() {
	ShowInstruction = !ShowInstruction
}
