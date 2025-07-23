package services

import (
	"github.com/sirupsen/logrus"
	"benchmark-api/internal/repository"
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
