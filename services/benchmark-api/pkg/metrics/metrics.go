package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// HTTP metrics
	HTTPRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "benchmark_http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)
	
	HTTPRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "benchmark_http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)
	
	// Benchmark metrics
	BenchmarkExecutionsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "benchmark_executions_total",
			Help: "Total number of benchmark executions",
		},
		[]string{"engine", "table_format", "status"},
	)
	
	QueryExecutionDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "benchmark_query_execution_duration_seconds",
			Help:    "Query execution duration in seconds",
			Buckets: []float64{0.1, 0.5, 1.0, 5.0, 10.0, 30.0, 60.0, 120.0, 300.0},
		},
		[]string{"engine", "table_format", "query_type"},
	)
	
	ActiveBenchmarks = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "benchmark_active_benchmarks",
			Help: "Number of currently active benchmarks",
		},
	)
)

func Init() {
	// Register custom metrics if needed
}

func RecordHTTPRequest(method, path, status string, duration float64) {
	HTTPRequestsTotal.WithLabelValues(method, path, status).Inc()
	HTTPRequestDuration.WithLabelValues(method, path).Observe(duration)
}

func RecordBenchmarkExecution(engine, tableFormat, status string) {
	BenchmarkExecutionsTotal.WithLabelValues(engine, tableFormat, status).Inc()
}

func RecordQueryExecution(engine, tableFormat, queryType string, duration float64) {
	QueryExecutionDuration.WithLabelValues(engine, tableFormat, queryType).Observe(duration)
}

func SetActiveBenchmarks(count float64) {
	ActiveBenchmarks.Set(count)
}
