package config

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// config database
type DbConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBDriver   string
}

// config jwt
type JwtConfig struct {
	JWTSecret     string
	JWTExpiration time.Duration
	JWTIssuer     string
}

type OsmApiConfig struct {
	OSMAPIEndpoint string
	OSMAPIKey      string
	LocationName   string
}

type Config struct {
	DbConfig
	JwtConfig
	OsmApiConfig
}

// type Config struct {
// 	DBDriver       string
// 	DBHost         string
// 	DBPort         string
// 	DBUser         string
// 	DBPassword     string
// 	DBName         string
// 	JWTSecret      string
// 	JWTExpiration  time.Duration
// 	JWTIssuer      string
// 	OSMAPIEndpoint string
// 	OSMAPIKey      string
// 	LocationName   string
// }

// NewConfig adalah constructor untuk inisialisasi Config
// func NewConfig() *Config {
// 	// Muat file .env jika ada
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatalf("Error loading .env file: %v", err)
// 	}

// add jwt configuration - by udin
// 	longTime, _ := strconv.Atoi(os.Getenv("JWT_LIFE_TIME"))
// 	c.JwtConfig = JwtConfig{
// 		Key:    os.Getenv("JWT_KEY"),
// 		Durasi: time.Duration(longTime),
// 		Issuer: os.Getenv("JWT_ISSUER_NAME"),
// 	}

// 	// Inisialisasi struct Config dengan nilai dari environment variables
// 	return &Config{
// 		DBDriver:       getEnv("DB_DRIVER", ""),
// 		DBHost:         getEnv("DB_HOST", ""),
// 		DBPort:         getEnv("DB_PORT", ""),
// 		DBUser:         getEnv("DB_USER", ""),
// 		DBPassword:     getEnv("DB_PASSWORD", ""),
// 		DBName:         getEnv("DB_NAME", ""),
// 		JWTSecret:      getEnv("JWT_SECRET", ""),
// 		JWTExpiration:  getDurationEnv("JWT_EXPIRATION", 1),
// 		JWTIssuer:      getEnv("JWT_ISSUER", ""),
// 		OSMAPIEndpoint: getEnv("OSM_API_ENDPOINT", ""),
// 		OSMAPIKey:      getEnv("OSM_API_KEY", ""),
// 		LocationName:   getEnv("LOCATION_NAME", ""),
// 	}
// }

// // getEnv mengambil variabel lingkungan dengan nilai default
// func getEnv(key, defaultValue string) string {
// 	if value, exists := os.LookupEnv(key); exists {
// 		return value
// 	}
// 	return defaultValue
// }

// getDurationEnv mengambil variabel lingkungan dengan nilai default waktu
// func getDurationEnv(key string, defaultValue int) time.Duration {
// 	if value, exists := os.LookupEnv(key); exists {
// 		duration, err := time.ParseDuration(value + "h")
// 		if err != nil {
// 			return time.Duration(defaultValue) * time.Hour
// 		}
// 		return duration
// 	}
// 	return time.Duration(defaultValue) * time.Hour
// }

// var AppConfig = struct {
// 	OSMAPIEndpoint string
// }{}

// read configuration
func (c *Config) readConfig() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ")
	}

	longTime, _ := strconv.Atoi(os.Getenv("JWT_EXPIRATION"))
	c.JwtConfig = JwtConfig{
		JWTSecret:     os.Getenv("JWT_SECRET"),
		JWTExpiration: time.Duration(longTime),
		JWTIssuer:     os.Getenv("JWT_ISSUER"),
	}

	c.DbConfig = DbConfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBDriver:   os.Getenv("DB_DRIVER"),
	}

	c.OsmApiConfig = OsmApiConfig{
		OSMAPIEndpoint: os.Getenv("OSM_API_ENDPOINT"),
		OSMAPIKey:      os.Getenv("OSM_API_KEY"),
		LocationName:   os.Getenv("LOCATION_NAME"),
	}

	if c.JwtConfig.JWTSecret == "" || c.JwtConfig.JWTExpiration < 0 || c.JwtConfig.JWTIssuer == "" || c.DbConfig.DBHost == "" || c.DbConfig.DBPort == "" || c.DbConfig.DBUser == "" || c.DbConfig.DBPassword == "" || c.DbConfig.DBName == "" || c.DbConfig.DBDriver == "" || c.OsmApiConfig.OSMAPIEndpoint == "" || c.OsmApiConfig.OSMAPIKey == "" || c.OsmApiConfig.LocationName == "" {
		return errors.New("empty env")
	}
	return nil
}

func NewConfig() (*Config, error) {
	config := &Config{}
	err := config.readConfig()
	if err != nil {
		log.Fatal(err)
	}

	return config, nil

}
