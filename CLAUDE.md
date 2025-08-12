# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

flow-run is an open-source runner for LLM workflows that makes prompts testable, deployments repeatable, and costs observable. This is a Go-based project using Go 1.24.4.

### Project Vision

flow-run treats LLM integration as an infrastructure component rather than an application component, similar to how we handle databases. It addresses common challenges in AI Engineering:

- **Reliability**: LLM providers lack reliability - this project provides a reliable execution layer
- **Coupling**: Modern approaches tightly couple prompts with application code - flow-run decouples prompt development from application development
- **Language restrictions**: AI frameworks are limited to specific languages - flow-run is language-agnostic
- **Testing complexity**: Prompts need evaluation tests to ensure new versions perform better than previous ones

### Core Concepts

- **Infrastructure as Code approach**: Prompts are treated as code, similar to Terraform
- **Decoupled execution**: Prompt execution is completely separated from application execution
- **Fire-and-forget semantics**: Reliable execution with guaranteed processing
- **Multi-provider support**: Easy switching between LLM providers without code changes

### Key Features (Roadmap)

**v1 - MVP:**
- Prompt development without traditional programming languages
- Workflow support for sequential prompt execution
- Easy deployment to dev/prod environments
- Multi-LLM provider support
- Application integration for executing AI flows

**v2 - Quality & Observability:**
- Prompt testing and evaluation capabilities
- Observability and cost monitoring
- CI/CD integration for automated testing and deployment

**v3 - Advanced Features:**
- AI Agent support within workflows

### User Personas

#### User Persona 1: Prompt Developer
**Role:** Develop, debug, evaluate, and deploy prompts  
**Background:** Former software engineer with knowledge of building software products and using development tools  
**Primary Pain Point:** No unified approach to prompt development; constantly writing Python scripts for quick testing, implementing workarounds for prompt evaluation

#### User Persona 2: Application Developer
**Role:** Develop, debug, and deploy application servers  
**Background:** Software engineer with expertise in building software products and using development tools  
**Primary Pain Point:** Lacks time or specialized knowledge for prompt development; focuses primarily on business logic implementation; existing LLM integrations are unreliable

### User Stories & Use Cases

#### Prompt Developer Use Cases

**US-1-1: Develop Prompts**
As a *Prompt Developer*, I want to define my prompts without traditional programming languages like Python, while maintaining the benefits of source code versioning (Git) and syntax highlighting (IDE).

**US-1-2: Test prompts**
As a *Prompt Developer*, I want to test my prompts immediately after development without writing custom Python code or waiting for CI builds to execute.

**US-1-3: Evaluate prompt versions**
As a *Prompt Developer*, I want to evaluate newly developed prompts against their production versions to ensure the new version performs better than the previous one.

**US-1-4: Deploy prompts**
As a *Prompt Developer*, I want to deploy my prompts to **dev** and **prod** environments easily and reliably.

**US-1-5: Automated prompts testing**
As a *Prompt Developer*, I want to automate my prompt testing and run tests in CI after each Git push.

**US-1-6: Automated prompts deployment**
As a *Prompt Developer*, I want to automate prompt deployments through CD after each Git push to the `main` branch.

**US-1-7: Prompts Workflows**
As a *Prompt Developer*, I want to build workflows with my prompts where each step executes sequentially.

**US-1-8: Prompts Agents**
As a *Prompt Developer*, I want to build agents with my prompts and incorporate them into workflows as described in *US-1-7*.

**US-1-9: Observability & Costs Management**
As a *Prompt Developer*, I want to monitor prompt execution and track LLM costs.

**US-1-10: Easy LLM swap**
As a *Prompt Developer*, I want to switch LLM providers easily without extensive code changes.

#### Application Developer Use Cases

**US-2-1: Execute AI Flow**
As an *Application Developer*, I want to reliably execute AI Flows defined by *Prompt Developers* in the `flow-run` service to add AI integration to my application.

**US-2-2: Get AI Flow results**
As an *Application Developer*, I want to retrieve AI Flow results from the `flow-run` service when they're ready.

### Complete Feature List

#### Core Infrastructure
- Reliable execution of AI flows using fire-and-forget semantics with guaranteed execution
- Infrastructure-as-Code approach for prompt development and deployment
- Multi-LLM provider support within the execution engine
- Language-agnostic application integration

