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

	"query-service/internal/config"
	"query-service/internal/handlers"
	"query-service/internal/services"
	"query-service/pkg/logger"
	"query-service/pkg/metrics"
)

func main() {
	// Initialize logger
	logger := logger.New()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Initialize metrics
	metrics.Init()

	// Initialize services
	trinoService, err := services.NewTrinoService(cfg.Trino, logger)
	if err != nil {
		log.Fatal("Failed to initialize Trino service:", err)
	}
	prestoService, err := services.NewPrestoService(cfg.Presto, logger)
	if err != nil {
		log.Fatal("Failed to initialize Presto service:", err)
	}
	// starrocksService, err := services.NewStarRocksService(cfg.StarRocks, logger)
	// if err != nil {
	// 	log.Fatal("Failed to initialize StarRocks service:", err)
	// }
	queryExecutor := services.NewQueryExecutor(trinoService, prestoService, nil, logger)

	// Initialize handlers
	queryHandler := handlers.NewQueryHandler(queryExecutor, logger)
	healthHandler := handlers.NewHealthHandler(logger)

	// Setup Gin router
	router := setupRouter(cfg, queryHandler, healthHandler)

	// Start server
	// Create HTTP server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}	// Graceful server shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	logger.Info("Query Service started on port 8080")

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

func setupRouter(cfg *config.Config, queryHandler *handlers.QueryHandler, healthHandler *handlers.HealthHandler) *gin.Engine {
	if cfg.Server.Mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

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
		api.POST("/execute", queryHandler.ExecuteQuery)
		api.GET("/engines", queryHandler.ListEngines)
		api.GET("/engines/:engine/status", queryHandler.GetEngineStatus)
		api.POST("/engines/:engine/test", queryHandler.TestEngine)
	}

	return router
}
