package services

import (
	"benchmark-api/internal/repository"
	"github.com/sirupsen/logrus"
)

type ResultService struct {
	repo   *repository.ResultRepository
	logger *logrus.Logger
}

func NewResultService(repo *repository.ResultRepository, logger *logrus.Logger) *ResultService {
	return &ResultService{
		repo:   repo,
		logger: logger,
	}
}

// TODO: Implement result service methods
