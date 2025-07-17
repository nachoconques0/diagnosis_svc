package env

import (
	"fmt"
	"os"
)

const (
	Test = "test"
)

// LoadOrDefault retrieves the value of the environment variable. Otherwise it return the provided default value
func LoadOrDefault(env, def string) string {
	val, ok := os.LookupEnv(env)
	if !ok {
		return def
	}

	return val
}

// LoadOrPanic retrieves the value of the environment variable. Otherwise it panics
func LoadOrPanic(env string) string {
	val, ok := os.LookupEnv(env)
	if !ok {
		panic(fmt.Sprintf("Missing '%s' environment variable\n", env))
	}

	return val
}

func IsTest(env string) bool {
	return env == Test
}
