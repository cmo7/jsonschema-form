package config

import (
	"nartex/ngr-stack/i18n"
	"time"
)

type EnviromentType string // "local" | "development" | "preproduction" | "production" | "container"
const (
	Local         = "local"
	Development   = "development"
	Preproduction = "preproduction"
	Production    = "production"
	Container     = "container"
)

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

type AppOptions struct {
	Enviroment string
	AppName    string
	Locale     i18n.Locale
}

type DebugOptions struct {
	DevTools     bool
	DatabaseSeed bool
}

type PaginationOptions struct {
	MaxPageSize     int
	DefaultPageSize int
}
