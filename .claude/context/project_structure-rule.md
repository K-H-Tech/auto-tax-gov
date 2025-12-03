# Go Project Structure Best Practices Guide

This document outlines best practices for structuring Go projects to promote maintainability, scalability, and collaboration. It draws from the official Go documentation on module layout, the widely referenced "Standard Go Project Layout" from the golang-standards community, and insights from talks like Peter Bourgon's "Best Practices for Industrial Programming" (GopherCon EU 2018) and Kat Zien's "How Do You Structure Your Go Apps" (GopherCon 2018). These guidelines emphasize simplicity: start minimal and add structure as needed. No single layout fits all projects—adapt based on size, type (CLI, web service, library), and team preferences. For small projects (<10k LOC), a flat structure suffices; for larger ones, use layered or modular designs.

The recommended layout balances Go idioms (e.g., `internal` for private code) with practical organization for non-Go assets (e.g., configs, scripts). Avoid over-engineering; refactor as the project grows. Use `go mod init` for modules and `go:generate` for automation.

## Core Principles

- **Start Simple**: Begin with a flat structure (all code in root) for prototypes. Evolve to `cmd/internal/pkg` as complexity increases.
- **Visibility Control**: Use `internal/` to hide private packages from external imports (enforced by Go compiler).
- **Single Responsibility**: Packages should be small and focused; avoid god packages. Group by use case (modular/hexagonal) over type (e.g., all models together).
- **Import Paths**: Keep them short and descriptive; use snake_case for directories.
- **Non-Go Assets**: Separate configs, scripts, and docs to keep the root clean.
- **Testing**: Place tests alongside code (`*_test.go`) or in `/test/` for integration/e2e.
- **Avoid**: `/src` (legacy GOPATH artifact); over-abstraction early; mixing public/private code.

## Recommended Directory Layout

Use this as a starting template for medium-to-large projects (e.g., web services with CLI tools). For libraries, omit `/cmd` and focus on `/pkg`. For tiny CLIs, just root + `go.mod`.

```
myproject/
├── cmd/                 # Applications (main packages)
│   ├── myapp/           # Main web server binary
│   │   └── main.go
│   └── myappctl/        # CLI tool binary
│       └── main.go
├── internal/            # Private app/lib code (unimportable externally)
│   ├── app/             # App-specific packages (e.g., handlers, services)
│   │   ├── handlers/
│   │   ├── services/
│   │   └── models/      # Domain models (use cases over types)
│   └── pkg/             # Private reusable libs (e.g., utils, db)
│       ├── db/
│       └── utils/
├── pkg/                 # Public reusable libs (importable externally)
│   └── mypubliclib/     # e.g., shared API client
├── api/                 # OpenAPI/Swagger specs, protobufs
│   └── openapi.yaml
├── web/                 # Web assets (templates, static files)
│   ├── templates/
│   └── static/
├── configs/             # Config templates/defaults
│   └── config.yaml
├── scripts/             # Build/deploy scripts
│   └── build.sh
├── build/               # Packaging/CI configs
│   └── docker/
├── deployments/         # Deployment manifests (k8s, docker-compose)
│   └── kubernetes/
├── test/                # Integration/e2e tests, testdata
│   ├── integration/
│   └── testdata/
├── docs/                # Design/user docs (beyond godoc)
│   └── README.md        # Project overview, setup, structure
├── go.mod               # Module definition
├── go.sum               # Dependency checksums
└── Makefile             # High-level build tasks (keep simple)
```

