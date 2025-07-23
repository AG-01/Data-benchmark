package services

import (
	"context"
	"database/sql"
	"fmt"
	"query-service/internal/config"
	"query-service/pkg/logger"

	_ "github.com/prestodb/presto-go-client/presto"
)

// PrestoService handles Presto-specific queries
type PrestoService struct {
	cfg    config.PrestoConfig
	logger *logger.Logger
	db     *sql.DB
}

// NewPrestoService creates a new PrestoService
func NewPrestoService(cfg config.PrestoConfig, logger *logger.Logger) (*PrestoService, error) {
	dsn := fmt.Sprintf("http://%s@%s:%s?catalog=%s&schema=%s",
		cfg.User, cfg.Host, cfg.Port, cfg.Catalog, cfg.Schema)

	db, err := sql.Open("presto", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Presto: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping Presto: %w", err)
	}

	return &PrestoService{
		cfg:    cfg,
		logger: logger,
		db:     db,
	}, nil
}

// ExecuteQuery executes a query on Presto
func (s *PrestoService) ExecuteQuery(ctx context.Context, query string) (*sql.Rows, error) {
	s.logger.WithField("query", query).Info("Executing Presto query")
	return s.db.QueryContext(ctx, query)
}

// GetStatus checks the status of the Presto service
func (s *PrestoService) GetStatus() error {
	return s.db.Ping()
}
