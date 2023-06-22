package initializers

import "time"

type DatabaseEngineType string // "postgres" | "mysql" | "sqlite3" | "mssql"
const (
	Postgres  DatabaseEngineType = "postgres"
	Mysql     DatabaseEngineType = "mysql"
	Sqlite3   DatabaseEngineType = "sqlite3"
	SQLServer DatabaseEngineType = "mssql"
)

type DatabaseOptions struct {
	Engine   DatabaseEngineType
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
	FrontTypes     bool
	FrontTypesPath string
	AutoMigrate    bool
}

type WebServerOptions struct {
	TLS      bool
	CertFile string
	KeyFile  string
	Port     uint64
}

type JwtOptions struct {
	Issuer     string
	Secret     string
	Expiration time.Duration
	MaxAge     time.Duration
}

type ConfigOptions struct {
	Enviroment string
	AppName    string
	Database   DatabaseOptions
	Middleware MiddlewareOptions
	Generate   GenerateOptions
	Client     ClientOptions
	WebServer  WebServerOptions
	Logger     LoggerOptions
	Jwt        JwtOptions
}
