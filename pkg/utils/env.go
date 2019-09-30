package env

import (
	"fmt"
	"os"
)

var (
	envPrefix = "GOBJECTION_"
)

// SetUpEnvPrefix set new environment var common prefix
func SetUpEnvPrefix(prefix string) {
	envPrefix = prefix
}

// GetEnv return environment var with common prefix
func GetEnv(key string) string {
	k := fmt.Sprintf("%s%s", envPrefix, key)
	return os.Getenv(k)
}