#### Development Tools
- CLI tool for running, evaluating, and deploying prompts
- Prompt development without traditional programming languages
- Source code versioning (Git) integration
- IDE syntax highlighting support

#### Testing & Quality Assurance
- Immediate prompt testing capabilities
- Prompt evaluation against production versions
- Automated prompt testing in CI pipelines
- Regression testing for prompt versions

#### Deployment & Operations
- Easy deployment to dev and prod environments
- Automated prompt deployments through CD
- Reliable prompt versioning and rollback capabilities

#### Workflow & Agent Support
- Sequential workflow execution
- AI Agent integration within workflows
- Support for common AI flow abstractions: tasks, workflows, and agents

#### Observability & Cost Management
- Monitoring of prompt execution
- LLM cost tracking and reporting
- Observability metrics for AI flow executions
- Performance analytics and insights

#### CI/CD Integration
- Automated testing after each Git push
- Continuous deployment to main branch
- Integration with popular CI/CD platforms

## Programming Rules & Guidelines

### Code Architecture & Design

1. **Follow clean code principles**: Write self-documenting code with meaningful names, small functions, and clear responsibilities
2. **Use hexagonal architecture**: Organize code into domain, app, and infra layers with clear separation of concerns
   - `domain/` - Business logic and domain entities (no external dependencies)
   - `app/` - Application services and use cases (orchestrates domain and infra)
   - `infra/` - Infrastructure concerns (database, HTTP, external APIs)
3. **Interfaces on receiver side**: Define interfaces where they are consumed, not where they are implemented
4. **Dependency injection**: Use constructor injection and avoid global state
5. **Start simple**: Don't over-engineer - implement the simplest solution that solves the problem, then refactor when needed

### Go-Specific Best Practices

6. **Follow Go idioms**: Use standard Go patterns, naming conventions, and code organization
7. **Error handling**: Always handle errors explicitly; use `errors.Is()` and `errors.As()` for error checking
8. **Context propagation**: Pass `context.Context` as the first parameter for operations that can be cancelled or have timeouts
9. **Package organization**: Keep packages focused with clear, single responsibilities
10. **Use Go modules**: Manage dependencies properly with `go mod` commands

### Code Quality & Testing

11. **Run `make lint`**: Always run linting after changes to fix code style issues
12. **Parallel tests**: All tests should be parallel using `t.Parallel()`
13. **Test naming**: Use pattern `Test<Function>If<Condition>` (e.g., `TestCreateUserIfValidInput`)
14. **Group tests logically**: Use subtests with `t.Run()` to organize related test cases
15. **Test coverage**: Maintain meaningful test coverage focusing on business logic
16. **Table-driven tests**: Use table-driven tests for testing multiple scenarios

### Performance & Security

17. **Avoid premature optimization**: Profile before optimizing; readability over micro-optimizations
18. **Memory management**: Be mindful of allocations; use object pools for high-frequency operations
19. **Input validation**: Validate all inputs at API boundaries
20. **Secrets management**: Never commit secrets; use environment variables or secret management systems
21. **SQL injection prevention**: Use parameterized queries and avoid string concatenation

### Development Workflow

22. **Atomic commits**: Make small, focused commits with descriptive messages
23. **Branch naming**: Use descriptive branch names (e.g., `feature/add-workflow-engine`, `fix/database-connection`)
24. **Code reviews**: All code must be reviewed before merging
25. **Documentation**: Document public APIs and complex business logic
26. **Configuration**: Use environment variables for configuration; provide sensible defaults

### Error Handling Patterns

27. **Wrap errors**: Use `fmt.Errorf("operation failed: %w", err)` to add context while preserving the original error
28. **Domain errors**: Create custom error types for business logic errors
29. **Graceful degradation**: Design systems to handle failures gracefully
30. **Logging**: Log errors with appropriate context but avoid excessive logging

### Database & Persistence

31. **Database migrations**: Use versioned database migrations for schema changes
32. **Transaction boundaries**: Keep transactions short and handle rollbacks properly
33. **Connection pooling**: Configure database connection pools appropriately
34. **Query optimization**: Write efficient queries and use database indexes wisely

### Git Commit Guidelines

35. **Commit message format**: Use concise, single-line commit messages describing what was done
36. **No authorship claims**: Never include co-authorship or attribution information in commit messages
37. **Action-based messages**: Use imperative mood describing the change (e.g., "Add user authentication", "Fix database connection")
38. **One line only**: Commit messages should be one line without additional descriptions or explanations

