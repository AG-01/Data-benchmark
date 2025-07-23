package services

import (
	"context"
	"database/sql"
	"fmt"
	"query-service/internal/config"
	"query-service/pkg/logger"

	_ "github.com/go-sql-driver/mysql"
)

// StarRocksService handles StarRocks-specific queries
type StarRocksService struct {
	cfg    config.StarRocksConfig
	logger *logger.Logger
	db     *sql.DB
}

// NewStarRocksService creates a new StarRocksService
func NewStarRocksService(cfg config.StarRocksConfig, logger *logger.Logger) (*StarRocksService, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to StarRocks: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping StarRocks: %w", err)
	}

	return &StarRocksService{
		cfg:    cfg,
		logger: logger,
		db:     db,
	}, nil
}

// ExecuteQuery executes a query on StarRocks
func (s *StarRocksService) ExecuteQuery(ctx context.Context, query string) (*sql.Rows, error) {
	s.logger.WithField("query", query).Info("Executing StarRocks query")
	return s.db.QueryContext(ctx, query)
}

// GetStatus checks the status of the StarRocks service
func (s *StarRocksService) GetStatus() error {
	return s.db.Ping()
}
