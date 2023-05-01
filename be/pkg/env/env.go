package env

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/joho/godotenv"
)

var ErrEnvNotFound = errors.New("No env for the specified key")
var ErrEnvWrongType = errors.New("The specified env is the wrong type")

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

func GetInt(key string) (int, error) {
	loadEnv()

	env := os.Getenv(key)
	if env == "" {
		return 0, ErrEnvNotFound
	}
	value, err := strconv.Atoi(env)
	if err != nil {
		return 0, ErrEnvWrongType
	}
	return value, nil
}

func GetIntWithDefault(key string, d int) int {
	loadEnv()

	env := os.Getenv(key)
	if env == "" {
		return d
	}
	value, err := strconv.Atoi(env)
	if err != nil {
		return d
	}
	return value
}

func GetString(key string) (string, error) {
	loadEnv()
	env := os.Getenv(key)
	if env == "" {
		return "", ErrEnvNotFound
	}

	return env, nil
}

func GetStringWithDefault(key string, d string) string {
	loadEnv()
	env := os.Getenv(key)
	if env == "" {
		return d
	}

	return env
}
