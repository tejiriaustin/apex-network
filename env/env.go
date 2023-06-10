package env

import (
	"github.com/joho/godotenv"
	"os"
)

const (
	Port  = "PORT"
	DbUrl = "DB_URL"
)

type Env map[string]interface{}

func NewEnv() Env {
	return Env{}
}
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("failed to load")
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
