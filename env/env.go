package env

import (
	"github.com/joho/godotenv"
	"os"
)

const (
	DbUsername = "DB_USERNAME"
	DbHost     = "DB_HOST"
	DbPassword = "DB_PASSWORD"
	DbDatabase = "DB_DATABASE"
	DbTimeZone = "DB_TIMEZONE"
	DbUrl      = "DB_URL"
)

type Env map[string]interface{}

func NewEnv() Env {
	return Env{}
}
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("failed to load ")
	}
}

func (e Env) SetEnv(key string, value interface{}) Env {
	e[key] = value

	return e
}

func (e Env) GetEnvString(serviceName string) string {
	value := e[serviceName]
	if value == nil {
		return ""
	}

	valueString, ok := value.(string)
	if !ok {
		return ""
	}
	return valueString
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		return ""
	}
	return value
}

func MustGetEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic("failed to get service env: " + key)
	}
	return value
}
