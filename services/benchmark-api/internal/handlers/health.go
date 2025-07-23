package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type HealthHandler struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewHealthHandler(db *gorm.DB, logger *logrus.Logger) *HealthHandler {
	return &HealthHandler{
		db:     db,
		logger: logger,
	}
}

// HealthCheck godoc
// @Summary Health check endpoint
// @Description Check if the service is healthy
// @Tags health
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	response := map[string]interface{}{
		"status":  "healthy",
		"service": "benchmark-api",
		"version": "1.0.0",
	}

	// Check database connection
	sqlDB, err := h.db.DB()
	if err != nil {
		h.logger.WithError(err).Error("Failed to get database connection")
		response["status"] = "unhealthy"
		response["database"] = "disconnected"
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	if err := sqlDB.Ping(); err != nil {
		h.logger.WithError(err).Error("Database ping failed")
		response["status"] = "unhealthy"
		response["database"] = "unreachable"
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	response["database"] = "connected"
	c.JSON(http.StatusOK, response)
}

// ReadinessCheck godoc
// @Summary Readiness check endpoint
// @Description Check if the service is ready to accept requests
// @Tags health
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 503 {object} map[string]interface{}
// @Router /ready [get]
func (h *HealthHandler) ReadinessCheck(c *gin.Context) {
	response := map[string]interface{}{
		"status":  "ready",
		"service": "benchmark-api",
	}

	// Check database connection
	sqlDB, err := h.db.DB()
	if err != nil {
		h.logger.WithError(err).Error("Failed to get database connection")
		response["status"] = "not ready"
		response["database"] = "disconnected"
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	if err := sqlDB.Ping(); err != nil {
		h.logger.WithError(err).Error("Database ping failed")
		response["status"] = "not ready"
		response["database"] = "unreachable"
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	response["database"] = "ready"
	c.JSON(http.StatusOK, response)
}
