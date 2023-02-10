package state

// ** Temp State
var TempBool = false // Nothing say
var TempStr = ""     // Nothing say

func ClearTemp() {
	Command = ""
	Args = []string{}
	TempBool = false

	ProductsIds = []string{}
}
