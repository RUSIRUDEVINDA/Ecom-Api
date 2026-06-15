package env

import (
	"os"

	"github.com/joho/godotenv"
)

// Load loads variables from a .env file if present.
func Load() {
	_ = godotenv.Load()
}

// GetString returns the environment variable value for key or fallback if not set.
func GetString(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