## Development Commands

### Building and Running
```bash
# Build the main application
go build -o bin/flowrun ./cmd/flowrun

# Run the application directly
go run ./cmd/flowrun

# Install dependencies (if go.sum exists)
go mod download

# Tidy dependencies
go mod tidy
```

### Testing and Quality
```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Format code
go fmt ./...

# Vet code for potential issues
go vet ./...

# Run specific package tests
go test ./internal/core
go test ./internal/flowrun
go test ./internal/flowruncli
```

### Database
The project uses PostgreSQL as its database. A Docker Compose configuration is provided:

```bash
# Start the database
docker-compose -f env/docker-compose.yaml up -d

# Stop the database
docker-compose -f env/docker-compose.yaml down
```

Database connection details:
- Host: localhost:5432
- Database: flowrun
- User: root
- Password: root

## Architecture

The project follows a **modular monolith** architecture pattern with clear separation of concerns and responsibilities:

### Module Structure

#### `internal/core`
**Domain Logic Module**
- Contains core business logic and domain entities
- Pure domain models with no external dependencies
- Shared structs and interfaces that can be reused across all modules
- Business rules and domain validation logic
- No dependencies on infrastructure or application layers

#### `internal/flowrun`
**Backend Engine Module**
- Implements the AI flow execution engine
- Provides HTTP endpoints for flow execution and configuration management
- RESTful API for creating, managing, and executing AI flows
- Handles flow orchestration and LLM provider integration
- **Primary User**: User Persona 2 (Application Developer)
- **Use Cases**: US-2-1 (Execute AI Flow), US-2-2 (Get AI Flow results)

#### `internal/flowruncli`
**CLI Client Module**
- Command-line interface for flow-run operations
- Responsible for creating flows from YAML files
- Provides commands for running, testing, and evaluating flows
- Handles flow deployment and versioning operations
- **Primary User**: User Persona 1 (Prompt Developer)
- **Use Cases**: US-1-1 through US-1-10 (all Prompt Developer use cases)

#### `internal/lib`
**Shared Utilities Module**
- Contains utility classes and helper functions
- Non-domain-specific code that can be reused across modules
- Common logging, configuration, and helper utilities
- Shared middleware and common patterns

#### `pkg/flowrunclient`
**Go HTTP Client Module**
- Public Go client library for calling flowrun backend
- Provides typed Go interface for Application Developers
- Handles HTTP communication with flowrun service
- Can be imported by external Go applications
- Simplifies integration for Go-based applications

### Entry Points

#### `cmd/flowrun/`
**Main Application Entry Point**
- Combines all modules into deployable applications
- Configuration and dependency injection setup
- Application bootstrapping and shutdown logic

#### `env/`
**Environment Configuration**
- Docker Compose configurations for development
- Environment-specific settings and configurations

### Module Dependencies

```
pkg/flowrunclient -> internal/flowrun (HTTP API)
internal/flowruncli -> internal/core -> internal/lib
internal/flowrun -> internal/core -> internal/lib
cmd/flowrun -> internal/flowrun + internal/flowruncli
```

### Hexagonal Architecture Within Modules

Each module follows hexagonal architecture principles:
- **Domain Layer**: Core business logic (in `internal/core` and domain packages within modules)
- **Application Layer**: Use cases and application services
- **Infrastructure Layer**: Database, HTTP, external API integrations

This modular monolith structure allows for:
- Clear separation of concerns between different user personas
- Independent development and testing of modules
- Potential future extraction into separate services if needed
- Shared domain logic through `internal/core`
- Reusable utilities through `internal/lib`

## Dependencies & Libraries

### Core Libraries

#### HTTP Framework
- **Gin** (`github.com/gin-gonic/gin`) - HTTP web framework for API endpoints
  - Lightweight and fast HTTP router
  - Middleware support for authentication, logging, CORS
  - JSON binding and validation

#### Logging
- **Logrus** (`github.com/sirupsen/logrus`) - Structured logging
  - JSON formatted logs for production
  - Multiple log levels (Debug, Info, Warn, Error)
  - Contextual logging with fields

#### Dependency Injection
- **Manual DI** - No framework dependency injection
  - Create dependencies explicitly in constructors
  - Makes dependencies visible and testable
  - Avoids magic and improves code clarity

#### Database & Migrations
- **golang-migrate** (`github.com/golang-migrate/migrate/v4`) - Database migrations
  - Version-controlled schema changes
  - Up/down migration support
  - Multiple database driver support

