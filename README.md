# Data Lake Benchmarking Platform

A high-performance benchmarking platform to compare Hive and Apache Iceberg table formats across multiple query engines (Trino, Presto) with real-time monitoring and visualization.

## 🏗️ Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   React Web UI  │────│   Go Services   │────│  Query Engines  │
│   (Frontend)    │    │   (APIs)        │    │ (Trino, Presto) │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         │              ┌─────────────────┐    ┌─────────────────┐
         │              │   PostgreSQL    │    │     MinIO       │
         └──────────────│  (Metastore)    │────│   (Storage)     │
                        └─────────────────┘    └─────────────────┘
```

## 🚀 Features

- **Multi-Engine Support**: Benchmark queries across Trino and Presto
- **Table Format Comparison**: Compare Hive and Apache Iceberg performance  
- **Real-time Monitoring**: Prometheus metrics with Grafana dashboards
- **Responsive Web UI**: Modern React interface for configuration and results
- **Local Development**: Complete Docker-based setup, no cloud dependencies
- **Microservices Architecture**: Clean Go services with proper API design

## 📁 Project Structure

```
data-benchmarks/
├── services/               # Go microservices
│   ├── benchmark-api/     # Main orchestration API
│   ├── query-service/     # Query execution service  
│   └── metrics-service/   # Metrics collection
├── web-ui/                # React frontend
├── infrastructure/        # Docker configs
└── docker-compose.yml    # Main orchestration
```

## 🛠️ Technology Stack

**Backend**: Go 1.21+, Gin, PostgreSQL, Prometheus  
**Frontend**: React 18, TypeScript, Material-UI  
**Infrastructure**: Docker Compose, Trino, Presto, MinIO, Grafana

## 🚀 Quick Start

### Prerequisites
- Docker & Docker Compose
- 8GB+ RAM recommended

### Launch Platform

```bash
git clone <repository>
cd data-benchmarks
docker-compose up -d
```

### Access Applications

- **Web UI**: http://localhost:3000 
- **API Documentation**: http://localhost:8080/swagger/index.html
- **Grafana**: http://localhost:3001 (admin/admin)
- **Prometheus**: http://localhost:9090
- **MinIO Console**: http://localhost:9001 (admin/password)

## 📊 Usage

1. **Access Web UI** at http://localhost:3000
2. **Configure Benchmarks** using the intuitive interface
3. **Execute Queries** across different engines and table formats
4. **Monitor Performance** in real-time via Grafana dashboards
5. **Analyze Results** to identify bottlenecks and optimize configurations

## 🔧 API Endpoints

### Core APIs
- `GET /api/v1/engines` - List available query engines
- `POST /api/v1/benchmarks` - Create new benchmark
- `POST /api/v1/benchmarks/{id}/run` - Execute benchmark  
- `GET /api/v1/results` - Retrieve benchmark results

Full API documentation: http://localhost:8080/swagger/index.html

## 📈 Monitoring

- **Prometheus**: Metrics collection at :9090
- **Grafana**: Visualization dashboards at :3001  
- **Health Checks**: Built-in endpoints for all services

## 🔍 Development

```bash
# Check service status
docker-compose ps

# View logs
docker-compose logs [service-name]

