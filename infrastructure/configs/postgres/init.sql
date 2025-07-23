-- Initialize the Hive Metastore database
-- This script sets up the necessary tables and schema for Hive Metastore

\c hive_metastore;

-- Create extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create benchmark-specific tables for our application
CREATE TABLE IF NOT EXISTS benchmarks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    table_format VARCHAR(50) NOT NULL CHECK (table_format IN ('hive', 'iceberg')),
    dataset_name VARCHAR(255) NOT NULL,
    dataset_size VARCHAR(50) CHECK (dataset_size IN ('small', 'medium', 'large')),
    engines TEXT[], -- Array of engine names
    status VARCHAR(50) DEFAULT 'created' CHECK (status IN ('created', 'running', 'completed', 'failed')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS queries (
    id SERIAL PRIMARY KEY,
    benchmark_id INTEGER NOT NULL REFERENCES benchmarks(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    sql_query TEXT NOT NULL,
    query_type VARCHAR(50) CHECK (query_type IN ('select', 'aggregation', 'join', 'window')),
    complexity VARCHAR(50) CHECK (complexity IN ('simple', 'medium', 'complex')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS query_executions (
    id SERIAL PRIMARY KEY,
    query_id INTEGER NOT NULL REFERENCES queries(id) ON DELETE CASCADE,
    engine VARCHAR(100) NOT NULL,
    status VARCHAR(50) DEFAULT 'pending' CHECK (status IN ('pending', 'running', 'completed', 'failed')),
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    execution_time_ms BIGINT,
    rows_processed BIGINT,
    bytes_processed BIGINT,
    cpu_usage DECIMAL(5,2),
    memory_usage BIGINT,
    io_read_bytes BIGINT,
    io_write_bytes BIGINT,
    error_message TEXT,
    query_plan TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS results (
    id SERIAL PRIMARY KEY,
    benchmark_id INTEGER NOT NULL REFERENCES benchmarks(id) ON DELETE CASCADE,
    engine VARCHAR(100) NOT NULL,
    table_format VARCHAR(50) NOT NULL,
    total_queries INTEGER,
    successful_queries INTEGER,
    failed_queries INTEGER,
    avg_execution_time_ms DECIMAL(15,2),
    min_execution_time_ms DECIMAL(15,2),
    max_execution_time_ms DECIMAL(15,2),
    total_rows_processed BIGINT,
    total_bytes_processed BIGINT,
    avg_cpu_usage DECIMAL(5,2),
    avg_memory_usage DECIMAL(15,2),
    total_io_read_bytes BIGINT,
    total_io_write_bytes BIGINT,
    throughput DECIMAL(10,4), -- queries per second
    efficiency_score DECIMAL(5,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS table_info (
    id SERIAL PRIMARY KEY,
    table_name VARCHAR(255) UNIQUE NOT NULL,
    table_format VARCHAR(50) NOT NULL,
    schema_definition TEXT,
    location VARCHAR(500),
    partition_by TEXT[], -- Array of partition columns
    row_count BIGINT,
    size_bytes BIGINT,
    file_count INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS engines (
    name VARCHAR(100) PRIMARY KEY,
    type VARCHAR(50) NOT NULL CHECK (type IN ('trino', 'presto', 'starrocks')),
    version VARCHAR(50),
    host VARCHAR(255),
    port INTEGER,
    is_active BOOLEAN DEFAULT true,
    last_checked TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS datasets (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    size VARCHAR(50) CHECK (size IN ('small', 'medium', 'large')),
    row_count BIGINT,
    size_bytes BIGINT,
    format VARCHAR(50) CHECK (format IN ('parquet', 'orc', 'avro')),
    location VARCHAR(500),
    schema_definition TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better performance
CREATE INDEX idx_benchmarks_status ON benchmarks(status);
CREATE INDEX idx_benchmarks_table_format ON benchmarks(table_format);
CREATE INDEX idx_benchmarks_created_at ON benchmarks(created_at);

CREATE INDEX idx_queries_benchmark_id ON queries(benchmark_id);
CREATE INDEX idx_queries_query_type ON queries(query_type);

CREATE INDEX idx_query_executions_query_id ON query_executions(query_id);
CREATE INDEX idx_query_executions_engine ON query_executions(engine);
CREATE INDEX idx_query_executions_status ON query_executions(status);
CREATE INDEX idx_query_executions_start_time ON query_executions(start_time);

CREATE INDEX idx_results_benchmark_id ON results(benchmark_id);
CREATE INDEX idx_results_engine ON results(engine);
CREATE INDEX idx_results_table_format ON results(table_format);

CREATE INDEX idx_table_info_table_format ON table_info(table_format);
CREATE INDEX idx_engines_type ON engines(type);
CREATE INDEX idx_engines_is_active ON engines(is_active);

-- Create triggers for updating updated_at timestamps
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_benchmarks_updated_at BEFORE UPDATE ON benchmarks
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_queries_updated_at BEFORE UPDATE ON queries
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_query_executions_updated_at BEFORE UPDATE ON query_executions
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_results_updated_at BEFORE UPDATE ON results
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_table_info_updated_at BEFORE UPDATE ON table_info
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_engines_updated_at BEFORE UPDATE ON engines
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_datasets_updated_at BEFORE UPDATE ON datasets
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Insert sample engines
INSERT INTO engines (name, type, version, host, port) VALUES
('trino', 'trino', '432', 'trino', 8080),
('presto', 'presto', '0.284', 'presto', 8080),
('starrocks', 'starrocks', '3.2', 'starrocks-fe', 9030)
ON CONFLICT (name) DO NOTHING;

-- Insert sample datasets
INSERT INTO datasets (name, description, size, format, location) VALUES
('customer_small', 'Small customer dataset for quick testing', 'small', 'parquet', 's3a://benchmark-data/customer_small/'),
('customer_medium', 'Medium customer dataset for standard benchmarks', 'medium', 'parquet', 's3a://benchmark-data/customer_medium/'),
('customer_large', 'Large customer dataset for stress testing', 'large', 'parquet', 's3a://benchmark-data/customer_large/'),
('orders_small', 'Small orders dataset', 'small', 'parquet', 's3a://benchmark-data/orders_small/'),
('orders_medium', 'Medium orders dataset', 'medium', 'parquet', 's3a://benchmark-data/orders_medium/'),
('orders_large', 'Large orders dataset', 'large', 'parquet', 's3a://benchmark-data/orders_large/'),
('lineitem_small', 'Small lineitem dataset', 'small', 'parquet', 's3a://benchmark-data/lineitem_small/'),
('lineitem_medium', 'Medium lineitem dataset', 'medium', 'parquet', 's3a://benchmark-data/lineitem_medium/'),
('lineitem_large', 'Large lineitem dataset', 'large', 'parquet', 's3a://benchmark-data/lineitem_large/')
ON CONFLICT (name) DO NOTHING;

-- Grant necessary permissions
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO hive;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO hive;

-- Show created tables
\dt;
