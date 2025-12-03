# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

A Go-based multi-service application for interacting with Iranian government portals through the central SSO (sso.my.gov.ir). The application provides:

- **MyGovAuth Service** - Core SSO authentication (captcha, OTP, session management)
- **MyTax Service** - Tax portal operations (my.tax.gov.ir)
- **Enamad Service** - E-trust symbol verification (reg2.enamad.ir)
- **Mojavez Service** - Business licensing (mojavez.ir)

The application simulates browser-like behavior to track redirect chains, handle captcha verification, and perform OTP-based authentication.

## Running the Application

```bash
# Build the CLI
go build -o auto-tax-gov ./cmd/auto-tax-gov

# Start the web server (default port 8080)
./auto-tax-gov serve

# Start on a custom port
./auto-tax-gov serve --port 3000

# Run redirect tracking from CLI
./auto-tax-gov track

# Show version
./auto-tax-gov version

# Get help
./auto-tax-gov --help
```

### Configuration

Copy `config.example.yaml` to `config.yaml` and update with your values:
```bash
cp config.example.yaml config.yaml
```

Configuration can also be set via environment variables with `AUTO_TAX_GOV_` prefix:
```bash
AUTO_TAX_GOV_SERVER_PORT=3000 ./auto-tax-gov serve
```

## Project Structure

```
tax-gov/
├── cmd/taxgov/               # CLI entry points (Cobra)
│   ├── main.go               # Main entry point
│   ├── root.go               # Root command and config init
│   ├── serve.go              # Web server command
│   └── track.go              # CLI tracking command
├── internal/
│   ├── client/               # HTTP client factory
│   │   └── http.go           # Client creation, headers helpers
│   ├── config/               # Configuration (Viper)
│   │   └── config.go         # Config structs for all services
│   ├── models/               # Shared data types
│   │   └── types.go          # RedirectStep, CaptchaData, etc.
│   ├── session/              # Shared session management
│   │   └── session.go        # Thread-safe session state
│   ├── server/               # HTTP server
│   │   ├── handlers.go       # API endpoint handlers
│   │   └── server.go         # Server lifecycle & routing
│   └── service/              # Business logic services
│       ├── mygovauth/        # Government SSO authentication
│       │   ├── service.go    # Auth flow, redirect tracking
│       │   ├── captcha.go    # Captcha fetching, CSRF extraction
│       │   └── otp.go        # OTP send/verify operations
│       ├── mytax/            # Tax portal service
│       │   └── service.go    # Dashboard, registration operations
│       ├── enamad/           # Enamad service (placeholder)
│       │   └── service.go    # E-trust symbol operations
│       └── mojavez/          # Mojavez service (placeholder)
│           └── service.go    # Business licensing operations
├── web/                      # Static web assets
│   └── index.html            # Persian UI
├── .claude/                  # Development rules & documentation
│   ├── agents/               # AI agent configurations
│   ├── rules/                # Code quality rules
│   ├── context/              # Technology reference rules (20 files)
│   └── redirect_steps/       # OAuth2 flow documentation
├── config.example.yaml       # Example configuration
├── go.mod                    # Go module definition
└── CLAUDE.md                 # This file
```

## CLI Reference

### Commands

| Command | Description | Example |
|---------|-------------|---------|
| `serve` | Start web server with UI | `./auto-tax-gov serve -p 3000` |
| `track` | CLI redirect tracking | `./auto-tax-gov track -o ./output` |
| `version` | Show version info | `./auto-tax-gov version` |

### Global Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--config` | `./config.yaml` | Config file path |
| `--log-level` | `info` | Log level (debug, info, warn, error) |

### Serve Command Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--port` | `-p` | `8080` | Server port |
| `--web-dir` | `-w` | `web` | Static assets directory |

### Track Command Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--output` | `-o` | `.` | Output directory |
| `--json` | `-j` | `false` | JSON output format |

## Architecture

### Service-Based Design

The application follows a service-oriented architecture where:
1. **MyGovAuth** is the core authentication service used by all other services
2. Each portal service (MyTax, Enamad, Mojavez) depends on MyGovAuth for authentication
3. Services are injected via dependency injection for testability

```
                    ┌─────────────────┐
                    │   HTTP Server   │
                    │   (handlers)    │
                    └────────┬────────┘
                             │
                    ┌────────▼────────┐
                    │  MyTax Service  │
                    └────────┬────────┘
                             │ uses
                    ┌────────▼────────┐
                    │ MyGovAuth Svc   │
                    │ (SSO, Captcha,  │
                    │  OTP, Session)  │
                    └─────────────────┘
```

### CLI Layer (cmd/taxgov/)
- **Cobra** for command-line interface with `serve` and `track` subcommands
- **Viper** for configuration management (YAML files + environment variables)
- Structured logging via `log/slog`

### Services (internal/service/)

