package env

import (
	"fmt"
	"os"
)

var (
	envPrefix = "GOBJECTION_"
)

// InitWithPrefix set new environment var common prefix
func InitWithPrefix(prefix string) {
	envPrefix = prefix
}

// GetEnv return environment var with common prefix
func GetEnv(key string) string {
	k := fmt.Sprintf("%s%s", envPrefix, key)
	return os.Getenv(k)
}

// SetEnv update environment var with common prefix
func SetEnv(key string, val string) error {
	k := fmt.Sprintf("%s%s", envPrefix, key)
	return os.Setenv(k, val)
}
