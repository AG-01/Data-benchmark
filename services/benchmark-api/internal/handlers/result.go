package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"benchmark-api/internal/services"
)

type ResultHandler struct {
	service *services.ResultService
	logger  *logrus.Logger
}

func NewResultHandler(service *services.ResultService, logger *logrus.Logger) *ResultHandler {
	return &ResultHandler{
		service: service,
		logger:  logger,
	}
}

// ListResults lists all results
func (h *ResultHandler) ListResults(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

// GetResult gets a result by ID
func (h *ResultHandler) GetResult(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

// CompareResults compares results
func (h *ResultHandler) CompareResults(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

// GetAnalytics gets analytics data
func (h *ResultHandler) GetAnalytics(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}
