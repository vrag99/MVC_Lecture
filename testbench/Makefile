.PHONY: build up down clean logs test-express test-go help

# Default target
help:
	@echo "Available commands:"
	@echo "  build         - Build all Docker images"
	@echo "  up            - Start all services"
	@echo "  down          - Stop all services"
	@echo "  clean         - Remove all containers, images, and volumes"
	@echo "  logs          - Show logs for all services"
	@echo "  logs-express  - Show logs for Express server"
	@echo "  logs-go       - Show logs for Go server"
	@echo "  logs-mysql    - Show logs for MySQL"
	@echo "  test-express  - Test Express server endpoint"
	@echo "  test-go       - Test Go server endpoint"
	@echo "  restart       - Restart all services"

# Build all images
build:
	docker-compose build

# Start all services
up:
	docker-compose up -d

# Stop all services
down:
	docker-compose down

# Clean everything
clean:
	docker-compose down -v --rmi all
	docker system prune -f

# Show logs
logs:
	docker-compose logs -f

logs-express:
	docker-compose logs -f express-server

logs-go:
	docker-compose logs -f go-server

logs-mysql:
	docker-compose logs -f mysql

# Test endpoints
test-express:
	@echo "Testing Express server..."
	@curl -X POST http://localhost:3000/process \
		-H "Content-Type: application/json" \
		-d '{"user_id": 1, "data": "test_data_for_performance_testing"}' \
		| jq '.'

test-go:
	@echo "Testing Go server..."
	@curl -X POST http://localhost:8080/process \
		-H "Content-Type: application/json" \
		-d '{"user_id": 1, "data": "test_data_for_performance_testing"}' \
		| jq '.'

bench-go:
	ab -n 5000 -c 100 -p test-payload.json -T application/json http://localhost:8080/process

bench-express:
	ab -n 5000 -c 100 -p test-payload.json -T application/json http://localhost:3000/process

# Health checks
health:
	@echo "Express server health:"
	@curl -s http://localhost:3000/health | jq '.'
	@echo "Go server health:"
	@curl -s http://localhost:8080/health | jq '.'

# Restart services
restart:
	docker-compose restart
