package handlers

import (
	"benchmark-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type QueryHandler struct {
	service *services.QueryService
	logger  *logrus.Logger
}

func NewQueryHandler(service *services.QueryService, logger *logrus.Logger) *QueryHandler {
	return &QueryHandler{
		service: service,
		logger:  logger,
	}
}

func (*QueryHandler) CreateQuery(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (*QueryHandler) ListQueries(c *gin.Context) {
	c.JSON(http.StatusOK, []interface{}{})
}

func (*QueryHandler) GetQuery(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (*QueryHandler) UpdateQuery(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (*QueryHandler) DeleteQuery(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (*QueryHandler) ExecuteQuery(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (*QueryHandler) GetQueryResults(c *gin.Context) {
	c.JSON(http.StatusOK, []interface{}{})
}

func (*QueryHandler) ListEngines(c *gin.Context) {
	engines := []map[string]interface{}{
		{"name": "trino", "type": "trino", "status": "active"},
		{"name": "presto", "type": "presto", "status": "active"},
		{"name": "starrocks", "type": "starrocks", "status": "active"},
	}
	c.JSON(http.StatusOK, engines)
}

func (*QueryHandler) GetEngineStatus(c *gin.Context) {
	engine := c.Param("engine")
	status := map[string]interface{}{
		"engine": engine,
		"status": "healthy",
		"uptime": "1h 30m",
	}
	c.JSON(http.StatusOK, status)
}

func (*QueryHandler) ListTableFormats(c *gin.Context) {
	formats := []string{"hive", "iceberg"}
	c.JSON(http.StatusOK, formats)
}

func (*QueryHandler) CreateTable(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (*QueryHandler) GetTableInfo(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}
