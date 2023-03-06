package configs

import (
	"os"
)

const (
	productionEnv = "production"
	environmentKey = "ENVIRONMENT"
)

var (
	environment string
)

// IsProduction checks the environment name and indicates if it is production or
// not.
func IsProduction() bool {
	return environment == productionEnv
}

func loadEnvsConfig() {
	environment = os.Getenv(environmentKey)
}

func init() {
	loadEnvsConfig()
}