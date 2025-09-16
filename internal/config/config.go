package config

import (
	"log"
	"os"

	"github.com/chinsiang99/simple-bank-go-project/pkg/confighelper"
	"github.com/joho/godotenv"
)

type Config struct {
	DB       *DBConfig
	APP      *AppConfig
	LOG      *LogConfig
	SECURITY *SecurityConfig
}

type DBConfig struct {
	Host       string
	Name       string
	User       string
	Pass       string
	Port       string
	Schema     string
	MaxOpenCon int
	MaxIdleCon int
}

type AppConfig struct {
	Environment string
	AppPort     string
}

type LogConfig struct {
	Level      string
	LogToFile  bool
	AppPath    string
	ErrPath    string
	MaxSize    int // in megabytes
	MaxBackups int // number of backups
	MaxAge     int //days
	Compress   bool
}

type SecurityConfig struct {
	CorsAllowOrigins     string
	CorsAllowMethods     string
	CorsAllowHeaders     string
	CorsExposeHeaders    string
	CorsAllowCredentials bool
	CorsMaxAge           string
}

var cfg Config

// init() is being called automatically when this pkg is imported & initialized
func init() {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}

	if env != "production" {
		// try load .env, but don’t panic if missing
		if err := godotenv.Load(); err != nil {
			log.Printf("No .env file found, relying on system env vars")
		}
	}

	cfg.DB = &DBConfig{
		Host:       confighelper.GetEnv("DB_HOST", ""),
		Name:       confighelper.GetEnv("DB_NAME", ""),
		User:       confighelper.GetEnv("DB_USER", ""),
		Pass:       confighelper.GetEnv("DB_PASS", ""),
		Port:       confighelper.GetEnv("DB_PORT", ""),
		Schema:     confighelper.GetEnv("DB_SCHEMA", ""),
		MaxOpenCon: confighelper.GetEnvAsInt("DB_MAX_OPEN_CON", 5),
		MaxIdleCon: confighelper.GetEnvAsInt("DB_MAX_IDLE_CON", 2),
	}

	cfg.APP = &AppConfig{
		Environment: confighelper.GetEnv("ENVIRONMENT", "development"),
		AppPort:     confighelper.GetEnv("APP_PORT", "9090"),
	}

	cfg.LOG = &LogConfig{
		Level:      confighelper.GetEnv("LOG_LEVEL", ""),
		LogToFile:  confighelper.GetEnvAsBool("LOG_TO_FILE", true),
		AppPath:    confighelper.GetEnv("LOG_APP_PATH", ""),
		ErrPath:    confighelper.GetEnv("LOG_ERR_PATH", ""),
		MaxSize:    confighelper.GetEnvAsInt("LOG_FILE_SIZE", 12),
		MaxBackups: confighelper.GetEnvAsInt("LOG_BACKUPS", 2),
		MaxAge:     confighelper.GetEnvAsInt("LOG_AGE", 7),
		Compress:   confighelper.GetEnvAsBool("LOG_COMPRESS", true),
	}

	cfg.SECURITY = &SecurityConfig{
		CorsAllowOrigins:     confighelper.GetEnv("CORS_ALLOW_ORIGINS", "*"),
		CorsAllowMethods:     confighelper.GetEnv("CORS_ALLOW_METHODS", "*"),
		CorsAllowHeaders:     confighelper.GetEnv("CORS_ALLOW_HEADERS", "*"),
		CorsExposeHeaders:    confighelper.GetEnv("CORS_EXPOSE_HEADERS", "*"),
		CorsAllowCredentials: confighelper.GetEnvAsBool("CORS_ALLOW_CREDENTIALS", false),
		CorsMaxAge:           confighelper.GetEnv("CORS_MAX_AGE", "*"),
	}
}

func New() *Config {
	return &cfg
}
