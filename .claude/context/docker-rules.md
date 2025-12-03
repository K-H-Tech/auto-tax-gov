# Docker and Docker Compose Best Practices Guide

This document outlines best practices for writing Dockerfiles and docker-compose.yml files in Go projects. It draws from the official Docker documentation on Dockerfile best practices (e.g., multi-stage builds, security, optimization) and Docker Compose environment management, as well as community guidelines for production-ready setups. These practices ensure secure, efficient, and maintainable containerization for Go applications, integrating with tools like Gin, Viper, Cobra, pgx, and Prometheus. Focus on minimal images, reproducibility, and scalability. Use Docker 24+ and Compose v2+ for Go 1.21+ projects.

## Core Principles

- **Security First**: Run as non-root, use minimal base images, scan for vulnerabilities (e.g., Trivy), and avoid unnecessary packages to reduce attack surface.
- **Efficiency**: Leverage build cache, multi-stage builds, and .dockerignore to minimize size and build time.
- **Reproducibility**: Pin image versions/digests; use environment variables for configs.
- **Compose Modularity**: Use multiple files for environments (dev/prod), named volumes, and health checks.
- **Go-Specific**: Use multi-stage with `golang` for build and `scratch`/`alpine` for runtime to create tiny, static binaries.

## Dockerfile Best Practices

Place Dockerfiles in project root or `/build/`; use `.dockerignore` to exclude `.git`, `node_modules`, tests, etc. Build with `docker build -t myapp .` or via Makefile.

### 1. Use Minimal, Official Base Images
- Prefer `golang:1.21-alpine` for builds (small, secure) and `scratch` or `alpine` for runtime.
- Pin versions/digests for reproducibility (e.g., `alpine:3.21@sha256:...`).
- Avoid full OS images (e.g., `ubuntu`) unless needed; they bloat size/vulnerabilities.

### 2. Multi-Stage Builds for Go Apps
- Separate build (compile) and runtime stages to exclude toolchain, reducing size to ~10MB.
- **Example Dockerfile for Go Web App (Gin)**:
  ```dockerfile:disable-run
  # syntax=docker/dockerfile:1
  # Build stage
  FROM golang:1.21-alpine AS builder
  WORKDIR /app
  # Copy go.mod/go.sum first for cache
  COPY go.mod go.sum ./
  RUN go mod download
  # Copy source
  COPY . .
  # Build static binary
  RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api/main.go

  # Runtime stage
  FROM alpine:3.21
  RUN apk --no-cache add ca-certificates tzdata  # For HTTPS/timezone
  WORKDIR /root/
  COPY --from=builder /app/main .
  # Expose port
  EXPOSE 8080
  # Run as non-root
  USER 1000:1000
  ENTRYPOINT ["./main"]
  ```
  - **Build**: `docker build -t myapp .`
  - **Tips**: Copy `go.mod` first for dependency caching; disable CGO for static binary; add healthcheck if needed (`HEALTHCHECK CMD wget --no-verbose --tries=1 --sput-output=/dev/null http://localhost:8080/health || exit 1`).

### 3. Optimize Layers and Cache
- Order instructions: Stable (e.g., `RUN apk add`) before volatile (e.g., `COPY .`).
- Combine `RUN` commands to reduce layers (e.g., `RUN apt-get update && apt-get install -y pkg && rm -rf /var/lib/apt/lists/*`).
- Use `.dockerignore`:
  ```
  **/.git
  **/node_modules
  **/*.md
  .env
  ```
- Rebuild often with `--no-cache` for security patches.

### 4. Security Enhancements
- Run as non-root: `USER 1000:1000` or create custom user.
- Minimal packages: Only install essentials (e.g., `ca-certificates` for HTTPS).
- Scan images: Integrate Trivy (`trivy image myapp`) in CI.
- Labels: Add metadata (`LABEL maintainer="team@example.com"`).

### 5. Go-Specific Tips
- Static binaries: `CGO_ENABLED=0` for cross-platform compatibility.
- Distroless: Alternative to `scratch` for runtime with glibc (`gcr.io/distroless/static-debian12`).
- Healthchecks: For Gin apps, check `/health` endpoint.

