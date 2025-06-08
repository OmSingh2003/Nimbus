# Docker Setup for Simple Bank

## Prerequisites
- Docker
- Docker Compose

## Quick Start

### 1. Build and Run with Docker Compose
```bash
# Build and start all services
docker-compose up --build

# Run in background
docker-compose up -d --build
```

### 2. Access the Application
- **API Server**: http://localhost:8080
- **PostgreSQL Database**: localhost:5432

### 3. Test the API
```bash
# Create a user
curl -X POST "http://localhost:8080/users" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "Password123!",
    "full_name": "Test User",
    "email": "test@example.com"
  }'

# Login
curl -X POST "http://localhost:8080/users/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "Password123!"
  }'
```

### 4. Management Commands
```bash
# View logs
docker-compose logs -f simple_bank_app
docker-compose logs -f postgres

# Stop services
docker-compose down

# Stop and remove volumes (careful: this deletes data)
docker-compose down -v

# Rebuild only the app
docker-compose build simple_bank_app
docker-compose up simple_bank_app
```

## Services

### PostgreSQL Database
- **Image**: postgres:17
- **Port**: 5432
- **Database**: simple_bank
- **User**: root
- **Password**: secret
- **Volume**: Persistent data storage
- **Health Check**: Ensures database is ready before starting the app

### Simple Bank Application
- **Build**: From local Dockerfile
- **Port**: 8080
- **Environment**: Configured for Docker networking
- **Dependencies**: Waits for PostgreSQL to be healthy
- **Restart**: Automatically restarts on failure

## Environment Variables

The Docker setup uses `app.docker.env` with these configurations:
- `DB_SOURCE`: Points to the PostgreSQL container
- `SERVER_ADDRESS`: Binds to all interfaces
- `TOKEN_SYMMETRIC_KEY`: JWT token encryption key
- `ACCESS_TOKEN_DURATION`: Token expiration time

## Troubleshooting

### Database Connection Issues
```bash
# Check if PostgreSQL is running
docker-compose ps postgres

# Check PostgreSQL logs
docker-compose logs postgres

# Connect to database manually
docker exec -it simple_bank_postgres psql -U root -d simple_bank
```

### Application Issues
```bash
# Check application logs
docker-compose logs simple_bank_app

# Restart just the application
docker-compose restart simple_bank_app

# Enter application container
docker exec -it simple_bank_app sh
```

### Clean Start
```bash
# Stop everything and clean up
docker-compose down -v
docker system prune -f

# Start fresh
docker-compose up --build
```

