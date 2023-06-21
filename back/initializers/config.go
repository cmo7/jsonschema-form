package initializers

import (
	"log"
	"os"

	"strconv"

	"github.com/joho/godotenv"
)

type DatabaseOptions struct {
	Engine   string
	Host     string
	Port     uint64
	User     string
	Password string
	Database string
}

type LoggerOptions struct {
	MainLogger  bool
	QueryLogger bool
}

type MiddlewareOptions struct {
	Helmet   bool
	Compress bool
	Cors     bool
	Logger   bool
	Cache    bool
}

type ClientModeType string // "internal" | "external"
const (
	Internal ClientModeType = "internal"
	External ClientModeType = "external"
)

type ClientOptions struct {
	Mode ClientModeType
	URL  string
}

type GenerateOptions struct {
	FrontTypes  bool
	AutoMigrate bool
}

type WebServerOptions struct {
	TLS      bool
	CertFile string
	KeyFile  string
	Port     uint64
}

type ConfigType struct {
	Enviroment string
	AppName    string
	Database   DatabaseOptions
	Middleware MiddlewareOptions
	Generate   GenerateOptions
	Client     ClientOptions
	WebServer  WebServerOptions
	Logger     LoggerOptions
}

var Config ConfigType

func LoadConfig(enviroment string) {

	err := godotenv.Load(".env." + enviroment)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var Config = new(ConfigType)
	// App config
	Config.Enviroment = os.Getenv("ENVIROMENT")
	Config.AppName = os.Getenv("APP_NAME")

	// Database config
	Config.Database.Engine = os.Getenv("DB_ENGINE")
	Config.Database.Host = os.Getenv("DB_HOST")
	p := os.Getenv("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Fatal("Error parsing port")
	}
	Config.Database.Port = port
	Config.Database.User = os.Getenv("DB_USER")
	Config.Database.Password = os.Getenv("DB_PASSWORD")
	Config.Database.Database = os.Getenv("DB_DATABASE")

	// Middleware config
	Config.Middleware.Helmet = toBool(os.Getenv("MID_HELMET"), true)
	Config.Middleware.Compress = toBool(os.Getenv("MID_COMPRESS"), true)
	Config.Middleware.Cors = toBool(os.Getenv("MID_CORS"), true)
	Config.Middleware.Cache = toBool(os.Getenv("MID_CACHE"), true)

	// Codegen config
	Config.Generate.FrontTypes = toBool(os.Getenv("GENERATE_FRONT_TYPES"), true)
	Config.Generate.AutoMigrate = toBool(os.Getenv("DB_AUTO_MIGRATE"), true)

	// Client config
	Config.Client.Mode = ClientModeType(os.Getenv("CLIENT_MODE"))
	Config.Client.URL = os.Getenv("CLIENT_URL")

	// WebServer config
	Config.WebServer.TLS = toBool(os.Getenv("TLS"), true)
	Config.WebServer.CertFile = os.Getenv("CERT_FILE")
	Config.WebServer.KeyFile = os.Getenv("KEY_FILE")

	if Config.WebServer.TLS {
		Config.WebServer.Port = toUInt(os.Getenv("TLS_PORT"), 8443)
	} else {
		Config.WebServer.Port = toUInt(os.Getenv("PORT"), 8080)
	}

	// Logger config
	Config.Logger.MainLogger = toBool(os.Getenv("MID_LOGGER"), true)
	Config.Logger.QueryLogger = toBool(os.Getenv("DB_QUERY_LOGGING"), true)
}

func toBool(s string, defaultValue bool) bool {
	if s == "" {
		return defaultValue
	}
	if s == "true" {
		return true
	}
	return false
}

func toUInt(s string, defaultValue uint64) uint64 {
	if s == "" {
		return defaultValue
	}
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return defaultValue
	}
	return v
}
