# HanJin

[中文](README.md) | English

## Project Overview

`HanJin` is a production-grade Go Web backend framework built on Gin, adopting the standard Go project layout (cmd / pkg / internal). Pre-integrated with databases, message queues, caching, logging, middleware, and encryption utilities, it is well-suited for small-to-medium web backends, microservice modules, and business systems requiring fast iteration with multiple data source integration.

## Core Values

- **Ready to Use** — Pre-configured production-grade components for rapid business development
- **Standard Compliance** — Follows Go project standard layout and best practices
- **Highly Modular** — Clear layered architecture (Controller → Service → Database), easy to extend
- **Enterprise-Grade** — Comprehensive logging, encryption, signature verification, and containerized deployment

**Feature Overview**:

| Capability | Technology |
|------------|------------|
| Web Framework & API | Gin + RESTful + Swagger Documentation |
| Data Storage | MySQL (GORM), Elasticsearch, Redis / PostgreSQL / TDengine (extensible) |
| Message Queues | RabbitMQ (implemented), Kafka / RocketMQ (extensible) |
| Logging & Monitoring | Zap structured logging + Lumberjack rotation + trace_id tracking |
| Security & Encryption | AES / RSA + SM2 / SM3 / SM4 National Cryptography |
| Middleware | Panic Recovery, HMAC-SHA256 Signature Verification |
| Scheduled Tasks | robfig/cron with second-level precision |
| Deployment | Docker multi-stage build + Compose / Kubernetes + Nginx / HAProxy |

## Project Structure

```
x-HanJin/
├── cmd/                            # Application entry points
│   └── server/
│       └── main.go                 # Main entry: init config/logger/routes, start HTTP server
├── internal/                       # Private application code (not importable by external projects)
│   ├── config/                     # Configuration loading (viper)
│   ├── constants/                  # Application-level constants
│   ├── controllers/                # HTTP controllers (parameter binding, response formatting)
│   ├── databases/                  # Database initialization
│   │   ├── es/                     # Elasticsearch client
│   │   ├── kaiwudb/                # KaiwuDB (placeholder)
│   │   ├── mysql/                  # MySQL/GORM connection
│   │   ├── postgresql/             # PostgreSQL (placeholder)
│   │   ├── redis/                  # Redis (placeholder)
│   │   └── tdengine/               # TDengine (placeholder)
│   ├── event/                      # Event processing framework
│   ├── message/                    # Message processing framework
│   ├── message_queues/             # Message queue integration
│   │   ├── kafka/                  # Kafka (placeholder)
│   │   ├── rabbitmq/               # RabbitMQ producer/consumer
│   │   │   ├── consumer/
│   │   │   └── producer/
│   │   └── rocketmq/              # RocketMQ (placeholder)
│   ├── middlewares/                 # HTTP middleware
│   │   ├── exception_middleware.go # Panic recovery
│   │   └── signature_middleware.go # HMAC-SHA256 signature verification
│   ├── models/                     # Data models
│   │   └── user/
│   │       └── request/            # Request DTOs
│   ├── api/                        # API route registration
│   │   ├── router.go              # Aggregate versioned API routes
│   │   └── v1/
│   │       └── users.go           # User API routes
│   ├── services/                   # Business logic layer
│   └── tasks/                      # Scheduled tasks
├── pkg/                            # Reusable public packages
│   ├── log/                        # Zap logger (JSON output, rotation, remote push)
│   └── utils/                      # Utility functions
│       ├── aes_util.go             # AES encryption/decryption (CBC/ECB/GCM)
│       ├── coding_util.go          # Base64/Hex encoding/decoding
│       ├── ctx_util.go             # Context value helpers
│       ├── file_util.go            # File operations
│       ├── gen_util.go             # Random generation (salt, password, IV)
│       ├── http_util.go            # HTTP client (GET/POST/upload)
│       ├── int_util.go             # Integer ternary helpers
│       ├── json_util.go            # JSON serialization/deserialization
│       ├── rsa_util.go             # RSA encryption/decryption
│       ├── sm2_util.go             # SM2 national cryptography asymmetric encryption
│       ├── sm3_util.go             # SM3 national cryptography hash
│       ├── sm4_util.go             # SM4 national cryptography symmetric encryption
│       ├── str_util.go             # String utilities
│       └── time_util.go            # Time formatting/calculation
├── scripts/                        # Build/deploy scripts
│   └── run.sh                      # Docker startup script
├── configs/                        # Configuration files
│   ├── config.yaml                 # Application config (port/DB/MQ/logger)
│   ├── nginx.conf                  # Nginx reverse proxy config
│   └── haproxy.conf                # HAProxy load balancer config
├── deploy/                         # Deployment orchestration
│   ├── docker-compose/             # Docker Compose
│   └── kubernetes/                 # Kubernetes manifests
├── docs/                           # Swagger auto-generated docs
├── statics/                        # Static assets
├── .air.toml                       # Air hot reload configuration
├── .gitignore                      # Git ignore rules
├── Dockerfile                      # Multi-stage Docker build
├── LICENSE                         # MIT License
├── README.md                       # Chinese documentation
├── README.en.md                    # English documentation
├── go.mod                          # Go module definition
└── go.sum                          # Dependency checksums
```

