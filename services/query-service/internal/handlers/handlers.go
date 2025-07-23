package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"query-service/internal/services"
	"query-service/pkg/logger"
)

type QueryHandler struct {
	executor *services.QueryExecutor
	logger   *logger.Logger
}

// NewQueryHandler creates a new QueryHandler
func NewQueryHandler(executor *services.QueryExecutor, logger *logger.Logger) *QueryHandler {
	return &QueryHandler{
		executor: executor,
		logger:   logger,
	}
}

func (h *QueryHandler) ExecuteQuery(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (h *QueryHandler) ListEngines(c *gin.Context) {
	engines := []map[string]interface{}{
		{"name": "trino", "status": "active"},
		{"name": "presto", "status": "active"},
		{"name": "starrocks", "status": "active"},
	}
	c.JSON(http.StatusOK, engines)
}

func (h *QueryHandler) GetEngineStatus(c *gin.Context) {
	engine := c.Param("engine")
	c.JSON(http.StatusOK, gin.H{"engine": engine, "status": "healthy"})
}

func (h *QueryHandler) TestEngine(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "test successful"})
}

type HealthHandler struct {
	logger *logger.Logger
}

// NewHealthHandler creates a new HealthHandler
func NewHealthHandler(logger *logger.Logger) *HealthHandler {
	return &HealthHandler{logger: logger}
}

func (h *HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy", "service": "query-service"})
}

func (h *HealthHandler) ReadinessCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ready", "service": "query-service"})
}
