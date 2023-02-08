package config

import "os"

var Env = ""
var Stage = 1

func init() {
	var exist bool
	if Env, exist = os.LookupEnv("ENVIRONMENT"); exist == false {
		Env = "production"
	}
}
