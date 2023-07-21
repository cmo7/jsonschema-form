package config

import (
	"log"
	"nartex/ngr-stack/i18n"
	"time"

	"github.com/joho/godotenv"
)

// Different options for the app are loaded into these variables
// The options are loaded from the .env file corresponding to the enviroment
// The main .env file is loaded first, then the .env.{enviroment} file is loaded
// Values from the .env.{enviroment} file override the values from the main .env file
// No config value is mandatory, if a value is missing, a default value is used

var (
	App        *AppOptions
	Database   *DatabaseOptions
	Debug      *DebugOptions
	Generate   *GenerateOptions
	Middleware *MiddlewareOptions
	WebServer  *WebServerOptions
	Client     *ClientOptions
	Jwt        *JwtOptions
	Logger     *LoggerOptions
	Pagination *PaginationOptions
)

// The values from the config file are parsed into corresponding types
// Missing values are replaced with default values, if any
func LoadConfig() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading main .env file")
		log.Fatal("Development enviroment is assumed")
	}
	enviroment := getEnviroment("ENVIROMENT", Development)

	err = godotenv.Load(".env." + enviroment)
	if err != nil {
		log.Fatalf("Error loading .env.%s file", enviroment)
		log.Fatal("Default values in the main .env file and hardcoded defaults are used as fallback")
	}

	App = &AppOptions{
		Enviroment: getEnvStr("ENVIROMENT", "development"),
		AppName:    getEnvStr("APP_NAME", "Nartex Go App"),
		Locale:     getEnvLocale("LOCALE", i18n.EN),
	}

	Database = &DatabaseOptions{
		Engine:   getEnvDatabaseEngine("DB_ENGINE", Postgres),
		Host:     getEnvStr("DB_HOST", "localhost"),
		Port:     getEnvUInt("DB_PORT", 5432),
		User:     getEnvStr("DB_USER", "postgres"),
		Password: getEnvStr("DB_PASSWORD", "postgres"),
		Database: getEnvStr("DB_DATABASE", "postgres"),
	}

	Middleware = &MiddlewareOptions{
		Helmet:   getEnvBool("MID_HELMET", true),
		Compress: getEnvBool("MID_COMPRESS", true),
		Cors:     getEnvBool("MID_CORS", true),
		Logger:   getEnvBool("MID_LOGGER", true),
		Cache:    getEnvBool("MID_CACHE", true),
	}

	Generate = &GenerateOptions{
		FrontTypes:     getEnvBool("GEN_FRONT_TYPES", true),
		FrontTypesPath: getEnvStr("GEN_FRONT_TYPES_PATH", "../../front/src/types.ts"),
		AutoMigrate:    getEnvBool("GEN_AUTO_MIGRATE", true),
	}

	Client = &ClientOptions{
		Mode: getEnvClientMode("CLIENT_MODE", Internal),
		URL:  getEnvStr("CLIENT_URL", "https://localhost:5173"),
	}

	WebServer = &WebServerOptions{
		TLS:      getEnvBool("TLS_ENABLED", true),
		CertFile: getEnvStr("TLS_CERT", "./certs/cert.pem"),
		KeyFile:  getEnvStr("TLS_KEY", "./certs/key.pem"),
		Port:     getEnvUInt("TLS_PORT", 5173),
	}

	Jwt = &JwtOptions{
		Issuer:     getEnvStr("JWT_ISSUER", "Nartex"),
		Secret:     getEnvStr("JWT_SECRET", "secret"),
		Expiration: getEnvDuration("JWT_EXPIRATION", 24*time.Hour),
		MaxAge:     getEnvDuration("JWT_MAX_AGE", 24*time.Hour),
	}

	Logger = &LoggerOptions{
		MainLogger:  getEnvBool("LOG_ENABLED", true),
		QueryLogger: getEnvBool("DB_QUERY_LOGGING", true),
	}

	Debug = &DebugOptions{
		DevTools:     getEnvBool("DEV_TOOLS", false),
		DatabaseSeed: getEnvBool("DATABASE_SEED", false),
	}

	Pagination = &PaginationOptions{
		MaxPageSize:     getEnvInt("PAGINATION_MAX_PAGE_SIZE", 100),
		DefaultPageSize: getEnvInt("PAGINATION_DEFAULT_PAGE_SIZE", 10),
	}
}
