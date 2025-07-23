package models

import (
	"time"
	"gorm.io/gorm"
)

// Benchmark represents a benchmark configuration
type Benchmark struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	TableFormat string    `json:"table_format" gorm:"not null"` // "hive" or "iceberg"
	DatasetName string    `json:"dataset_name" gorm:"not null"`
	DatasetSize string    `json:"dataset_size"`                 // "small", "medium", "large"
	Engines     []string  `json:"engines" gorm:"type:text[]"`   // JSON array of engine names
	Status      string    `json:"status" gorm:"default:'created'"` // "created", "running", "completed", "failed"
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// Relationships
	Queries []Query  `json:"queries,omitempty" gorm:"foreignKey:BenchmarkID"`
	Results []Result `json:"results,omitempty" gorm:"foreignKey:BenchmarkID"`
}

// Query represents a SQL query to be benchmarked
type Query struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	BenchmarkID uint   `json:"benchmark_id" gorm:"not null"`
	Name        string `json:"name" gorm:"not null"`
	SQLQuery    string `json:"sql_query" gorm:"type:text;not null"`
	QueryType   string `json:"query_type"` // "select", "aggregation", "join", "window"
	Complexity  string `json:"complexity"` // "simple", "medium", "complex"
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// Relationships
	Benchmark   Benchmark `json:"benchmark,omitempty" gorm:"foreignKey:BenchmarkID"`
	Executions  []QueryExecution `json:"executions,omitempty" gorm:"foreignKey:QueryID"`
}

// QueryExecution represents a single execution of a query on a specific engine
type QueryExecution struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	QueryID          uint      `json:"query_id" gorm:"not null"`
	Engine           string    `json:"engine" gorm:"not null"`
	Status           string    `json:"status" gorm:"default:'pending'"` // "pending", "running", "completed", "failed"
	StartTime        *time.Time `json:"start_time"`
	EndTime          *time.Time `json:"end_time"`
	ExecutionTimeMs  *int64     `json:"execution_time_ms"`
	RowsProcessed    *int64     `json:"rows_processed"`
	BytesProcessed   *int64     `json:"bytes_processed"`
	CPUUsage         *float64   `json:"cpu_usage"`
	MemoryUsage      *int64     `json:"memory_usage"`
	IOReadBytes      *int64     `json:"io_read_bytes"`
	IOWriteBytes     *int64     `json:"io_write_bytes"`
	ErrorMessage     *string    `json:"error_message"`
	QueryPlan        *string    `json:"query_plan" gorm:"type:text"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	
	// Relationships
	Query  Query  `json:"query,omitempty" gorm:"foreignKey:QueryID"`
}

// Result represents aggregated benchmark results
type Result struct {
	ID                    uint      `json:"id" gorm:"primaryKey"`
	BenchmarkID           uint      `json:"benchmark_id" gorm:"not null"`
	Engine                string    `json:"engine" gorm:"not null"`
	TableFormat           string    `json:"table_format" gorm:"not null"`
	TotalQueries          int       `json:"total_queries"`
	SuccessfulQueries     int       `json:"successful_queries"`
	FailedQueries         int       `json:"failed_queries"`
	AvgExecutionTimeMs    float64   `json:"avg_execution_time_ms"`
	MinExecutionTimeMs    float64   `json:"min_execution_time_ms"`
	MaxExecutionTimeMs    float64   `json:"max_execution_time_ms"`
	TotalRowsProcessed    int64     `json:"total_rows_processed"`
	TotalBytesProcessed   int64     `json:"total_bytes_processed"`
	AvgCPUUsage           float64   `json:"avg_cpu_usage"`
	AvgMemoryUsage        float64   `json:"avg_memory_usage"`
	TotalIOReadBytes      int64     `json:"total_io_read_bytes"`
	TotalIOWriteBytes     int64     `json:"total_io_write_bytes"`
	Throughput            float64   `json:"throughput"` // queries per second
	EfficiencyScore       float64   `json:"efficiency_score"` // custom metric
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	
	// Relationships
	Benchmark Benchmark `json:"benchmark,omitempty" gorm:"foreignKey:BenchmarkID"`
}

// TableInfo represents metadata about a table
type TableInfo struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	TableName    string    `json:"table_name" gorm:"uniqueIndex;not null"`
	TableFormat  string    `json:"table_format" gorm:"not null"`
	Schema       string    `json:"schema" gorm:"type:text"`
	Location     string    `json:"location"`
	PartitionBy  []string  `json:"partition_by" gorm:"type:text[]"`
	RowCount     *int64    `json:"row_count"`
	SizeBytes    *int64    `json:"size_bytes"`
	FileCount    *int      `json:"file_count"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Engine represents a query engine configuration
type Engine struct {
	Name        string    `json:"name" gorm:"primaryKey"`
	Type        string    `json:"type"` // "trino", "presto", "starrocks"
	Version     string    `json:"version"`
	Host        string    `json:"host"`
	Port        int       `json:"port"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	LastChecked time.Time `json:"last_checked"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Dataset represents a test dataset
type Dataset struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"uniqueIndex;not null"`
	Description string    `json:"description"`
	Size        string    `json:"size"` // "small", "medium", "large"
	RowCount    int64     `json:"row_count"`
	SizeBytes   int64     `json:"size_bytes"`
	Format      string    `json:"format"` // "parquet", "orc", "avro"
	Location    string    `json:"location"`
	Schema      string    `json:"schema" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
