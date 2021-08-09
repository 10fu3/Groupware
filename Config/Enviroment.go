package Config

import "github.com/kelseyhightower/envconfig"
import "../Entity"

var Env = getEnv()

func getEnv() Entity.Env {
	var env Entity.Env
	envconfig.Process("", &env)
	return env
}
