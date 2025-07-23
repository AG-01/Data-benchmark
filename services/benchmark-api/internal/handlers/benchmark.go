package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"benchmark-api/internal/models"
	"benchmark-api/internal/services"
)

type BenchmarkHandler struct {
	service *services.BenchmarkService
	logger  *logrus.Logger
}

func NewBenchmarkHandler(service *services.BenchmarkService, logger *logrus.Logger) *BenchmarkHandler {
	return &BenchmarkHandler{
		service: service,
		logger:  logger,
	}
}

// CreateBenchmark godoc
// @Summary Create a new benchmark
// @Description Create a new benchmark configuration
// @Tags benchmarks
// @Accept json
// @Produce json
// @Param benchmark body models.Benchmark true "Benchmark configuration"
// @Success 201 {object} models.Benchmark
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/benchmarks [post]
func (h *BenchmarkHandler) CreateBenchmark(c *gin.Context) {
	var benchmark models.Benchmark
	if err := c.ShouldBindJSON(&benchmark); err != nil {
		h.logger.WithError(err).Error("Failed to bind benchmark JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateBenchmark(&benchmark); err != nil {
		h.logger.WithError(err).Error("Failed to create benchmark")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create benchmark"})
		return
	}

	c.JSON(http.StatusCreated, benchmark)
}

// ListBenchmarks godoc
// @Summary List all benchmarks
// @Description Get a list of all benchmarks with optional filtering
// @Tags benchmarks
// @Produce json
// @Param status query string false "Filter by status"
// @Param table_format query string false "Filter by table format"
// @Param limit query int false "Limit number of results" default(20)
// @Param offset query int false "Offset for pagination" default(0)
// @Success 200 {array} models.Benchmark
// @Failure 500 {object} map[string]string
// @Router /api/v1/benchmarks [get]
func (h *BenchmarkHandler) ListBenchmarks(c *gin.Context) {
	filters := make(map[string]interface{})
	
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}
	if tableFormat := c.Query("table_format"); tableFormat != "" {
		filters["table_format"] = tableFormat
	}

	limit := 20
	if l := c.Query("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil {
			limit = parsed
		}
	}

	offset := 0
	if o := c.Query("offset"); o != "" {
		if parsed, err := strconv.Atoi(o); err == nil {
			offset = parsed
		}
	}

	benchmarks, err := h.service.ListBenchmarks(filters, limit, offset)
	if err != nil {
		h.logger.WithError(err).Error("Failed to list benchmarks")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list benchmarks"})
		return
	}

	c.JSON(http.StatusOK, benchmarks)
}

// GetBenchmark godoc
// @Summary Get a benchmark by ID
// @Description Get a specific benchmark by its ID
// @Tags benchmarks
// @Produce json
// @Param id path int true "Benchmark ID"
// @Success 200 {object} models.Benchmark
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/benchmarks/{id} [get]
func (h *BenchmarkHandler) GetBenchmark(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid benchmark ID"})
		return
	}

	benchmark, err := h.service.GetBenchmarkByID(uint(id))
	if err != nil {
		h.logger.WithError(err).Error("Failed to get benchmark")
		c.JSON(http.StatusNotFound, gin.H{"error": "Benchmark not found"})
		return
	}

	c.JSON(http.StatusOK, benchmark)
}

// UpdateBenchmark godoc
// @Summary Update a benchmark
// @Description Update an existing benchmark configuration
// @Tags benchmarks
// @Accept json
// @Produce json
// @Param id path int true "Benchmark ID"
// @Param benchmark body models.Benchmark true "Updated benchmark configuration"
// @Success 200 {object} models.Benchmark
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/benchmarks/{id} [put]
func (h *BenchmarkHandler) UpdateBenchmark(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid benchmark ID"})
		return
	}

	var benchmark models.Benchmark
	if err := c.ShouldBindJSON(&benchmark); err != nil {
		h.logger.WithError(err).Error("Failed to bind benchmark JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	benchmark.ID = uint(id)
	if err := h.service.UpdateBenchmark(&benchmark); err != nil {
		h.logger.WithError(err).Error("Failed to update benchmark")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update benchmark"})
		return
	}

	c.JSON(http.StatusOK, benchmark)
}

// DeleteBenchmark godoc
// @Summary Delete a benchmark
// @Description Delete a benchmark by ID
// @Tags benchmarks
// @Param id path int true "Benchmark ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/benchmarks/{id} [delete]
func (h *BenchmarkHandler) DeleteBenchmark(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid benchmark ID"})
		return
	}

	if err := h.service.DeleteBenchmark(uint(id)); err != nil {
		h.logger.WithError(err).Error("Failed to delete benchmark")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete benchmark"})
		return
	}

	c.Status(http.StatusNoContent)
}

// RunBenchmark godoc
// @Summary Run a benchmark
// @Description Start executing all queries in a benchmark across specified engines
// @Tags benchmarks
// @Param id path int true "Benchmark ID"
// @Success 202 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/benchmarks/{id}/run [post]
func (h *BenchmarkHandler) RunBenchmark(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid benchmark ID"})
		return
	}

	if err := h.service.RunBenchmark(uint(id)); err != nil {
		h.logger.WithError(err).Error("Failed to run benchmark")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to run benchmark"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Benchmark execution started"})
}

// GetBenchmarkStatus godoc
// @Summary Get benchmark status
// @Description Get the current status of a benchmark execution
// @Tags benchmarks
// @Produce json
// @Param id path int true "Benchmark ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/benchmarks/{id}/status [get]
func (h *BenchmarkHandler) GetBenchmarkStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid benchmark ID"})
		return
	}

	status, err := h.service.GetBenchmarkStatus(uint(id))
	if err != nil {
		h.logger.WithError(err).Error("Failed to get benchmark status")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get benchmark status"})
		return
	}

	c.JSON(http.StatusOK, status)
}

// GetBenchmarkResults godoc
// @Summary Get benchmark results
// @Description Get aggregated results for a benchmark
// @Tags benchmarks
// @Produce json
// @Param id path int true "Benchmark ID"
// @Success 200 {array} models.Result
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/benchmarks/{id}/results [get]
func (h *BenchmarkHandler) GetBenchmarkResults(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid benchmark ID"})
		return
	}

	results, err := h.service.GetBenchmarkResults(uint(id))
	if err != nil {
		h.logger.WithError(err).Error("Failed to get benchmark results")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get benchmark results"})
		return
	}

	c.JSON(http.StatusOK, results)
}