## System Architecture

### System Layered Architecture

```mermaid
graph TD
    client["Client"]
    proxy["Nginx / HAProxy<br/>Reverse Proxy &amp; Load Balancer"]

    subgraph gin["Gin HTTP Server"]
        subgraph mw["Middleware Pipeline"]
            recovery["RecoveryMiddleware"]
            signature["SignatureMiddleware"]
            logger_mw["LoggerMiddleware"]
        end
        routes["Routes /api/v1/*"]
        controllers["Controllers<br/>Parameter Binding &amp; Response Formatting"]
        services["Services<br/>Business Logic"]
        db["Databases<br/>MySQL / ES / Redis / ..."]
        mq["Message Queues<br/>RabbitMQ / Kafka / RocketMQ"]
        tasks["Tasks<br/>Cron Scheduler"]
        pkg["pkg Public Packages<br/>log · utils encryption/HTTP/time/string"]
    end

    client -->|HTTP| proxy
    proxy --> mw
    mw --> routes --> controllers --> services
    services --> db
    services --> mq
    services --> tasks
```

### Core Function Business Flow

```mermaid
flowchart TD
    client["Client Request"] --> nginx["Nginx Reverse Proxy\n(80 → 8080)"]
    nginx --> gin["Gin Engine"]
    gin --> recovery["RecoveryMiddleware\nCapture panic, log stack"]
    recovery --> signature["SignatureMiddleware\nHMAC-SHA256 signature verification\n(Skip /swagger path)"]
    signature --> router["Router Dispatch\n(/api/v1/*)"]
    router --> controller["Controller\nParameter binding & validation"]
    controller --> service["Service Business Logic"]
    service --> mysql["MySQL (GORM)\nData persistence"]
    service --> rabbitmq["RabbitMQ Producer\nMessage notification"]
    service --> es["Elasticsearch\nSearch index"]
    service --> task["Task Scheduler\nScheduled tasks"]
    service --> response["Unified response format"]
```

### Module Dependency Diagram

```mermaid
flowchart TD
    cmd["cmd/server"]

    config["internal/config\nConfiguration (viper)"]
    databases["internal/databases\nDatabase initialization"]
    mysql["mysql"]
    es["es"]
    redis_pg["redis / postgresql / kaiwudb / tdengine\n(placeholder)"]

    middlewares["internal/middlewares\nHTTP middleware"]
    exception["exception"]
    signature["signature"]

    routes["internal/api\nAPI route registration"]
    controllers["controllers"]
    services["services"]

    event["internal/event\nEvent processing"]
    message["internal/message\nMessage processing"]
    message_queues["internal/message_queues\nMessage queues"]
    rabbitmq["rabbitmq"]

    tasks["internal/tasks\nScheduled tasks"]
    pkg_log["pkg/log"]
    pkg_utils["pkg/utils"]

    cmd --> config
    cmd --> databases
    cmd --> middlewares
    cmd --> routes
    cmd --> event
    cmd --> message
    cmd --> message_queues
    cmd --> tasks

    databases --> mysql
    databases --> es
    databases --> redis_pg

    mysql --> config
    mysql --> pkg_log
    es --> pkg_log

    middlewares --> exception
    middlewares --> signature
    exception --> pkg_log
    signature --> config

    routes --> controllers
    controllers --> services
    services --> mysql
    services --> rabbitmq
    services --> pkg_log

    event --> pkg_utils
    event --> pkg_log
    message --> pkg_log
    message_queues --> rabbitmq
    rabbitmq --> config
    rabbitmq --> pkg_log
    tasks --> pkg_log

    pkg_log --> config
    pkg_utils --> pkg_log
```

