package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDriver       string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	JWTSecret      string
	JWTExpiration  time.Duration
	JWTIssuer      string
	OSMAPIEndpoint string
	OSMAPIKey      string
	LocationName   string
}

var AppConfig Config

// LoadConfig memuat konfigurasi dari file .env
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	AppConfig = Config{
		DBDriver:       getEnv("DB_DRIVER", ""),
		DBHost:         getEnv("DB_HOST", ""),
		DBPort:         getEnv("DB_PORT", ""),
		DBUser:         getEnv("DB_USER", ""),
		DBPassword:     getEnv("DB_PASSWORD", ""),
		DBName:         getEnv("DB_NAME", ""),
		JWTSecret:      getEnv("JWT_SECRET", ""),
		JWTExpiration:  getDurationEnv("JWT_EXPIRATION", 1),
		JWTIssuer:      getEnv("JWT_ISSUER", ""),
		OSMAPIEndpoint: getEnv("OSM_API_ENDPOINT", ""),
		OSMAPIKey:      getEnv("OSM_API_KEY", ""),
		LocationName:   getEnv("LOCATION_NAME", ""),
	}
}

// getEnv mengambil variabel lingkungan dengan nilai default
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getDurationEnv mengambil variabel lingkungan dengan nilai default waktu
func getDurationEnv(key string, defaultValue int) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		duration, err := time.ParseDuration(value + "h")
		if err != nil {
			return time.Duration(defaultValue) * time.Hour
		}
		return duration
	}
	return time.Duration(defaultValue) * time.Hour
}
