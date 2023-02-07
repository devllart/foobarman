package config

import "os"

var Env = ""

func init() {
	var exist bool
	if Env, exist = os.LookupEnv("ENVIRONMENT"); exist == false {
		Env = "production"
	}
}
