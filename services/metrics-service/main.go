package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"metrics-service/internal/config"
	"metrics-service/internal/handlers"
	"metrics-service/pkg/logger"
	"metrics-service/pkg/metrics"
)

func main() {
	// Initialize logger
	logger := logger.New()

	// Load configuration
	cfg := config.Load()

	// Initialize metrics
	metrics.Init()

	// Initialize handlers
	metricsHandler := handlers.NewMetricsHandler(logger)
	healthHandler := handlers.NewHealthHandler(logger)

	// Setup Gin router
	router := setupRouter(cfg, metricsHandler, healthHandler)

	// Start server
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	// Graceful server shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	logger.Info("Metrics Service started on port " + cfg.Port)

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	logger.Info("Server exited")
}

func setupRouter(cfg *config.Config, metricsHandler *handlers.MetricsHandler, healthHandler *handlers.HealthHandler) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Health check
	router.GET("/health", healthHandler.HealthCheck)
	router.GET("/ready", healthHandler.ReadinessCheck)

	// Metrics endpoint
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// API routes
	api := router.Group("/api/v1")
	{
		api.GET("/metrics", metricsHandler.GetMetrics)
		api.GET("/metrics/benchmark/:id", metricsHandler.GetBenchmarkMetrics)
		api.GET("/metrics/query/:id", metricsHandler.GetQueryMetrics)
		api.GET("/metrics/resource/:service", metricsHandler.GetResourceMetrics)
	}

	return router
}
