package services

import (
	"benchmark-api/internal/config"
	"benchmark-api/internal/repository"
	"github.com/sirupsen/logrus"
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
