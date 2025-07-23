export interface Benchmark {
  id: number;
  name: string;
  description: string;
  table_format: 'hive' | 'iceberg';
  dataset_name: string;
  dataset_size: 'small' | 'medium' | 'large';
  engines: string[];
  status: 'created' | 'running' | 'completed' | 'failed';
  created_at: string;
  updated_at: string;
  queries?: Query[];
  results?: Result[];
}

export interface Query {
  id: number;
  benchmark_id: number;
  name: string;
  sql_query: string;
  query_type: 'select' | 'aggregation' | 'join' | 'window';
  complexity: 'simple' | 'medium' | 'complex';
  created_at: string;
  updated_at: string;
  executions?: QueryExecution[];
}

export interface QueryExecution {
  id: number;
  query_id: number;
  engine: string;
  status: 'pending' | 'running' | 'completed' | 'failed';
  start_time?: string;
  end_time?: string;
  execution_time_ms?: number;
  rows_processed?: number;
  bytes_processed?: number;
  cpu_usage?: number;
  memory_usage?: number;
  io_read_bytes?: number;
  io_write_bytes?: number;
  error_message?: string;
  query_plan?: string;
  created_at: string;
  updated_at: string;
}

export interface Result {
  id: number;
  benchmark_id: number;
  engine: string;
  table_format: string;
  total_queries: number;
  successful_queries: number;
  failed_queries: number;
  avg_execution_time_ms: number;
  min_execution_time_ms: number;
  max_execution_time_ms: number;
  total_rows_processed: number;
  total_bytes_processed: number;
  avg_cpu_usage: number;
  avg_memory_usage: number;
  total_io_read_bytes: number;
  total_io_write_bytes: number;
  throughput: number;
  efficiency_score: number;
  created_at: string;
  updated_at: string;
}

export interface Engine {
  name: string;
  type: 'trino' | 'presto' | 'starrocks';
  version: string;
  host: string;
  port: number;
  is_active: boolean;
  last_checked: string;
  created_at: string;
  updated_at: string;
}

export interface Dataset {
  id: number;
  name: string;
  description: string;
  size: 'small' | 'medium' | 'large';
  row_count: number;
  size_bytes: number;
  format: 'parquet' | 'orc' | 'avro';
  location: string;
  schema: string;
  created_at: string;
  updated_at: string;
}

export interface TableInfo {
  id: number;
  table_name: string;
  table_format: string;
  schema: string;
  location: string;
  partition_by: string[];
  row_count?: number;
  size_bytes?: number;
  file_count?: number;
  created_at: string;
  updated_at: string;
}

export interface BenchmarkStatus {
  benchmark_id: number;
  status: string;
  total_queries: number;
  completed_queries: number;
  failed_queries: number;
  progress_percentage: number;
  estimated_completion_time?: string;
  current_execution?: {
    query_name: string;
    engine: string;
    started_at: string;
  };
}

export interface ComparisonResult {
  engines: string[];
  table_formats: string[];
  metrics: {
    execution_time: ComparisonMetric;
    throughput: ComparisonMetric;
    resource_usage: ComparisonMetric;
    efficiency: ComparisonMetric;
  };
}

export interface ComparisonMetric {
  name: string;
  values: { [engine: string]: { [format: string]: number } };
  unit: string;
  description: string;
}