| Service | Purpose | Key Operations |
|---------|---------|----------------|
| `mygovauth` | SSO authentication | InitiateAuth, FetchCaptcha, SendOTP, VerifyOTP |
| `mytax` | Tax portal | AccessDashboard, StartTaxFileRegistration |
| `enamad` | E-trust symbol | (placeholder - future implementation) |
| `mojavez` | Business licensing | (placeholder - future implementation) |

### API Endpoints (internal/server/)
- `GET /` - Serves the web UI
- `POST /api/start-tracker` - Initiates login flow and fetches captcha
- `POST /api/refresh-captcha` - Fetches new captcha using existing session
- `POST /api/send-otp` - Sends OTP request with mobile + captcha
- `POST /api/verify-otp` - Verifies OTP code and completes login
- `POST /api/access-dashboard` - Accesses the tax dashboard
- `POST /api/start-tax-file` - Initiates tax file registration

### Core Components
- **Session** (`internal/session/`) - Thread-safe session state with `sync.RWMutex`
- **Client** (`internal/client/`) - HTTP client factory with no-redirect policy
- **Config** (`internal/config/`) - Multi-service configuration via Viper

### Frontend (web/index.html)
- Single-page RTL Persian UI with multi-step login flow
- Steps: Load captcha → Enter mobile + captcha → Receive OTP → Verify OTP → Access dashboard

## Key Implementation Details

- **Service Composition**: MyTax/Enamad/Mojavez services compose MyGovAuth for authentication
- **HTTP Client**: Configured to NOT auto-follow redirects (`http.ErrUseLastResponse`)
- **Session Management**: Thread-safe session with cookie merging across services
- **CSRF Protection**: Tokens extracted from HTML via regex patterns
- **Captcha**: Fetched from `/client/v1/captcha` with Basic auth header
- **Response Handling**: Automatic gzip decompression
- **Configuration**: All URLs stored in config, not hardcoded
- **Dependency Injection**: Services injected via constructors for testability

---

## Development Standards

This project follows a hierarchical rule system. CLAUDE.md provides project context, detailed standards are in `.claude/` subdirectories.

### Master Rule Index

| Category | Location | When to Apply |
|----------|----------|---------------|
| **Go Engineering** | `.claude/agents/go-engineer.md` | All Go code changes |
| **Code Quality** | `.claude/rules/code-quality.md` | All code changes |
| **Security** | `.claude/rules/security.md` | Auth, input handling, secrets |
| **Testing** | `.claude/rules/testing.md` | Writing/modifying tests |
| **API Design** | `.claude/rules/api-design.md` | HTTP endpoints, JSON responses |
| **Technology Rules** | `.claude/context/*.md` | See Technology Context section below |
| **OAuth Flow Docs** | `.claude/redirect_steps/` | Tax portal integration |

---

## Technology Context Rules

Reference documentation for technologies used in this project. Located in `.claude/context/`.

### Complete Technology Index

| File | Technology | Purpose | Key Library |
|------|-----------|---------|-------------|
| `circuit-breaker-rule.md` | Resilience | Fault tolerance, cascading failure prevention | `sony/gobreaker/v2` |
| `cobra-rule.md` | CLI | Command-line application structure | `spf13/cobra` |
| `design_pattern-rule.md` | Architecture | Go design patterns & idioms | Go conventions |
| `docker-rules.md` | Containerization | Multi-stage builds, deployment | Docker |
| `gin-rule.md` | Web Framework | REST API development | `gin-gonic/gin` |
| `gitlab_ci-rule.md` | CI/CD | Pipeline automation | GitLab CI |
| `jwt-rule.md` | Authentication | Token-based auth | `golang-jwt/jwt/v5` |
| `MinIO-rule.md` | Object Storage | S3-compatible file storage | `minio/minio-go/v7` |
| `mock-rule.md` | Testing | Mock generation | `uber-go/mock` |
| `mongodb-rule.md` | Database (NoSQL) | Document store | `mongo-go-driver/v2` |
| `open_telementry-rule.md` | Observability | Traces, metrics, logs | OpenTelemetry |
| `pgx-rule.md` | Database (SQL) | PostgreSQL driver | `jackc/pgx/v5` |
| `project_structure-rule.md` | Organization | Go project layout | Go conventions |
| `prometheus-rule.md` | Metrics | Application monitoring | `prometheus/client_golang` |
| `rabbitmq-rule.md` | Messaging | Event-driven architecture | `amqp091-go` |
| `rate-limiter-rule.md` | Throttling | Request rate limiting | `x/time/rate` |
| `redis-rule.md` | Caching | In-memory cache/sessions | `go-redis/v9` |
| `retry-rule.md` | Resilience | Transient failure handling | `avast/retry-go/v4` |
| `viper-rule.md` | Configuration | Multi-source config management | `spf13/viper` |
| `zap-rule.md` | Logging | Structured logging | `uber-go/zap` |

