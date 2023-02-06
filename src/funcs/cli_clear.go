package funcs

import (
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	clear["android"] = clear["linux"]
	clear["darwin"] = clear["linux"]
}

func CliClear() {
	value, ok := clear[runtime.GOOS] // Runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          // If we defined a clear func for that platform:
		value() // We execute it
	} else { // Unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