# Restart specific service
docker-compose restart [service-name]
```

## 🛡️ Health Status

All services include health check endpoints:
- `/health` - Service health status
- `/ready` - Readiness probe  
- `/metrics` - Prometheus metrics

## 📝 Notes

- PostgreSQL runs on port 5433 to avoid conflicts
- All services use Docker networking for internal communication
- Web UI proxies API calls through nginx for CORS handling

## 🚀 Features

- **Multi-Engine Support**: Benchmark queries across Trino, Presto, and StarRocks
- **Table Format Comparison**: Compare Hive and Apache Iceberg performance
- **Real-time Monitoring**: Prometheus metrics with Grafana dashboards
- **Responsive Web UI**: Modern React interface for configuration and results
- **Local Development**: Complete Docker-based setup, no cloud dependencies
- **Microservices Architecture**: Clean separation of concerns with Go services
- **Performance Analytics**: Detailed bottleneck identification and analysis

## 📁 Project Structure

```
data-benchmarks/
├── services/
│   ├── benchmark-api/          # Main API orchestrator
│   ├── query-service/          # Query execution service
│   └── metrics-service/        # Metrics collection service
├── web-ui/                     # React frontend
├── infrastructure/
│   ├── docker/                 # Service Dockerfiles
│   ├── configs/               # Configuration files
│   └── monitoring/            # Prometheus/Grafana configs
├── data/
│   ├── sample-datasets/       # Test datasets
│   └── schemas/              # Table schemas
├── scripts/                   # Utility scripts
├── docs/                      # Documentation
└── docker-compose.yml        # Main orchestration
```

## 🛠️ Technology Stack

### Backend Services
- **Language**: Go 1.21+
- **Framework**: Gin (HTTP), GORM (ORM)
- **Database**: PostgreSQL (Hive Metastore)
- **Monitoring**: Prometheus client
- **Containerization**: Docker

### Frontend
- **Framework**: React 18+ with TypeScript
- **UI Library**: Material-UI / Ant Design
- **State Management**: Redux Toolkit / Zustand
- **Charts**: Chart.js / Recharts
- **HTTP Client**: Axios

### Infrastructure
- **Query Engines**: Trino, Presto, StarRocks
- **Table Formats**: Apache Hive, Apache Iceberg
- **Object Storage**: MinIO (S3-compatible)
- **Metastore**: Hive Metastore
- **Monitoring**: Prometheus + Grafana
- **Orchestration**: Docker Compose

## 🚀 Quick Start

### Prerequisites
- Docker & Docker Compose
- Node.js 18+ (for frontend development)
- Go 1.21+ (for backend development)

### Running the Platform

1. **Clone and setup**:
   ```bash
   git clone <repository>
   cd data-benchmarks
   ```

2. **Start all services**:
   ```bash
   docker-compose up -d
   ```

3. **Access the applications**:
   - Web UI: http://localhost:3000
   - Benchmark API: http://localhost:8080
   - Grafana: http://localhost:3001 (admin/admin)
   - Prometheus: http://localhost:9090
   - MinIO Console: http://localhost:9001 (admin/password)

### Development Mode

1. **Start infrastructure only**:
   ```bash
   docker-compose up -d postgres minio hive-metastore prometheus grafana
   ```

2. **Run services locally**:
   ```bash
   # Backend services
   cd services/benchmark-api && go run main.go
   cd services/query-service && go run main.go
   cd services/metrics-service && go run main.go
   
   # Frontend
   cd web-ui && npm start
   ```

## 📊 Benchmarking Workflow

1. **Setup**: Configure table formats and datasets
2. **Execute**: Run benchmark queries across engines
3. **Monitor**: Real-time metrics collection
4. **Analyze**: Compare performance results
5. **Optimize**: Identify bottlenecks and tune configurations

## 🔧 Configuration

### Query Engines
- Trino: `infrastructure/configs/trino/`
- Presto: `infrastructure/configs/presto/`
- StarRocks: `infrastructure/configs/starrocks/`

### Table Formats
- Hive: Traditional partitioned tables
- Iceberg: Modern table format with time travel, schema evolution

### Monitoring
- Prometheus: `infrastructure/monitoring/prometheus/`
- Grafana: `infrastructure/monitoring/grafana/`

## 📈 Metrics & Analytics

### Performance Metrics
- Query execution time
- Resource utilization (CPU, Memory, I/O)
- Data scan efficiency
- Parallelism effectiveness

### Bottleneck Analysis
- Engine-specific performance characteristics
- Table format overhead comparison
- Storage I/O patterns
- Memory usage patterns

## 🧪 Testing

```bash
# Run all tests
make test

# Run specific service tests
make test-benchmark-api
make test-query-service
make test-metrics-service

# Run frontend tests
cd web-ui && npm test
```

## 📝 API Documentation

API documentation is available at:
- Swagger UI: http://localhost:8080/swagger/
- OpenAPI Spec: http://localhost:8080/api/v1/openapi.json

## 🤝 Contributing

1. Fork the repository
2. Create feature branch: `git checkout -b feature/amazing-feature`
3. Commit changes: `git commit -m 'Add amazing feature'`
4. Push to branch: `git push origin feature/amazing-feature`
5. Submit pull request

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Apache Iceberg community
- Trino/Presto projects
- StarRocks community
- MinIO team
- Prometheus & Grafana teams