### Rules by Category

#### Core Infrastructure
| Rule | When to Apply |
|------|---------------|
| `project_structure-rule.md` | Project organization, directory layout |
| `docker-rules.md` | Containerization, multi-stage builds |
| `gitlab_ci-rule.md` | CI/CD pipelines, automated testing |

#### Web & API
| Rule | When to Apply |
|------|---------------|
| `gin-rule.md` | HTTP handlers, middleware, routing |
| `cobra-rule.md` | CLI commands, flags, subcommands |
| `viper-rule.md` | Configuration files, env vars, flags |

#### Data Layer
| Rule | When to Apply |
|------|---------------|
| `pgx-rule.md` | PostgreSQL queries, transactions, pooling |
| `mongodb-rule.md` | MongoDB operations, aggregations |
| `redis-rule.md` | Caching, sessions, pub/sub |
| `MinIO-rule.md` | File uploads, object storage |

#### Messaging & Events
| Rule | When to Apply |
|------|---------------|
| `rabbitmq-rule.md` | Message queues, pub/sub, work queues |

#### Security & Auth
| Rule | When to Apply |
|------|---------------|
| `jwt-rule.md` | Token generation, validation, claims |

#### Resilience & Performance
| Rule | When to Apply |
|------|---------------|
| `circuit-breaker-rule.md` | Fault tolerance, cascading failure prevention |
| `retry-rule.md` | Transient failures, exponential backoff |
| `rate-limiter-rule.md` | Request throttling, API protection |

#### Observability
| Rule | When to Apply |
|------|---------------|
| `zap-rule.md` | Structured logging, log levels |
| `prometheus-rule.md` | Metrics collection, instrumentation |
| `open_telementry-rule.md` | Distributed tracing, spans |

#### Testing & Patterns
| Rule | When to Apply |
|------|---------------|
| `mock-rule.md` | Unit test mocking with mockgen |
| `design_pattern-rule.md` | Repository, DI, middleware patterns |

---

## OAuth2 Authentication Flow

Complete documentation of the Iranian Tax Portal OAuth2 flow. Located in `.claude/redirect_steps/`.

### Flow Overview

```
1. OAuth Init → 2. HTTPS Upgrade → 3. Login (Captcha+OTP) → 4-6. Post-Login → 7. Auth Code → 8. Callback → 9. Authenticated
```

### Key Endpoints

| Endpoint | URL |
|----------|-----|
| SSO Authorization | `https://sso.my.gov.ir/oauth2/authorize` |
| Login Page | `https://sso.my.gov.ir/login` |
| Logged-in Handler | `https://sso.my.gov.ir/logged-in` |
| Tax Callback | `https://my.tax.gov.ir/myiran/sso` |
| Tax Home | `https://my.tax.gov.ir/` |

### Step Documentation Index

| Step | File | Status | Description |
|------|------|--------|-------------|
| 1 | `step_01.md` | 302 | OAuth2 authorization request to SSO |
| 2 | `step_02.md` | 302 | HTTP→HTTPS security upgrade |
| 3 | `step_03.md` | 200 | Login page with captcha + OTP form |
| - | `step_03_captcha.png` | - | Sample captcha image |
| 4 | `step_04.md` | 302 | Post-OTP redirect to logged-in |
| 5 | `step_05.md` | 302 | Logged-in handler with `&continue` |
| 6 | `step_06.md` | 302 | OAuth HTTPS security redirect |
| 7 | `step_07.md` | 302 | **Critical:** Authorization code generation |
| 8 | `step_08.md` | 302 | Callback with code→token exchange |
| 9 | `step_09.md` | 200 | Final authenticated home page |

### Summary Files

| File | Description |
|------|-------------|
| `summary.md` | Quick reference of all steps |
| `summary_full.md` | Comprehensive technical documentation |

### Authentication Details

- **Method:** Mobile OTP + Captcha
- **CSRF Protection:** State parameter validated throughout flow
- **Session Cookies:** SESSION, f5avraaaaaaaaaaaaaaaa_session_, s2
- **Security:** HTTPS enforcement, HSTS, anti-XSS headers

---

## Before Implementation

1. Review this file for project architecture
2. Check `.claude/agents/go-engineer.md` for Go standards
3. Apply relevant rules from `.claude/rules/`
4. Consult `.claude/context/` for technology-specific guidance
5. Reference `.claude/redirect_steps/` for OAuth flow understanding
6. Run `go test -race` and linters before committing

## Implementation Checklist

Before writing or modifying code:

- [ ] Understand the relevant architecture section above
- [ ] Review applicable rules in `.claude/agents/` or `.claude/rules/`
- [ ] Check `.claude/context/` for technology-specific best practices
- [ ] For Go: Apply go-engineer agent standards
- [ ] Run tests and linters before committing
