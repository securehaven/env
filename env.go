package env

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrNotFound = errors.New("environment variable not found")
)

func Get[T Allowed](name string, fallback T) T {
	raw, exists := os.LookupEnv(name)

	if !exists {
		return fallback
	}

	parsed, err := Parse[T](raw)

	if err != nil {
		return fallback
	}

	return parsed
}

func GetStrict[T Allowed](name string, fallback T) (T, error) {
	raw, exists := os.LookupEnv(name)

	if !exists {
		return fallback, ErrNotFound
	}

	parsed, err := Parse[T](raw)

	if err != nil {
		return fallback, err
	}

	return parsed, nil
}

func MustGet[T Allowed](name string) T {
	raw, exists := os.LookupEnv(name)

	if !exists {
		panic(fmt.Sprintf("missing environment variable %q", name))
	}

	return MustParse[T](raw)
}
