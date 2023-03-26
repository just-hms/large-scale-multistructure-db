package env

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/joho/godotenv"
)

var EnvNotFound = errors.New("No env for the specified key")
var EnvWrongType = errors.New("The specified env is the wrong type")

var loaded = false

func loadEnv() {
	if loaded {
		return
	}
	loaded = true

	_, caller, _, ok := runtime.Caller(0)
	if !ok {
		panic("error retrevieng the .env file")
	}

	envPath := filepath.Join(filepath.Dir(caller), "../..", "../.env")

	err := godotenv.Load(envPath)
	if err != nil {
		panic("error retrevieng the .env file")
	}
}

func GetInteger(key string) (int, error) {
	loadEnv()

	env := os.Getenv(key)
	if env == "" {
		return 0, EnvNotFound
	}
	value, err := strconv.Atoi(env)
	if err != nil {
		return 0, EnvWrongType
	}
	return value, nil
}

func GetString(key string) (string, error) {
	loadEnv()
	env := os.Getenv(key)
	if env == "" {
		return "", EnvNotFound
	}

	return env, nil
}
