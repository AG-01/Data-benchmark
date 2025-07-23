package services

import (
	"context"
	"database/sql"
	"fmt"
	"query-service/internal/config"
	"query-service/pkg/logger"

	_ "github.com/trinodb/trino-go-client/trino"
)

// TrinoService handles Trino-specific queries
type TrinoService struct {
	cfg    config.TrinoConfig
	logger *logger.Logger
	db     *sql.DB
}

// NewTrinoService creates a new TrinoService
func NewTrinoService(cfg config.TrinoConfig, logger *logger.Logger) (*TrinoService, error) {
	dsn := fmt.Sprintf("http://%s@%s:%s?catalog=%s&schema=%s",
		cfg.User, cfg.Host, cfg.Port, cfg.Catalog, cfg.Schema)

	db, err := sql.Open("trino", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Trino: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping Trino: %w", err)
	}

	return &TrinoService{
		cfg:    cfg,
		logger: logger,
		db:     db,
	}, nil
}

// ExecuteQuery executes a query on Trino
func (s *TrinoService) ExecuteQuery(ctx context.Context, query string) (*sql.Rows, error) {
	s.logger.WithField("query", query).Info("Executing Trino query")
	return s.db.QueryContext(ctx, query)
}

// GetStatus checks the status of the Trino service
func (s *TrinoService) GetStatus() error {
	return s.db.Ping()
}
