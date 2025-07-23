#!/bin/bash

# Data Lake Benchmark Platform Setup Script
set -e

echo "🚀 Setting up Data Lake Benchmark Platform..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if Docker is installed
check_docker() {
    if ! command -v docker &> /dev/null; then
        print_error "Docker is not installed. Please install Docker first."
        exit 1
    fi
    
    if ! command -v docker-compose &> /dev/null; then
        print_error "Docker Compose is not installed. Please install Docker Compose first."
        exit 1
    fi
    
    print_success "Docker and Docker Compose are installed"
}

# Check if required ports are available
check_ports() {
    local ports=(3000 8080 8081 8082 8030 9000 9001 9090 3001 5432)
    local busy_ports=()
    
    for port in "${ports[@]}"; do
        if lsof -Pi :$port -sTCP:LISTEN -t >/dev/null 2>&1; then
            busy_ports+=($port)
        fi
    done
    
    if [ ${#busy_ports[@]} -ne 0 ]; then
        print_warning "The following ports are already in use: ${busy_ports[*]}"
        print_warning "Please stop the services using these ports or they will conflict"
        read -p "Continue anyway? (y/N): " -n 1 -r
        echo
        if [[ ! $REPLY =~ ^[Yy]$ ]]; then
            exit 1
        fi
    else
        print_success "All required ports are available"
    fi
}

# Create necessary directories
create_directories() {
    print_status "Creating necessary directories..."
    
    # Data directories
    mkdir -p data/minio
    mkdir -p data/postgres
    mkdir -p data/prometheus
    mkdir -p data/grafana
    
    # Log directories
    mkdir -p logs/api
    mkdir -p logs/query-service
    mkdir -p logs/metrics-service
    
    print_success "Directories created"
}

# Set environment variables
setup_environment() {
    print_status "Setting up environment variables..."
    
    if [ ! -f .env ]; then
        cat > .env << EOF
# Database Configuration
DB_HOST=postgres
DB_PORT=5432
DB_USER=hive
DB_PASSWORD=hivepass
DB_NAME=hive_metastore

# MinIO Configuration
MINIO_ENDPOINT=minio:9000
MINIO_ACCESS_KEY=admin
MINIO_SECRET_KEY=password

# Engine Configuration
TRINO_HOST=trino:8080
PRESTO_HOST=presto:8080

# Monitoring Configuration
PROMETHEUS_URL=http://prometheus:9090

# Application Configuration
GIN_MODE=debug
EOF
        print_success "Environment file created"
    else
        print_warning ".env file already exists, skipping creation"
    fi
}

# Build images
build_images() {
    print_status "Building Docker images..."
    
    if ! docker-compose build; then
        print_error "Failed to build Docker images"
        exit 1
    fi
    
    print_success "Docker images built successfully"
}

# Start services
start_services() {
    print_status "Starting services..."
    
    # Start infrastructure services first
    print_status "Starting infrastructure services..."
    docker-compose up -d postgres minio
    
    # Wait for database to be ready
    print_status "Waiting for database to be ready..."
    sleep 10
    
    # Start metastore
    print_status "Starting Hive Metastore..."
    docker-compose up -d hive-metastore
    sleep 15
    
    # Start query engines
    print_status "Starting query engines..."
    docker-compose up -d trino presto
    sleep 10
    
    # Start monitoring
    print_status "Starting monitoring services..."
    docker-compose up -d prometheus grafana
    
    # Start application services
    print_status "Starting application services..."
    docker-compose up -d benchmark-api query-service metrics-service web-ui
    
    print_success "All services started!"
}

# Show service status
show_status() {
    print_status "Checking service status..."
    
    echo ""
    echo "Service Status:"
    docker-compose ps
    
    echo ""
    echo "🌐 Application URLs:"
    echo "  Web UI:           http://localhost:3000"
    echo "  Benchmark API:    http://localhost:8080"
    echo "  API Documentation: http://localhost:8080/swagger/"
    echo "  Grafana:          http://localhost:3001 (admin/admin)"
    echo "  Prometheus:       http://localhost:9090"
    echo "  MinIO Console:    http://localhost:9001 (admin/password)"
    echo "  Trino UI:         http://localhost:8081"
    echo "  Presto UI:        http://localhost:8082"
    echo ""
}

# Wait for services to be healthy
wait_for_services() {
    print_status "Waiting for services to be healthy..."
    
    local max_attempts=30
    local attempt=0
    
    while [ $attempt -lt $max_attempts ]; do
        if curl -s http://localhost:8080/health > /dev/null 2>&1; then
            print_success "Benchmark API is healthy"
            break
        fi
        
        attempt=$((attempt + 1))
        echo -n "."
        sleep 2
    done
    
    if [ $attempt -eq $max_attempts ]; then
        print_warning "Benchmark API health check timeout, but continuing..."
    fi
    
    echo ""
}

# Seed sample data
seed_data() {
    print_status "Seeding sample data..."
    
    # Wait a bit more for all services to be ready
    sleep 5
    
    # Create sample benchmark data via API
    curl -s -X POST http://localhost:8080/api/v1/benchmarks \
        -H "Content-Type: application/json" \
        -d '{
            "name": "Sample Hive vs Iceberg Benchmark",
            "description": "Compare performance between Hive and Iceberg table formats",
            "table_format": "hive",
            "dataset_name": "sample_dataset",
            "dataset_size": "small",
            "engines": ["trino", "presto"]
        }' > /dev/null 2>&1 || print_warning "Failed to create sample benchmark"
    
    print_success "Sample data seeded"
}

# Main execution
main() {
    echo "============================================="
    echo "  Data Lake Benchmark Platform Setup"
    echo "============================================="
    echo ""
    
    check_docker
    check_ports
    create_directories
    setup_environment
    build_images
    start_services
    wait_for_services
    seed_data
    show_status
    
    echo ""
    print_success "🎉 Setup completed successfully!"
    echo ""
    echo "To stop all services, run: docker-compose down"
    echo "To view logs, run: docker-compose logs -f"
    echo "To rebuild and restart, run: make build && make up"
    echo ""
    echo "Happy benchmarking! 🚀"
}

# Handle script interruption
trap 'echo ""; print_warning "Setup interrupted"; exit 1' INT

# Run main function
main "$@"