## Quick Start

### Environment Requirements

#### Windows

- Go 1.24+ (required; `go.mod` pinned to `go 1.24.0`)
- Git (for cloning the project)
- Configuration file: `configs/config.yaml` (needs to be filled with actual MySQL/Redis/ES/RabbitMQ addresses and credentials)
- Optional dependencies (enable as needed; services can be left unused if corresponding features are not used):
  - MySQL 5.7+ (data persistence)
  - Redis 6.0+ (cache/sessions)
  - Elasticsearch 8.x (search indexing)
  - RabbitMQ 3.8+ (async message queue)
- (Optional) Swagger documentation generation: install `swag` (`go install github.com/swaggo/swag/cmd/swag@latest`)
- (Optional) Hot reload development: install `air` (`go install github.com/cosmtrek/air@latest`)

#### Linux

- Go 1.24+ (required; `go.mod` pinned to `go 1.24.0`)
- Git (for cloning the project)
- Configuration file: `configs/config.yaml` (needs to be filled with actual MySQL/Redis/ES/RabbitMQ addresses and credentials)
- Optional dependencies (enable as needed; services can be left unused if corresponding features are not used):
  - MySQL 5.7+ (data persistence)
  - Redis 6.0+ (cache/sessions)
  - Elasticsearch 8.x (search indexing)
  - RabbitMQ 3.8+ (async message queue)
- (Optional) Swagger documentation generation: install `swag` (`go install github.com/swaggo/swag/cmd/swag@latest`)
- (Optional) Hot reload development: install `air` (`go install github.com/cosmtrek/air@latest`)

### Project Clone

```bash
git clone https://github.com/cross-lang/x-HanJin.git
cd x-HanJin
```

### Dependency Installation

```bash
go mod tidy
```

### Configuration File

Configuration file path: `configs/config.yaml`

```yaml
# Web service configuration
Web:
  host: localhost          # Service listen address
  port: 8080               # Service listen port

# MySQL database configuration
MySQL:
  host: localhost          # Database address
  port: 3306               # Database port
  user: root               # Username
  password: your_password  # Password (change to actual value)
  default_dbname: hanjin # Default database name

# Redis cache configuration
Redis:
  host: localhost          # Redis address
  port: 6379               # Redis port
  default_db: 0            # Default database number

# Elasticsearch configuration
ES:
  address: localhost       # ES address
  user: elastic            # Username
  password: your_password  # Password (change to actual value)

# RabbitMQ message queue configuration
RabbitMQ:
  host: localhost          # RabbitMQ address
  port: 5672               # RabbitMQ port
  user: guest              # Username
  password: guest          # Password
  default_queue_name: x-hanjin-queue  # Default queue name

# Application configuration
App:
  app_id: x-HanJin         # Application ID (for signature verification)
  app_key: your_app_key    # Application key (change to actual value)

# Logger configuration
Logger:
  LogDir: "./log"          # Log directory
  Level: "info"            # Log level: debug / info / warn / error
  EnableRemote: false      # Whether to enable remote log push
  RemoteURL: ""            # Remote log service address
```

| Configuration | Description |
|---------------|-------------|
| `Web.host` / `Web.port` | Web service listen address and port |
| `MySQL.*` | MySQL connection info (host/port/user/password/dbname) |
| `Redis.*` | Redis connection info (host/port/db) |
| `ES.*` | Elasticsearch connection info (address/user/password) |
| `RabbitMQ.*` | RabbitMQ connection info (host/port/user/password/queue) |
| `App.app_id` / `App.app_key` | Application identity and HMAC signing key |
| `Logger.*` | Logger configuration (directory/level/remote push) |

### Service Startup

#### Method 1: Local Development Mode (Hot Reload and Debug Mode Support)

```bash
# 1. Modify configuration (replace sensitive values like passwords with actual values)
# Edit configs/config.yaml

# 2. Install air hot reload tool (if not already installed)
go install github.com/air-verse/air@latest

# 3. Start service with air (supports hot reload)
air

# 4. Or start with debug mode
# Windows
set GIN_MODE=debug
go run ./cmd/server/

# Linux / macOS
export GIN_MODE=debug
go run ./cmd/server/

# 5. Build and start
go build -o server ./cmd/server/
./server
```

After service startup, access:
- API Service: `http://localhost:8080`
- Swagger Documentation: `http://localhost:8080/swagger/index.html`

