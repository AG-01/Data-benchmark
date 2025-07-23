package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// QueryExecutionDuration tracks query execution time
	QueryExecutionDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "query_execution_duration_seconds",
		Help: "Duration of query execution",
	}, []string{"engine", "query_type"})

	// QueryExecutionCount tracks number of queries executed
	QueryExecutionCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "query_execution_total",
		Help: "Total number of queries executed",
	}, []string{"engine", "status"})

	// EngineHealthStatus tracks engine health
	EngineHealthStatus = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "engine_health_status",
		Help: "Health status of query engines (1=healthy, 0=unhealthy)",
	}, []string{"engine"})

	// ActiveConnections tracks active database connections
	ActiveConnections = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "active_connections",
		Help: "Number of active database connections",
	}, []string{"engine", "active_connections"})
)

// Init initializes the metrics collectors.
func Init() {
	// This function can be used to register custom collectors or perform other setup.
	// For now, it's a placeholder to satisfy the call in main.go.
}

// RecordQueryExecution records metrics for a query execution
func RecordQueryExecution(engine, queryType string, duration float64, success bool) {
	status := "success"
	if !success {
		status = "error"
	}
	
	QueryExecutionDuration.WithLabelValues(engine, queryType).Observe(duration)
	QueryExecutionCount.WithLabelValues(engine, status).Inc()
}

// SetEngineHealth sets the health status for an engine
func SetEngineHealth(engine string, healthy bool) {
	value := float64(0)
	if healthy {
		value = 1
	}
	EngineHealthStatus.WithLabelValues(engine).Set(value)
}

// SetActiveConnections sets the number of active connections for an engine
func SetActiveConnections(engine string, count int) {
	ActiveConnections.WithLabelValues(engine).Set(float64(count))
}
