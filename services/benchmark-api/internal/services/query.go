package services

import (
	"github.com/sirupsen/logrus"
	"benchmark-api/internal/config"
	"benchmark-api/internal/repository"
)

type QueryService struct {
	repo   *repository.QueryRepository
	config *config.Config
	logger *logrus.Logger
}

func NewQueryService(repo *repository.QueryRepository, config *config.Config, logger *logrus.Logger) *QueryService {
	return &QueryService{
		repo:   repo,
		config: config,
		logger: logger,
	}
}

// TODO: Implement query service methods