#### Task Scheduling
- **gocron** (`github.com/go-co-op/gocron/v2`) - Task scheduling and cron jobs
  - Background task execution
  - Cron-like scheduling syntax
  - Job persistence and recovery

#### Validation
- **validator** (`github.com/go-playground/validator/v10`) - Struct validation
  - Tag-based validation rules
  - Custom validation functions
  - Internationalization support

#### Testing
- **testify** (`github.com/stretchr/testify`) - Testing toolkit
  - Rich assertion library (`assert`, `require`)
  - Test suites and setup/teardown
  - Mock generation and verification

#### JSON Schema
- **jsonschema** (`github.com/invopop/jsonschema`) - JSON Schema generation
  - Generate JSON schemas from Go structs
  - Validation of JSON payloads
  - API documentation generation

### External Dependencies

#### LLM Integration
- **OpenAI Go SDK** (`github.com/openai/openai-go`) - LLM API client
  - Used to integrate with OpenRouter for multi-LLM support
  - Provides OpenAI-compatible interface
  - Supports streaming and function calling
  - **Purpose**: OpenRouter allows access to multiple LLM providers (OpenAI, Anthropic, Google, etc.) through a single API

#### Database
- **PostgreSQL** - Primary database for `internal/flowrun` module
  - Stores flow configurations, execution history, and metadata
  - ACID compliance for reliable data storage
  - JSON/JSONB support for flexible schema
  - Full-text search capabilities

### Development Dependencies

#### Dependency Management
- **Go Modules** - Native Go dependency management
- **Vendor Directory** - All dependencies stored locally
  ```bash
  go mod vendor  # Create vendor folder with all dependencies
  ```

### Dependency Philosophy

1. **Explicit Dependencies**: Manual dependency injection makes all dependencies visible
2. **Minimal External Dependencies**: Choose well-maintained, stable libraries
3. **Standard Library First**: Prefer Go standard library when possible
4. **Vendor Everything**: Use `go mod vendor` to ensure reproducible builds
5. **Interface-Based Design**: Define interfaces to make dependencies swappable

### Library Usage by Module

#### `internal/core`
- `validator` - Domain model validation
- `jsonschema` - Schema definitions
- No external infrastructure dependencies

#### `internal/flowrun`
- `gin` - HTTP API endpoints
- `logrus` - Request/response logging
- `golang-migrate` - Database schema management
- `gocron` - Background task scheduling
- `openai-go` - LLM provider integration
- PostgreSQL driver

#### `internal/flowruncli`
- `logrus` - CLI operation logging
- `validator` - YAML/JSON validation
- File system operations (standard library)

#### `internal/lib`
- `logrus` - Shared logging utilities
- Configuration helpers (standard library)
- Common middleware and utilities

#### `pkg/flowrunclient`
- HTTP client (standard library)
- JSON encoding/decoding (standard library)
- Minimal external dependencies for easy integration

## Development Environment

The project supports SQLite for local development (files ignored in .gitignore: `*.sqlite`, `*.sqlite-shm`, `*.sqlite-wal`) and PostgreSQL for production/testing via Docker Compose.

Build artifacts are placed in `bin/` directory and ignored by Git.

## Deployment

### Prototype Deployment Strategy

For the initial prototype version, flow-run will be deployed using **Docker Compose** on a **Hetzner** cloud machine. This approach provides a simple, cost-effective solution for early development and testing.

### Deployment Architecture

#### Infrastructure
- **Cloud Provider**: Hetzner Cloud
- **Server Type**: Shared vCPU instance (cost-effective for prototyping)
- **Operating System**: Ubuntu LTS
- **Container Orchestration**: Docker Compose
- **Reverse Proxy**: Nginx (containerized)
- **SSL/TLS**: Let's Encrypt certificates via Nginx

#### Services
The Docker Compose setup includes:

1. **flowrun-api** - Main application container
   - Builds from project Dockerfile
   - Exposes internal port 8080
   - Environment variables for configuration
   - Connects to PostgreSQL container

2. **postgresql** - Database container
   - Official PostgreSQL image
   - Persistent volume for data storage
   - Automated backups (planned)

3. **nginx** - Reverse proxy container
   - Routes traffic to flowrun-api
   - SSL termination
   - Static file serving (if needed)

