# simple-bank

[![Go Version](https://img.shields.io/badge/Go-1.24+-blue.svg)](https://golang.org/)
[![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen)](https://github.com/YOUR_USERNAME/simple-bank/actions) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) Full-featured backend system for a simple bank built with Go. Implements REST/gRPC APIs (Gin/gRPC Gateway), PostgreSQL DB access (SQLC & migrations), JWT/PASETO authentication, asynchronous task processing (Redis/Asynq), Docker deployment strategies, and CI/CD pipelines.

**Status:** Currently under active development.

## ✨ Features

* **RESTful API:** Built with [Gin](https://github.com/gin-gonic/gin) for core banking operations (accounts, transfers, entries).
* **gRPC API:** High-performance RPC framework for inter-service communication and client SDK generation.
* **gRPC Gateway:** Translates RESTful JSON API requests into gRPC messages, allowing one API definition for both. Includes embedded Swagger UI.
* **Database:** PostgreSQL database integration.
* **Type-Safe SQL:** Uses [SQLC](https://github.com/sqlc-dev/sqlc) to generate type-safe Go code from SQL queries.
* **DB Migrations:** Manages database schema changes using [golang-migrate](https://github.com/golang-migrate/migrate).
* **Authentication:** Secure user authentication using both JWT and PASETO tokens.
* **Asynchronous Tasks:** Background task processing (e.g., sending verification emails) using Redis and [Asynq](https://github.com/hibiken/asynq).
* **Dockerized:** Fully containerized application using Docker and Docker Compose for consistent environments.
* **Configuration:** Centralized configuration management using [Viper](https://github.com/spf13/viper).
* **Logging:** Structured logging using [Zerolog](https://github.com/rs/zerolog).
* **CI/CD:** Automated testing and build pipelines (e.g., using GitHub Actions - *placeholder*).

## 🛠️ Tech Stack

* **Language:** [Go](https://golang.org/) (version 1.19+)
* **API Frameworks:** [Gin](https://github.com/gin-gonic/gin), [gRPC](https://grpc.io/), [gRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway)
* **Database:** [PostgreSQL](https://www.postgresql.org/)
* **ORM/SQL Builder:** [SQLC](https://github.com/sqlc-dev/sqlc)
* **Migrations:** [golang-migrate](https://github.com/golang-migrate/migrate)
* **Async Tasks:** [Redis](https://redis.io/), [Asynq](https://github.com/hibiken/asynq)
* **Authentication:** JWT ([golang-jwt](https://github.com/golang-jwt/jwt)), PASETO ([paseto](https://github.com/o1egl/paseto))
* **Configuration:** [Viper](https://github.com/spf13/viper)
* **Logging:** [Zerolog](https://github.com/rs/zerolog)
* **Containerization:** [Docker](https://www.docker.com/), [Docker Compose](https://docs.docker.com/compose/)
* **Build/Task Runner:** [Make](https://www.gnu.org/software/make/)
* **DB Documentation:** [DBML](https://dbml.org/), [dbdocs](https://dbdocs.io/)
* **Mocking:** [Gomock](https://github.com/golang/mock)

## 📋 Prerequisites & Tool Installation

Ensure you have the following core tools installed:

* [Go](https://golang.org/doc/install) (version 1.19 or higher)
* [Docker Desktop](https://docs.docker.com/get-docker/) (includes Docker & Docker Compose)
* [Make](https://www.gnu.org/software/make/)
* Optional DB GUI: [TablePlus](https://tableplus.com/) or similar

Install project-specific CLI tools:

* **Homebrew (macOS):** If using macOS, [install Homebrew](https://brew.sh/) first.
* **Migrate CLI:** Manages database migrations.
    ```bash
    # Using Homebrew (macOS)
    brew install golang-migrate
    # Or build from source (other OS) - see golang-migrate docs
    ```
* **SQLC CLI:** Generates Go code from SQL.
    ```bash
    # Using Homebrew (macOS)
    brew install sqlc
    # Or build from source (other OS) - see sqlc docs
    ```
* **buf CLI:** Manages Protobuf files.
    ```bash
    # Using Homebrew (macOS)
    brew install bufbuild/buf/buf
    # Or see buf installation docs
    ```
* **Gomock:** Generates mock code for testing.
    ```bash
    go install [github.com/golang/mock/mockgen@v1.6.0](https://github.com/golang/mock/mockgen@v1.6.0)
    ```
* **DBML CLI:** Converts DBML to SQL (for schema generation).
    ```bash
    npm install -g @dbml/cli
    # Verify installation
    dbml2sql --version
    ```
* **dbdocs CLI:** Generates database documentation website.
    ```bash
    npm install -g dbdocs
    # Login to dbdocs service (required for hosting)
    dbdocs login
    ```

## 🚀 Local Development Setup

Follow these steps to set up the project for local development:

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/OmSingh2003/simple-bank.git](https://github.com/OmSingh2003/simple-bank.git)
    cd simple-bank
    ```

2.  **Set up Environment Variables:**
    Copy the example environment file and update it with your specific configuration.
    ```bash
    cp .env.example .env
    ```
    * Modify `.env` with your database credentials (use defaults if using `make postgres`), token symmetric keys, Redis address, etc.

3.  **Setup Infrastructure (Docker Network & Postgres):**
    This creates a dedicated Docker network and starts a PostgreSQL container.
    ```bash
    # Create the docker network (if it doesn't exist)
    make network

    # Start postgres container
    make postgres
    ```

4.  **Create Database:**
    Connects to the running Postgres container and creates the `simple_bank` database.
    ```bash
    make createdb
    ```

5.  **Run Initial Database Migration:**
    Applies all available migrations to set up the database schema.
    ```bash
    make migrateup
    ```

6.  **Install Go Dependencies:**
    ```bash
    go mod tidy
    ```

7.  **Generate Code (SQLC, Mocks, Protobuf):**
    This step generates necessary Go code based on SQL queries, Protobuf definitions, and creates mocks for testing.
    ```bash
    # Generate Go code from SQL queries in db/query/
    make sqlc

    # Generate gRPC, gRPC-gateway, and Swagger code from proto files
    make proto

    # Generate mock code for testing interfaces
    make mock
    ```

8.  **Build the Application:**
    Compiles the Go application binaries.
    ```bash
    make build
    ```

## ▶️ Running the Application

**Using Docker Compose (Recommended):**

This starts all services (PostgreSQL, Redis, API server, gRPC server, async worker) defined in `docker-compose.yml`.

```bash
docker-compose up --build
(The --build flag ensures images are rebuilt if code changes)Using Make (Requires Manual Service Management):If you prefer not to use Docker Compose for the Go services (but still use make postgres for the DB):# Start the main API server (REST & gRPC Gateway)
make server

# (In another terminal) Start the asynchronous task worker
make worker
The REST API server typically runs on :8080.The gRPC server typically runs on :9090.The Asynq web UI (if enabled) runs on :8081.🧪 Running TestsTo run the test suite:make test
(This usually runs go test -v -cover ./...)🔄 Migrations ManagementUse Make commands to manage database schema migrations:Apply all pending migrations:make migrateup
Apply the next pending migration:make migrateup1
Roll back all migrations:make migratedown
Roll back the last applied migration:make migratedown1
Create a new migration file:Replace <migration_name> with a descriptive name (e.g., add_users_table).make new_migration name=<migration_name>
📄 Database DocumentationGenerate and view database documentation using DBML and dbdocs:Generate DBML schema file (if needed):(This might be manual or part of another process depending on your setup)Generate schema SQL file from DBML:(Useful for visualizing or comparing)make db_schema
Generate and publish documentation website:(Requires prior dbdocs login)make db_docs
Access the DB documentation at the URL provided by the command output. (Password: secret - as noted in your input, consider if this should be documented or secured differently)📄 API Documentation (Swagger)API documentation is automatically generated from the Protobuf definitions and served via Swagger UI.Once the server is running (using docker-compose up or make server), access the Swagger UI at:http://localhost:8080/swagger/☁️ Deployment (Kubernetes Example)These are example steps for setting up prerequisites in a Kubernetes cluster for deployment:Install Nginx Ingress Controller:(Example for AWS, check provider docs for others)kubectl apply -f [https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v0.48.1/deploy/static/provider/aws/deploy.yaml](https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v0.48.1/deploy/static/provider/aws/deploy.yaml)
Install Cert-Manager:(For automatic TLS certificate management)kubectl apply -f [https://github.com/jetstack/cert-manager/releases/download/v1.4.0/cert-manager.yaml](https://github.com/jetstack/cert-manager/releases/download/v1.4.0/cert-manager.yaml)
(Note: Ensure you use versions compatible with your cluster. These are examples.)(Add specific deployment steps for the simple-bank application itself here, e.g., applying Kubernetes manifests for Deployments, Services, Ingress, Secrets, etc.)🏗️ Project Structure (Overview).
├── api         # Gin handlers, middleware, server setup
├── cmd         # Main application entry points (server, worker)
├── db          # Database migrations and SQLC queries/schema
├── gapi        # gRPC handlers, interceptors, server setup
├── internal    # Core business logic, domain types (shared internal code)
├── mail        # Mail sending logic
├── pb          # Generated Protobuf Go code
├── proto       # Protobuf definition files
├── token       # JWT/PASETO token generation and verification logic
├── util        # Utility functions (config, logging, etc.)
├── worker      # Asynq task definitions and processor setup
├── .env.example # Example environment variables
├── Dockerfile  # Docker build instructions
├── docker-compose.yml # Docker Compose service definitions
├── go.mod      # Go module dependencies
├── Makefile    # Make targets for common tasks
└── main.go     # Main application entry point (often calls cmd)
(Adjust this structure based on your actual project layout)🔄 CI/CDThis project uses GitHub Actions for continuous integration. The workflow includes:Running linters (golangci-lint).Running unit tests.Building the application.(Describe your specific CI/CD setup here)🤝 ContributingContributions are welcome! Please follow standard Go practices and ensure tests pass before submitting a pull request.(Add more detailed contribution guidelines if needed)📜 LicenseThis project is licensed under the MIT License - see the LICENSE file for details.
