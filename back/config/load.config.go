package config

import (
	"log"
	"os"
	"time"

	"strconv"

	"github.com/joho/godotenv"
)

// Options is the struct that holds the config values
// The values are loaded from .env file using LoadConfig function
var Options *ConfigOptions

// LoadConfig loads the config from .env file corresponding to the enviroment into Config variable
// The enviroment is passed as a parameter
// The values from the config file are parsed into corresponding types
// Missing values are replaced with default values, if any
func LoadConfig(enviroment string) {

	err := godotenv.Load(".env." + enviroment)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Options = &ConfigOptions{
		Enviroment: getEnvStr("ENVIROMENT", "development"),
		AppName:    getEnvStr("APP_NAME", "Nartex Go App"),
		DevTools:   getEnvBool("DEV_TOOLS", false),
		Database: DatabaseOptions{
			Engine:   getEnvDatabaseEngine("DB_ENGINE", Postgres),
			Host:     getEnvStr("DB_HOST", "localhost"),
			Port:     getEnvUInt("DB_PORT", 5432),
			User:     getEnvStr("DB_USER", "postgres"),
			Password: getEnvStr("DB_PASSWORD", "postgres"),
			Database: getEnvStr("DB_DATABASE", "postgres"),
		},
		Middleware: MiddlewareOptions{
			Helmet:   getEnvBool("MID_HELMET", true),
			Compress: getEnvBool("MID_COMPRESS", true),
			Cors:     getEnvBool("MID_CORS", true),
			Cache:    getEnvBool("MID_CACHE", true),
		},
		Generate: GenerateOptions{
			FrontTypes:     getEnvBool("GENERATE_FRONT_TYPES", true),
			FrontTypesPath: getEnvStr("GENERATE_FRONT_TYPES_PATH", "../front/src/types/generated"),
			AutoMigrate:    getEnvBool("DB_AUTO_MIGRATE", true),
		},
		Client: ClientOptions{
			Mode: getEnvClientMode("CLIENT_MODE", Internal),
			URL:  getEnvStr("CLIENT_URL", "https://localhost:5173"),
		},
		WebServer: WebServerOptions{
			TLS:      getEnvBool("WEB_SERVER_TLS", true),
			CertFile: getEnvStr("TLS_CERT", "./certs/cert.pem"),
			KeyFile:  getEnvStr("TLS_KEY", "./certs/key.pem"),
			Port:     getEnvUInt("TLS_PORT", 8443),
		},
		Jwt: JwtOptions{
			Issuer:     getEnvStr("JWT_ISSUER", "Nartex Go App"),
			Secret:     getEnvStr("JWT_SECRET", "secret"),
			Expiration: getEnvDuration("JWT_EXPIRATION", 24*time.Hour),
			MaxAge:     getEnvDuration("JWT_MAX_AGE", 24*time.Hour),
		},
		Logger: LoggerOptions{
			MainLogger:  getEnvBool("LOG_ENABLED", true),
			QueryLogger: getEnvBool("DB_QUERY_LOGGING", true),
		},
	}
}

// Helper functions to get enviroment variables with default values

// getEnvStr returns the value of the enviroment variable with the given key, if the variable is empty returns the defaultValue
func getEnvStr(key string, defaultValue string) string {
	s := os.Getenv(key)
	if s == "" {
		return defaultValue
	}
	return s
}

// getEnvBool returns the value of the enviroment variable with the given key, if the variable is empty returns the defaultValue
func getEnvBool(key string, defaultValue bool) bool {
	s := os.Getenv(key)
	if s == "" {
		return defaultValue
	}
	if s == "true" {
		return true
	}
	return false
}

// getEnvUInt returns the value of the enviroment variable with the given key, if the variable is empty returns the defaultValue
func getEnvUInt(key string, defaultValue uint64) uint64 {
	s := os.Getenv(key)
	if s == "" {
		return defaultValue
	}
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return defaultValue
	}
	return v
}

// getEnvInt returns the value of the enviroment variable with the given key, if the variable is empty returns the defaultValue
func getEnvInt(key string, defaultValue int64) int64 {
	s := os.Getenv(key)
	if s == "" {
		return defaultValue
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return defaultValue
	}
	return v
}

// getEnvFloat returns the value of the enviroment variable with the given key, if the variable is empty returns the defaultValue
func getEnvFloat(key string, defaultValue float64) float64 {
	s := os.Getenv(key)
	if s == "" {
		return defaultValue
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return defaultValue
	}
	return v
}

// getEnvDuration returns the value of the enviroment variable with the given key, if the variable is empty returns the defaultValue
func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	s := os.Getenv(key)
	if s == "" {
		return defaultValue
	}
	v, err := time.ParseDuration(s)
	if err != nil {
		return defaultValue
	}
	return v
}

// getEnvTime returns the value of the enviroment variable with the given key, if the variable is empty returns the defaultValue
func getEnvTime(key string, defaultValue time.Time) time.Time {
	s := os.Getenv(key)
	if s == "" {
		return defaultValue
	}
	v, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return defaultValue
	}
	return v
}

// getEnvDatabaseEngine returns the value of the enviroment variable with the given key, if the variable is empty returns the defaultValue
func getEnvDatabaseEngine(key string, defaultValue DatabaseEngineType) DatabaseEngineType {
	s := os.Getenv(key)
	if s == "" {
		return defaultValue
	}
	return DatabaseEngineType(s)
}

// getEnvClientMode returns the value of the enviroment variable with the given key, if the variable is empty returns the defaultValue
func getEnvClientMode(key string, defaultValue ClientModeType) ClientModeType {
	s := os.Getenv(key)
	if s == "" {
		return defaultValue
	}
	return ClientModeType(s)
}
