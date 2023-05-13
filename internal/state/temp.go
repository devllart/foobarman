package state

// ** Temp State
var (
	TempBool = false // Nothing say
	TempStr = ""     // Nothing say
)

func ClearTemp() {
	Command = ""
	Args = []string{}
	TempBool = false

	ProductsIds = []string{}
}
