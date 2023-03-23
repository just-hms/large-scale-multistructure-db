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

var ()

func load() {
	loaded = true

	_, b, _, ok := runtime.Caller(0)
	if !ok {
		panic("error retrevieng the .env file")
	}

	envPath := filepath.Join(filepath.Dir(b), "../..", "../.env")

	err := godotenv.Load(envPath)
	if err != nil {
		panic("error retrevieng the .env file")
	}
}

func GetInteger(key string) (int, error) {
	if !loaded {
		load()
	}
	env := os.Getenv(key)
	if env == "" {
		return 0, EnvNotFound
	}
	return strconv.Atoi(env)
}

func GetString(key string) (string, error) {
	if !loaded {
		load()
	}
	env := os.Getenv(key)
	if env == "" {
		return "", EnvNotFound
	}

	return env, nil
}
