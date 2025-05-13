package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	SshDefaultUser     string
	SshDefaultPassword string
}

func InitEnviromentVariables() (*Configuration, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	globalConfig := &Configuration{
		SshDefaultUser:     getEnvVarOrDefault("SSH_DEFAULT_SSH_USER", ""),
		SshDefaultPassword: getEnvVarOrDefault("SSH_DEFAULT_SSH_PASSWORD", ""),
	}
	return globalConfig, nil
}

func getEnvVarOrDefault(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}
