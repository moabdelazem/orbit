package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	// Server configs
	Port            string
	LogLevel        string
	Env             string
	APITimeout      time.Duration
	APIReadTimeout  time.Duration
	APIWriteTimeout time.Duration
	APIIdleTimeout  time.Duration

	// Database configs
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	// Redis configs
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       string

	// JWT configs
	JWTSecret          string
	AccessTokenExpiry  time.Duration
	RefreshTokenExpiry time.Duration
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	// Load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Server configuration
	port := getEnv("PORT", "8080")
	logLevel := getEnv("LOG_LEVEL", "info")
	env := getEnv("ENV", "development")
	apiTimeout := getEnvDuration("API_TIMEOUT", 30*time.Second)
	apiReadTimeout := getEnvDuration("API_READ_TIMEOUT", 15*time.Second)
	apiWriteTimeout := getEnvDuration("API_WRITE_TIMEOUT", 15*time.Second)
	apiIdleTimeout := getEnvDuration("API_IDLE_TIMEOUT", 60*time.Second)

	// Database configuration
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "orbit")
	dbSSLMode := getEnv("DB_SSL_MODE", "disable")

	// Redis configuration
	redisHost := getEnv("REDIS_HOST", "localhost")
	redisPort := getEnv("REDIS_PORT", "6379")
	redisPassword := getEnv("REDIS_PASSWORD", "")
	redisDB := getEnv("REDIS_DB", "0")

	// JWT configuration
	jwtSecret := getEnv("JWT_SECRET", "your_jwt_secret_key_here")
	accessTokenExpiry := getEnvDuration("ACCESS_TOKEN_EXPIRY", 15*time.Minute)
	refreshTokenExpiry := getEnvDuration("REFRESH_TOKEN_EXPIRY", 7*24*time.Hour)

	return &Config{
		// Server configs
		Port:            port,
		LogLevel:        logLevel,
		Env:             env,
		APITimeout:      apiTimeout,
		APIReadTimeout:  apiReadTimeout,
		APIWriteTimeout: apiWriteTimeout,
		APIIdleTimeout:  apiIdleTimeout,

		// Database configs
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     dbName,
		DBSSLMode:  dbSSLMode,

		// Redis configs
		RedisHost:     redisHost,
		RedisPort:     redisPort,
		RedisPassword: redisPassword,
		RedisDB:       redisDB,

		// JWT configs
		JWTSecret:          jwtSecret,
		AccessTokenExpiry:  accessTokenExpiry,
		RefreshTokenExpiry: refreshTokenExpiry,
	}
}

// Helper function to get environment variable with a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// Helper function to get environment variable as duration
func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		log.Printf("Warning: invalid duration for %s, using default", key)
		return defaultValue
	}

	return duration
}
