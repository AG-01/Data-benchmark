package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Config holds the application configuration
type Config struct {
	Server    ServerConfig
	Trino     TrinoConfig
	Presto    PrestoConfig
	StarRocks StarRocksConfig
	Logger    *logrus.Logger
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port string
	Mode string
}

// TrinoConfig holds Trino connection details
type TrinoConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Catalog  string
	Schema   string
}

// PrestoConfig holds Presto connection details
type PrestoConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Catalog  string
	Schema   string
}

// StarRocksConfig holds StarRocks connection details
type StarRocksConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// Load creates a new configuration object
func Load() (*Config, error) {
	cfg := &Config{
		Server: ServerConfig{
			Port: getEnv("QUERY_SERVICE_PORT", "8080"),
			Mode: getEnv("GIN_MODE", "release"),
		},
		Trino: TrinoConfig{
			Host:    getEnv("TRINO_HOST", "trino"),
			Port:    getEnv("TRINO_PORT", "8080"),
			User:    getEnv("TRINO_USER", "admin"),
			Catalog: getEnv("TRINO_CATALOG", "hive"),
			Schema:  getEnv("TRINO_SCHEMA", "default"),
		},
		Presto: PrestoConfig{
			Host:    getEnv("PRESTO_HOST", "presto"),
			Port:    getEnv("PRESTO_PORT", "8080"),
			User:    getEnv("PRESTO_USER", "admin"),
			Catalog: getEnv("PRESTO_CATALOG", "hive"),
			Schema:  getEnv("PRESTO_SCHEMA", "default"),
		},
		StarRocks: StarRocksConfig{
			Host:     getEnv("STARROCKS_HOST", "starrocks-fe"),
			Port:     getEnv("STARROCKS_PORT", "9030"),
			User:     getEnv("STARROCKS_USER", "root"),
			Password: getEnv("STARROCKS_PASSWORD", ""),
			Database: getEnv("STARROCKS_DATABASE", "default_catalog"),
		},
		Logger: logrus.New(),
	}

	// Configure logger
	cfg.Logger.SetFormatter(&logrus.JSONFormatter{})
	cfg.Logger.SetLevel(logrus.InfoLevel)
	cfg.Logger.SetOutput(os.Stdout)

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
