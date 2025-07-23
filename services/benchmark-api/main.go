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
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "benchmark-api/docs"
	"benchmark-api/internal/config"
	"benchmark-api/internal/handlers"
	"benchmark-api/internal/middleware"
	"benchmark-api/internal/repository"
	"benchmark-api/internal/services"
	"benchmark-api/pkg/database"
	"benchmark-api/pkg/logger"
	"benchmark-api/pkg/metrics"
)

// @title Data Lake Benchmark API
// @version 1.0
// @description API for benchmarking data lake table formats across multiple query engines
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

func main() {
	// Initialize logger
	logger := logger.New()
	
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Initialize database
	db, err := database.NewConnection(cfg.Database)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	
	// Get underlying sql.DB to close properly
	defer func() {
		if sqlDB, err := db.DB(); err == nil {
			sqlDB.Close()
		}
	}()

	// Initialize metrics
	metrics.Init()

	// Initialize repositories
	benchmarkRepo := repository.NewBenchmarkRepository(db)
	queryRepo := repository.NewQueryRepository(db)
	resultRepo := repository.NewResultRepository(db)

	// Initialize services
	benchmarkService := services.NewBenchmarkService(benchmarkRepo, logger)
	queryService := services.NewQueryService(queryRepo, cfg, logger)
	resultService := services.NewResultService(resultRepo, logger)
	// metricService := services.NewMetricService(cfg.Prometheus.URL, logger) // TODO: Use this service

	// Initialize handlers
	benchmarkHandler := handlers.NewBenchmarkHandler(benchmarkService, logger)
	queryHandler := handlers.NewQueryHandler(queryService, logger)
	resultHandler := handlers.NewResultHandler(resultService, logger)
	healthHandler := handlers.NewHealthHandler(db, logger)

	// Setup Gin router
	router := setupRouter(cfg, benchmarkHandler, queryHandler, resultHandler, healthHandler)

	// Start server
	srv := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: router,
	}

	// Graceful server shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	logger.Info("Server started on port " + cfg.Server.Port)

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

func setupRouter(cfg *config.Config, benchmarkHandler *handlers.BenchmarkHandler, queryHandler *handlers.QueryHandler, resultHandler *handlers.ResultHandler, healthHandler *handlers.HealthHandler) *gin.Engine {
	if cfg.Server.Mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	
	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())
	router.Use(middleware.RequestID())
	router.Use(middleware.Metrics())
	
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

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Benchmark routes
		benchmarks := v1.Group("/benchmarks")
		{
			benchmarks.POST("", benchmarkHandler.CreateBenchmark)
			benchmarks.GET("", benchmarkHandler.ListBenchmarks)
			benchmarks.GET("/:id", benchmarkHandler.GetBenchmark)
			benchmarks.PUT("/:id", benchmarkHandler.UpdateBenchmark)
			benchmarks.DELETE("/:id", benchmarkHandler.DeleteBenchmark)
			benchmarks.POST("/:id/run", benchmarkHandler.RunBenchmark)
			benchmarks.GET("/:id/status", benchmarkHandler.GetBenchmarkStatus)
			benchmarks.GET("/:id/results", benchmarkHandler.GetBenchmarkResults)
		}

		// Query routes
		queries := v1.Group("/queries")
		{
			queries.POST("", queryHandler.CreateQuery)
			queries.GET("", queryHandler.ListQueries)
			queries.GET("/:id", queryHandler.GetQuery)
			queries.PUT("/:id", queryHandler.UpdateQuery)
			queries.DELETE("/:id", queryHandler.DeleteQuery)
			queries.POST("/:id/execute", queryHandler.ExecuteQuery)
			queries.GET("/:id/results", queryHandler.GetQueryResults)
		}

		// Result routes
		results := v1.Group("/results")
		{
			results.GET("", resultHandler.ListResults)
			results.GET("/:id", resultHandler.GetResult)
			results.GET("/compare", resultHandler.CompareResults)
			results.GET("/analytics", resultHandler.GetAnalytics)
		}

		// Engine routes
		engines := v1.Group("/engines")
		{
			engines.GET("", queryHandler.ListEngines)
			engines.GET("/:engine/status", queryHandler.GetEngineStatus)
		}

		// Table format routes
		tables := v1.Group("/tables")
		{
			tables.GET("/formats", queryHandler.ListTableFormats)
			tables.POST("/create", queryHandler.CreateTable)
			tables.GET("/:table/info", queryHandler.GetTableInfo)
		}
	}

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
