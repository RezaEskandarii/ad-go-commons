package env_manager

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
	"time"
)

// loadEnv initializes environment variables from a .env file
func loadEnv() {
	_ = godotenv.Load()
}

// GetString returns the environment variable as a string
func GetString(key string) string {
	loadEnv()
	return strings.TrimSpace(os.Getenv(key))
}

// GetInt returns the environment variable as an int
func GetInt(key string) (int, error) {
	loadEnv()
	value := strings.TrimSpace(os.Getenv(key))
	return strconv.Atoi(value)
}

// GetInt64 returns the environment variable as an int64
func GetInt64(key string) (int64, error) {
	loadEnv()
	value := strings.TrimSpace(os.Getenv(key))
	return strconv.ParseInt(value, 10, 64)
}

// GetUint returns the environment variable as a uint
func GetUint(key string) (uint, error) {
	loadEnv()
	value := strings.TrimSpace(os.Getenv(key))
	v, err := strconv.ParseUint(value, 10, 64)
	return uint(v), err
}

// GetUint64 returns the environment variable as a uint64
func GetUint64(key string) (uint64, error) {
	loadEnv()
	value := strings.TrimSpace(os.Getenv(key))
	return strconv.ParseUint(value, 10, 64)
}

// GetBool returns the environment variable as a boolean
func GetBool(key string) (bool, error) {
	loadEnv()
	value := strings.TrimSpace(os.Getenv(key))
	return strconv.ParseBool(value)
}

// GetFloat32 returns the environment variable as a float32
func GetFloat32(key string) (float32, error) {
	loadEnv()
	value := strings.TrimSpace(os.Getenv(key))
	v, err := strconv.ParseFloat(value, 32)
	return float32(v), err
}

// GetFloat64 returns the environment variable as a float64
func GetFloat64(key string) (float64, error) {
	loadEnv()
	value := strings.TrimSpace(os.Getenv(key))
	return strconv.ParseFloat(value, 64)
}

// GetDuration returns the environment variable as a time.Duration
func GetDuration(key string) (time.Duration, error) {
	loadEnv()
	value := strings.TrimSpace(os.Getenv(key))
	return time.ParseDuration(value)
}
