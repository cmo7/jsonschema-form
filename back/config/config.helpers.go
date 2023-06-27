package config

import (
	"nartex/ngr-stack/i18n"
	"os"
	"strconv"
	"time"
)

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
func getEnvInt(key string, defaultValue int) int {
	s := os.Getenv(key)
	if s == "" {
		return defaultValue
	}
	v, err := strconv.Atoi(s)
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

// getEnviroment returns the value of the enviroment variable with the given key, if the variable is empty returns the defaultValue
func getEnviroment(key string, defaultValue EnviromentType) string {
	s := os.Getenv(key)
	if s == "" {
		return string(defaultValue)
	}
	return s
}

func getEnvLocale(key string, defaultValue i18n.Locale) i18n.Locale {
	s := os.Getenv(key)
	if s == "" {
		return defaultValue
	}
	return i18n.Locale(s)
}
