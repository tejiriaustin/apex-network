package env

import (
	"errors"
	"os"
)

const (
	DbUsername = "db_user_name"
	DbHost     = "db_host"
	DbPassword = "db_password"
	DbDatabase = "db_database"
	DbTimeZone = "db_time_zone"
)

type Env map[string]interface{}

func NewEnv() Env {
	return Env{}
}

func (e Env) SetEnv(key string, value interface{}) Env {
	e[key] = value

	return e
}

func (e Env) GetEnvString(serviceName string) (string, error) {
	value := e[serviceName]
	if value == nil {
		return "", errors.New("failed to get service env")
	}

	valueString, ok := value.(string)
	if !ok {
		return "", errors.New("failed to convert env to string")
	}
	return valueString, nil
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		return ""
	}
	return value
}

func MustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("failed to get service env: " + key)
	}
	return value
}
