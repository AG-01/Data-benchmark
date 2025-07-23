# Getting Started

This guide will help you set up and run the Data Lake Benchmarking Platform on your local machine.

## Prerequisites

Before you begin, ensure you have the following installed:

- **Docker** (20.10+) and **Docker Compose** (2.0+)
- **Git** for cloning the repository
- **Make** (optional, but recommended for easier project management)

### System Requirements

- **RAM**: Minimum 8GB, recommended 16GB+
- **Storage**: At least 10GB free space
- **CPU**: Multi-core processor recommended

## Quick Start

### 1. Clone the Repository

```bash
git clone <repository-url>
cd data-benchmarks
```

### 2. Run the Setup Script

The easiest way to get started is using our automated setup script:

```bash
chmod +x scripts/setup.sh
./scripts/setup.sh
```

This script will:
- Check system requirements
- Create necessary directories
- Set up environment variables
- Build Docker images
- Start all services
- Seed sample data

### 3. Alternative: Manual Setup

If you prefer to set up manually:

```bash
# Build and start all services
make build
make up

# Or using Docker Compose directly
docker-compose build
docker-compose up -d
```

### 4. Verify Installation

Once all services are running, you can access:

- **Web UI**: http://localhost:3000
- **API Documentation**: http://localhost:8080/swagger/
- **Grafana**: http://localhost:3001 (admin/admin)
- **Prometheus**: http://localhost:9090
- **MinIO Console**: http://localhost:9001 (admin/password)

## Application URLs

| Service | URL | Credentials |
|---------|-----|-------------|
| Web UI | http://localhost:3000 | - |
| Benchmark API | http://localhost:8080 | - |
| Swagger UI | http://localhost:8080/swagger/ | - |
| Grafana | http://localhost:3001 | admin/admin |
| Prometheus | http://localhost:9090 | - |
| MinIO Console | http://localhost:9001 | admin/password |
| Trino UI | http://localhost:8081 | - |
| Presto UI | http://localhost:8082 | - |
| StarRocks UI | http://localhost:8030 | - |

## Using the Platform

### Creating Your First Benchmark

1. Open the Web UI at http://localhost:3000
2. Navigate to "Benchmarks" section
3. Click "Create New Benchmark"
4. Configure your benchmark:
   - Choose table format (Hive or Iceberg)
   - Select dataset size
   - Choose query engines to test
   - Add SQL queries to benchmark

### Running Benchmarks

1. Go to your benchmark details page
2. Click "Run Benchmark"
3. Monitor progress in real-time
4. View results and performance metrics

### Analyzing Results

- Use the built-in analytics dashboard
- View Grafana dashboards for detailed metrics
- Export results for further analysis
- Compare performance across different configurations

## Development Mode

For development, you can run services locally while keeping infrastructure in Docker:

```bash
# Start only infrastructure services
make dev-infra

# Run services locally (in separate terminals)
make dev-api    # Starts Go API server
make dev-ui     # Starts React development server
```

### Prerequisites for Development

- **Go** 1.21+
- **Node.js** 18+
- **npm** or **yarn**

Install dependencies:
```bash
make install-deps
```

## Common Commands

```bash
# View service logs
make logs

# View specific service logs
make logs-api
make logs-ui

# Stop all services
make down

# Stop and clean up everything
make down-clean

# Run tests
make test

# Check service status
make status

# Open all dashboards
make dashboard
```

## Troubleshooting

### Port Conflicts

If you encounter port conflicts, check what's running on the required ports:

```bash
# Check if ports are in use
lsof -i :3000,8080,8081,8082,9000,9001,9090,3001,5432
```

Stop conflicting services or modify the ports in `docker-compose.yml`.

### Services Not Starting

1. Check Docker logs:
   ```bash
   docker-compose logs [service-name]
   ```

2. Verify Docker resources:
   ```bash
   docker system df
   docker system prune  # Clean up if needed
   ```

3. Check available memory:
   ```bash
   free -h
   ```

### Database Connection Issues

If you see database connection errors:

1. Wait for PostgreSQL to fully initialize (can take 30-60 seconds)
2. Check if the database container is healthy:
   ```bash
   docker-compose ps postgres
   ```

### Query Engine Connection Issues

1. Verify engines are running:
   ```bash
   docker-compose ps trino presto starrocks-fe
   ```

2. Check engine logs:
   ```bash
   docker-compose logs trino
   docker-compose logs presto
   docker-compose logs starrocks-fe
   ```

### Memory Issues

If services are running out of memory:

1. Increase Docker's memory limit
2. Close unnecessary applications
3. Consider running fewer services simultaneously

## Configuration

### Environment Variables

Key environment variables are set in `.env` file:

```bash
# Database
DB_HOST=postgres
DB_PASSWORD=hivepass

# MinIO
MINIO_ACCESS_KEY=admin
MINIO_SECRET_KEY=password

# Engines
TRINO_HOST=trino:8080
PRESTO_HOST=presto:8080
```

### Custom Datasets

To add your own datasets:

1. Upload data to MinIO via the console (http://localhost:9001)
2. Create table definitions in the Web UI
3. Add queries for benchmarking

### Custom Queries

Add your SQL queries in the Query Editor or via the API:

```bash
curl -X POST http://localhost:8080/api/v1/queries \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Custom Query",
    "sql_query": "SELECT * FROM my_table LIMIT 100",
    "query_type": "select"
  }'
```

## Performance Tips

1. **Resource Allocation**: Ensure Docker has sufficient memory (8GB+)
2. **Storage**: Use SSD storage for better I/O performance
3. **Network**: Use wired connection for stable networking
4. **Monitoring**: Keep an eye on resource usage via Grafana

## Getting Help

- Check the logs: `make logs`
- Review the API documentation: http://localhost:8080/swagger/
- Monitor system resources via Grafana
- Check the troubleshooting section above

## Next Steps

Once you have the platform running:

1. Explore the sample datasets and queries
2. Create your own benchmarks
3. Compare Hive vs Iceberg performance
4. Analyze query execution patterns
5. Tune query engine configurations
6. Export results for presentations

Happy benchmarking! 🚀