## docker-compose.yml Best Practices

Place `docker-compose.yml` in project root; use overrides (e.g., `docker-compose.dev.yml`) for environments. Run with `docker compose up -d`.

### 1. Environment Management
- Use `.env` files for secrets/vars; reference with `${VAR}`. Never commit `.env`.
- Multiple files: `docker compose -f docker-compose.yml -f docker-compose.prod.yml up`.
- **Example .env**:
  ```
  DB_URL=postgres://user:pass@db:5432/mydb
  MINIO_ENDPOINT=play.min.io
  ```

### 2. Service Definition and Networking
- Explicit networks: Use `networks:` for isolation; default bridge for simple setups.
- Depends_on with healthchecks: Wait for dependencies.
- **Example docker-compose.yml for Go App + Postgres + MinIO**:
  ```yaml
  version: '3.8'
  services:
    api:
      build: .  # Uses Dockerfile in root
      ports:
        - "8080:8080"
      environment:
        - DB_URL=${DB_URL}
        - MINIO_ENDPOINT=${MINIO_ENDPOINT}
      volumes:
        - .:/app  # Dev only
      depends_on:
        db:
          condition: service_healthy
        minio:
          condition: service_healthy
      networks:
        - app-network
      healthcheck:
        test: ["CMD", "wget", "--no-verbose", "--tries=1", "--sput-output=/dev/null", "http://localhost:8080/health"]
        interval: 30s
        timeout: 10s
        retries: 3

    db:
      image: postgres:16-alpine
      environment:
        POSTGRES_DB: mydb
        POSTGRES_USER: user
        POSTGRES_PASSWORD: pass
      volumes:
        - pgdata:/var/lib/postgresql/data  # Named volume
      healthcheck:
        test: ["CMD-SHELL", "pg_isready -U user -d mydb"]
        interval: 10s
        timeout: 5s
        retries: 5
      networks:
        - app-network

    minio:
      image: minio/minio:latest
      command: server /data --console-address ":9001"
      environment:
        MINIO_ROOT_USER: minioadmin
        MINIO_ROOT_PASSWORD: minioadmin
      ports:
        - "9000:9000"
        - "9001:9001"
      volumes:
        - miniodata:/data
      healthcheck:
        test: ["CMD", "curl", "-f", "http://localhost:9000/minio/health/live"]
        interval: 30s
      networks:
        - app-network

  volumes:
    pgdata:
    miniodata:

  networks:
    app-network:
      driver: bridge
  ```
  - **Run**: `docker compose up -d`; scale with `docker compose up --scale api=3`.
  - **Tips**: Use `condition: service_healthy` for deps; named volumes for persistence.

### 3. Resource Limits and Scaling
- Set CPU/memory: `deploy.resources.limits.memory: 512M` for prod.
- Healthchecks: Essential for orchestration.

### 4. Production Tips
- Secrets: Use `secrets:` or external tools (e.g., Docker Secrets).
- Logging: `logging.driver: "json-file"` with rotation.
- Avoid dev volumes in prod; use bind mounts sparingly.

## Integration with Go Project Structure

- **Build in Makefile**: `docker-build: docker build -t myapp .`
- **Compose with Viper**: Load env vars in Go code.
- **CI/CD**: Build/test in GitHub Actions; push to registry.
- **Prometheus**: Expose `/metrics` in Gin; scrape via Compose.

## Testing and CI

- **Local**: `docker compose up --build`; test with `docker compose exec api go test ./...`
- **CI**: Use `docker/build-push-action` in GitHub Actions; scan with Trivy.
- **Security Scans**: Integrate Hadolint for Dockerfile linting.

## Additional Resources

- Official Dockerfile Best Practices: https://docs.docker.com/build/building/best-practices/
- Docker Compose Best Practices: https://docs.docker.com/compose/best-practices/
- Go Docker Example: https://github.com/golang/go/wiki/Docker

Adhere to these practices for robust Docker setups in Go projects. Update as Docker evolves.
```