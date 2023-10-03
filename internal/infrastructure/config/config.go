package config

import (
	"os"
	"strconv"
)

func getEnvWithDefaults(key string, defaults string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaults
	}

	return value
}

func getIntEnvValueWithDefaults(key string, defaults int) int {
	value := os.Getenv(key)

	if value == "" {
		return defaults
	}

	val, err := strconv.Atoi(value)

	if err != nil {
		return defaults
	}

	return val
}

func getBoolEnvValueWithDefaults(key string, defaults bool) bool {
	value := os.Getenv(key)

	if value == "" {
		return defaults
	}

	val, err := strconv.ParseBool(value)

	if err != nil {
		return defaults
	}

	return val
}
