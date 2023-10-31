package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	GinMode    = getGinMode()
	DBUsername = getEnvDefault("DBUSERNAME", "ruth")
	DBPassword = getEnvDefault("DBPASSWORD", "1234")
	DBPort     = getEnvDefault("DBPORT", "3306")
	DBHost     = getEnvDefault("DBHOST", "0.0.0.0")
	DBName     = getEnvDefault("DBNAME", "exampledb")
)

func getEnvDefault(env string, defaultEnv string) string {
	envContent := os.Getenv(env)
	if envContent == "" {
		envContent = defaultEnv
	}

	return envContent
}
func getEnvDefaultBool(env string, defaultEnv bool) bool {
	envContent, err := strconv.ParseBool(os.Getenv(env))
	if err != nil {
		envContent = defaultEnv
	}
	return envContent
}

func getGinMode() string {
	godotenv.Load(".env")

	gin_mode := os.Getenv("GIN_MODE")
	if gin_mode != "debug" && gin_mode != "release" {
		gin_mode = "debug"
	}

	return gin_mode
}