- **cmd/**: Entry points for executables. Subdirs named after binaries (e.g., `myapp`). Keep minimal—wire dependencies and run servers/CLIs. Import from `internal/` and `pkg/`. Multiple binaries? Multiple subdirs.
- **internal/**: Core app logic. Unimportable outside the module. Use subdirs like `app/` for domain-specific code (handlers, services) and `pkg/` for utilities (db, auth). Enforces encapsulation.
- **pkg/**: Public, reusable code (e.g., shared libs). Optional for apps; essential for libraries. Avoid if code isn't go-gettable.
- **api/**: Interface definitions (Swagger, gRPC protos). Keeps contracts separate.
- **web/**: Frontend assets for web apps. Use if serving static files or templates.
- **configs/**: YAML/TOML templates. Load via Viper; gitignore secrets.
- **scripts/**: Automation (e.g., lint, build). Call from root Makefile.
- **build/deployments/**: CI/CD pipelines, Dockerfiles, k8s manifests.
- **test/**: Non-unit tests (e2e, load). Use `testdata/` for fixtures (Go ignores it).
- **docs/**: Markdown guides. Use godoc for package docs.

## Examples

### Minimal CLI Project (Flat Structure)
For a simple tool (<1k LOC):
```
mycli/
├── main.go              # All logic here
├── go.mod
└── README.md
```
- Pros: Zero overhead. Cons: Scales poorly.

### Web Service with CLI (Layered/Modular)
Extend the recommended layout for a Gin-based API with Cobra CLI:
```
myproject/
├── cmd/
│   ├── api/             # Gin server
│   │   └── main.go      # Wires router, DB, etc.
│   └── cli/             # Cobra commands
│       └── main.go
├── internal/
│   ├── app/
│   │   ├── handlers/    # Gin handlers
│   │   │   └── user.go
│   │   ├── services/    # Business logic
│   │   │   └── usersvc.go
│   │   └── models/      # Domain types
│   │       └── user.go
│   └── pkg/
│       ├── db/          # pgx wrapper
│       │   └── postgres.go
│       └── resilience/  # Retry, circuit breaker
│           └── retry.go
├── api/
│   └── openapi.yaml
├── configs/
│   └── config.yaml      # Viper defaults
├── test/
│   └── integration/
│       └── db_test.go
├── go.mod
└── Makefile             # go build ./cmd/api
```
- `internal/app/handlers/user.go`: Gin handler imports `services.UserService`.
- `cmd/api/main.go`:
  ```go
  package main

  import (
      "github.com/gin-gonic/gin"
      "myproject/internal/app/handlers"
      "myproject/internal/pkg/db"
      "github.com/spf13/viper"
  )

  func main() {
      viper.SetConfigFile("configs/config.yaml")
      viper.ReadInConfig()
      dbConn := db.New(viper.GetString("db.url"))
      r := gin.Default()
      handlers.RegisterUserRoutes(r, dbConn)
      r.Run(":8080")
  }
  ```
- Run: `go build -o bin/api ./cmd/api`

### Library Project
Omit `/cmd`; focus on `/pkg/myproject/` with examples in `/examples/`.
```
mylib/
├── pkg/
│   └── myproject/
│       └── mylib.go
├── examples/
│   └── simple/
│       └── main.go
├── go.mod
└── README.md
```

## Integration with Other Tools

- **Viper/Cobra**: Configs in `/configs/`, CLI in `/cmd/cli/`.
- **Gin/JWT/Zap**: Handlers in `/internal/app/handlers/`, middleware in `/internal/pkg/resilience/`.
- **pgx/Mongo/Redis**: Wrappers in `/internal/pkg/db/` or `/internal/pkg/cache/`.
- **Prometheus/Metrics**: Expose in `/internal/pkg/metrics/`.
- **Retry/Circuit Breaker/Rate Limit**: In `/internal/pkg/resilience/`.
- **Mocking**: Generate in `/internal/pkg/mocks/` via `go:generate`.
- **Docker/K8s**: Files in `/build/` and `/deployments/`.

## Testing and CI

- **Unit Tests**: `*_test.go` next to code; use `go test ./...`.
- **Integration**: In `/test/integration/`; run via Makefile.
- **CI**: Use `/build/ci/` for GitHub Actions/Travis configs. Lint with `golangci-lint`.
- **Makefile Example**:
  ```makefile
  test:
      go test ./...
  build:
      go build -o bin/api ./cmd/api
  lint:
      golangci-lint run
  ```

## Additional Resources

- Standard Go Project Layout: https://github.com/golang-standards/project-layout
- Official Go Module Layout: https://go.dev/doc/modules/layout
- Talks: Peter Bourgon's Industrial Go (GopherCon EU 2018); Kat Zien's Structuring Go Apps (GopherCon 2018)

Adhere to these guidelines to build scalable Go projects. Review and refactor structure periodically as the codebase evolves.