package config

import (
	"crypto/sha1"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/pavelkg/tradem-mon-api/internal/database"
)

type Config struct {
	IsDebug       bool `json:"isDebug,omitempty"`
	IsDevelopment bool `json:"isDevelopment,omitempty"`
	Listen        struct {
		Host string
		Port string
	}
	DbParams   database.DbParams
	JwtSecret  string
	HostPrefix string
}

func GetConfig(fileNames ...string) *Config {

	instance := Config{}

	// load .env file
	err := godotenv.Load(fileNames...)
	if err != nil {
		fmt.Print("Error loading .env file\n")
	}
	instance.IsDebug = getEnvVar("IS_DEBUG", "") == "true"
	instance.IsDevelopment = getEnvVar("APP_STATUS", "dev") == "dev"

	instance.Listen.Host = getEnvVar("SERVER_HOST", "0.0.0.0")
	instance.Listen.Port = getEnvVar("SERVER_PORT", "8765")

	instance.DbParams.DbName = getEnvVar("DB_NAME", "")
	instance.DbParams.Host = getEnvVar("DB_HOST", "localhost")
	instance.DbParams.Port = getEnvVar("DB_PORT", "5432")
	instance.DbParams.User = getEnvVar("DB_USER", "postgres")
	instance.DbParams.Password = getEnvVar("DB_PASS", "")

	data := []byte(fmt.Sprintf("%s", time.Now()))
	instance.JwtSecret = getEnvVar("JWT_SECRET", fmt.Sprintf("%x", sha1.Sum(data)))
	instance.HostPrefix = getEnvVar("HOST_PREFIX", "")

	return &instance
}

// GetEnvVar func to get env value
func getEnvVar(key string, def string) string {
	val := os.Getenv(key)
	if val == "" {
		val = def
	}
	return val
}
