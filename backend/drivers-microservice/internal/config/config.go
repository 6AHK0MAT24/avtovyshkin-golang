package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Database  DatabaseConfig
	Server    ServerConfig
	Upload    UploadConfig
	WebSocket WebSocketConfig
	CORS      CORSConfig
	Env       string
	LogLevel  string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type ServerConfig struct {
	Port         string
	Host         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type UploadConfig struct {
	Dir           string
	MaxUploadSize int64
	AllowedTypes  []string
}

type WebSocketConfig struct {
	ReadBufferSize  int
	WriteBufferSize int
	PingPeriod      time.Duration
	PongWait        time.Duration
	WriteWait       time.Duration
}

type CORSConfig struct {
	AllowedOrigins []string
	AllowedMethods []string
	AllowedHeaders []string
}

func Load() *Config {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	return &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			User:     getEnv("DB_USER", "avtovyshkin_db_go"),
			Password: getEnv("DB_PASSWORD", "avtovyshkin_db_go_admin"),
			DBName:   getEnv("DB_NAME", "avtovyshkin_db_go"),
		},
		Server: ServerConfig{
			Port:         getEnv("SERVER_PORT", "8080"),
			Host:         getEnv("SERVER_HOST", "0.0.0.0"),
			ReadTimeout:  getDurationEnv("READ_TIMEOUT", 15*time.Second),
			WriteTimeout: getDurationEnv("WRITE_TIMEOUT", 15*time.Second),
			IdleTimeout:  getDurationEnv("IDLE_TIMEOUT", 60*time.Second),
		},
		Upload: UploadConfig{
			Dir:           getEnv("UPLOAD_DIR", "./uploads"),
			MaxUploadSize: getInt64Env("MAX_UPLOAD_SIZE", 10485760), // 10MB
			AllowedTypes:  []string{"jpg", "jpeg", "png", "gif", "webp", "bmp", "tiff", "svg", "pdf"},
		}, WebSocket: WebSocketConfig{
			ReadBufferSize:  getIntEnv("WS_READ_BUFFER_SIZE", 1024),
			WriteBufferSize: getIntEnv("WS_WRITE_BUFFER_SIZE", 1024),
			PingPeriod:      getDurationEnv("WS_PING_PERIOD", 54*time.Second),
			PongWait:        getDurationEnv("WS_PONG_WAIT", 60*time.Second),
			WriteWait:       getDurationEnv("WS_WRITE_WAIT", 10*time.Second),
		},
		CORS: CORSConfig{
			AllowedOrigins: getSliceEnv("CORS_ALLOWED_ORIGINS", []string{"http://localhost:5173", "http://localhost:3000"}),
			AllowedMethods: getSliceEnv("CORS_ALLOWED_METHODS", []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			AllowedHeaders: getSliceEnv("CORS_ALLOWED_HEADERS", []string{"Content-Type", "Authorization", "X-Requested-With"}),
		}, Env: getEnv("ENV", "development"),
		LogLevel: getEnv("LOG_LEVEL", "debug"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getIntEnv(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

func getInt64Env(key string, defaultValue int64) int64 {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.ParseInt(value, 10, 64); err == nil {
			return intVal
		}
	}
	return defaultValue
}

func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}

func getSliceEnv(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		return strings.Split(value, ",")
	}
	return defaultValue
}