4. **migration** - Database migration container
   - Runs on startup to apply schema changes
   - Uses golang-migrate tool
   - Exits after successful migration

#### Configuration Management
- **Environment Variables**: Store configuration and secrets
- **.env files**: Local development configuration
- **Docker Secrets**: Production secrets management
- **Health Checks**: Container health monitoring
- **Deployment Files**: All deployment configuration files stored in `deployments/` folder

#### Deployments Folder Structure
The `deployments/` folder contains all deployment-related configuration files:

```
deployments/
├── docker-compose.dev.yml          # Development environment
├── docker-compose.prod.yml         # Production environment
├── Dockerfile                      # Application container image
├── nginx/
│   ├── nginx.conf                  # Nginx configuration
│   └── ssl/                        # SSL certificate configs
├── postgres/
│   ├── init.sql                    # Database initialization
│   └── backup.sh                   # Backup scripts
├── .env.example                    # Environment variables template
├── .env.dev                        # Development environment variables
└── scripts/
    ├── deploy.sh                   # Deployment automation script
    ├── backup.sh                   # Database backup script
    └── healthcheck.sh              # Health check script
```

### Deployment Process

#### Initial Setup
```bash
# 1. Server provisioning on Hetzner
# 2. Docker and Docker Compose installation
# 3. Clone repository
git clone https://github.com/vitalii-honchar/flow-run.git
cd flow-run

# 4. Configure environment
cp deployments/.env.example deployments/.env.production
# Edit deployments/.env.production with production values

# 5. Build and deploy
docker-compose -f deployments/docker-compose.prod.yml up -d
```

#### CI/CD Pipeline (Future)
- **Git Push** → GitHub Actions
- **Build** → Docker image creation
- **Test** → Automated testing in containers
- **Deploy** → Push to Hetzner machine
- **Health Check** → Verify deployment success

### Environment Configuration

#### Required Environment Variables
```bash
# Database
DB_HOST=postgresql
DB_PORT=5432
DB_NAME=flowrun
DB_USER=flowrun
DB_PASSWORD=${DB_PASSWORD}

# API Configuration
API_PORT=8080
API_HOST=0.0.0.0
LOG_LEVEL=info

# LLM Integration
OPENROUTER_API_KEY=${OPENROUTER_API_KEY}
OPENROUTER_BASE_URL=https://openrouter.ai/api/v1

# Domain and SSL
DOMAIN=api.flowrun.dev
LETSENCRYPT_EMAIL=${ADMIN_EMAIL}
```

### Monitoring and Maintenance

#### Basic Monitoring
- **Container Health**: Docker health checks
- **Application Logs**: Centralized logging via Docker
- **Database Monitoring**: PostgreSQL metrics
- **Disk Space**: Automated cleanup of old logs

#### Backup Strategy
- **Database Backups**: Daily PostgreSQL dumps
- **Configuration Backup**: Environment files and compose configs
- **Application State**: Minimal state to backup (stateless design)

#### Security Considerations
- **Firewall**: UFW configured for minimal port exposure
- **SSL/TLS**: HTTPS only with automatic certificate renewal
- **Container Security**: Non-root containers where possible
- **Network Isolation**: Docker networks for service communication
- **Regular Updates**: Automated security updates for the host OS

### Scaling Considerations

This deployment approach supports the prototype phase with clear upgrade paths:

#### Current Limitations
- **Single Server**: No high availability
- **Vertical Scaling Only**: Limited by single machine resources
- **Manual Deployment**: No automated CI/CD initially

#### Future Migration Path
- **Kubernetes**: For horizontal scaling and high availability
- **Cloud Database**: Managed PostgreSQL for better reliability
- **Load Balancer**: Multiple application instances
- **CDN Integration**: For static assets and global performance

### Cost Optimization

#### Hetzner Advantages
- **Cost-Effective**: Lower costs compared to AWS/GCP for prototype
- **European Data Center**: GDPR compliance ready
- **Simple Pricing**: Predictable monthly costs
- **Good Performance**: Dedicated resources even on shared plans

#### Resource Planning
- **Initial Setup**: 2 vCPU, 4GB RAM, 40GB SSD (~€5-10/month)
- **Database Storage**: Separate volume for persistence
- **Backup Storage**: Additional storage for database backups
- **Bandwidth**: Sufficient for prototype traffic levels

This deployment strategy balances simplicity, cost-effectiveness, and functionality for the prototype phase while maintaining a clear path for future scaling.