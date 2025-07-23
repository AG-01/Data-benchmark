package services

import (
	"github.com/sirupsen/logrus"
)

type MetricService struct {
	prometheusURL string
	logger        *logrus.Logger
}

func NewMetricService(prometheusURL string, logger *logrus.Logger) *MetricService {
	return &MetricService{
		prometheusURL: prometheusURL,
		logger:        logger,
	}
}

// TODO: Implement metric service methods
