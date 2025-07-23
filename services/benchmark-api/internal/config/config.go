package config

import (
	"os"
	"strconv"
)

type Config struct {
	Server     ServerConfig
	Database   DatabaseConfig
	MinIO      MinIOConfig
	Engines    EnginesConfig
	Prometheus PrometheusConfig
}

type ServerConfig struct {
	Port string
	Mode string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type MinIOConfig struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	UseSSL    bool
}

type EnginesConfig struct {
	Trino     EngineConfig
	Presto    EngineConfig
	StarRocks EngineConfig
}

type EngineConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

type PrometheusConfig struct {
	URL string
}

func Load() (*Config, error) {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
			Mode: getEnv("GIN_MODE", "debug"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "hive"),
			Password: getEnv("DB_PASSWORD", "hivepass"),
			DBName:   getEnv("DB_NAME", "hive_metastore"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		MinIO: MinIOConfig{
			Endpoint:  getEnv("MINIO_ENDPOINT", "localhost:9000"),
			AccessKey: getEnv("MINIO_ACCESS_KEY", "admin"),
			SecretKey: getEnv("MINIO_SECRET_KEY", "password"),
			UseSSL:    getEnvBool("MINIO_USE_SSL", false),
		},
		Engines: EnginesConfig{
			Trino: EngineConfig{
				Host:     getEnv("TRINO_HOST", "localhost:8081"),
				Username: getEnv("TRINO_USERNAME", "admin"),
				Password: getEnv("TRINO_PASSWORD", ""),
			},
			Presto: EngineConfig{
				Host:     getEnv("PRESTO_HOST", "localhost:8082"),
				Username: getEnv("PRESTO_USERNAME", "admin"),
				Password: getEnv("PRESTO_PASSWORD", ""),
			},
			StarRocks: EngineConfig{
				Host:     getEnv("STARROCKS_HOST", "localhost:8030"),
				Username: getEnv("STARROCKS_USERNAME", "root"),
				Password: getEnv("STARROCKS_PASSWORD", ""),
			},
		},
		Prometheus: PrometheusConfig{
			URL: getEnv("PROMETHEUS_URL", "http://localhost:9090"),
		},
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if parsed, err := strconv.ParseBool(value); err == nil {
			return parsed
		}
	}
	return defaultValue
}
