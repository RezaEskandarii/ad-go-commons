package env_manager

import (
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func Get(key string) string {
	err := godotenv.Load()
	if err != nil {
		return strings.TrimSpace(os.Getenv(key))
	}
	return strings.TrimSpace(os.Getenv(key))
}
