package config

import "os"

var (
	Env = ""
	Stage = 1
)

func init() {
	var exist bool
	if Env, exist = os.LookupEnv("ENVIRONMENT"); !exist {
		Env = "production"
	}
}