#### Method 2: Docker Container Deployment

```bash
# Single container startup
docker build -t x-hanjin .
docker run -p 8080:8080 --name x-hanjin x-hanjin

# Docker Compose (includes MySQL + Redis + RabbitMQ)
cd deploy/docker-compose

# Basic services (without Elasticsearch)
docker-compose up -d

# Full services (includes Elasticsearch)
docker-compose --profile full up -d

# View service status
docker-compose ps

# View logs
docker-compose logs -f x-hanjin

# Stop services
docker-compose down

# Stop and delete data volumes
docker-compose down -v

# Kubernetes
kubectl apply -f deploy/kubernetes/test-gin.yaml
```

### Common Commands

```bash
# Build
go build -o server ./cmd/server/

# Run
go run ./cmd/server/

# Hot reload development
air

# Format code
go fmt ./...

# Static code check
go vet ./...

# Run tests
go test ./...

# Install dependencies
go mod tidy

# Generate Swagger documentation
swag init -g cmd/server/main.go -o docs/

# Docker build
docker build -t x-hanjin .

# Docker Compose startup
docker-compose -f deploy/docker-compose/docker-compose.yaml up -d
```

## Technology Stack

### Web Framework
- [Gin](https://github.com/gin-gonic/gin) - High-performance HTTP framework

### Data Storage
- [GORM](https://gorm.io) - Go language ORM library
- MySQL - Main data storage
- Elasticsearch - Search engine
- Redis - Cache and session storage
- PostgreSQL - Relational database (placeholder)

### Message Queues
- RabbitMQ - Async messaging
- Kafka - High-throughput message system (placeholder)
- RocketMQ - Distributed message system (placeholder)

### Tool Libraries
- [Viper](https://github.com/spf13/viper) - Configuration management, supports YAML/JSON/ENV multi-format
- [Zap](https://github.com/uber-go/zap) + [Lumberjack](https://github.com/natefinch/lumberjack) - Structured JSON logging + file rotation
- [robfig/cron](https://github.com/robfig/cron) - Scheduled tasks with second-level precision
- [Swaggo](https://github.com/swaggo/swag) - Swagger auto-generation

### Encryption and Security
- AES / RSA - Common encryption algorithms
- [SM2-SM4](https://github.com/tjfoc/gmsm) - National cryptography algorithms

### Deployment Tools
- Docker - Containerization
- Docker Compose - Multi-container orchestration
- Kubernetes - Container orchestration
- Nginx / HAProxy - Reverse proxy and load balancing

## API Documentation

The project integrates Swagger auto-generated API documentation with online interactive and offline export support.

- **Swagger UI Interactive Documentation**: http://localhost:8080/swagger/index.html
- **ReDoc Read-Only Documentation**: http://localhost:8080/swagger/doc.html
- **OpenAPI JSON Documentation**: http://localhost:8080/swagger/doc.json

Generate documentation command:

```bash
swag init -g cmd/server/main.go -o docs/
```

## Storage Configuration

### Local Storage

Local storage configuration is located in `configs/config.yaml`, mainly including:
- MySQL database connection configuration
- Redis cache configuration
- Elasticsearch search engine configuration

### Object Storage

Object storage functionality (such as Alibaba Cloud OSS, AWS S3, etc.) is currently a reserved module and can be extended and integrated according to business needs.

## License

This project is licensed under the [MIT License](LICENSE).

## References

- [Gin Framework Documentation](https://gin-gonic.com/docs/)
- [GORM Documentation](https://gorm.io/docs/)
- [Viper Configuration Management](https://github.com/spf13/viper)
- [Zap Logging Library](https://pkg.go.dev/go.uber.org/zap)
- [Go Project Standard Layout](https://github.com/golang-standards/project-layout)
- [Swaggo Swagger Generation](https://github.com/swaggo/swag)
- [Air Hot Reload Tool](https://github.com/cosmtrek/air)
- [Docker Official Documentation](https://docs.docker.com/)
- [Docker Compose Documentation](https://docs.docker.com/compose/)
- [Kubernetes Documentation](https://kubernetes.io/docs/)

## Contact

- **Author**: John Young (Online nickname: 夜雨诗来)
- **Email**: [john.young@foxmail.com](mailto:john.young@foxmail.com)
- **Gitee**: [https://gitee.com/yeyushilai](https://gitee.com/yeyushilai)
- **GitHub**: [https://github.com/yeyushilai](https://github.com/yeyushilai)
