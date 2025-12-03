# GitLab CI/CD Best Practices Guide for .gitlab-ci.yml

This document outlines best practices for writing `.gitlab-ci.yml` files in Go projects using GitLab CI/CD. It draws from the official GitLab documentation on YAML syntax, pipeline configuration, optimization, and examples, as well as community recommendations (e.g., from Medium articles and Reddit discussions) to ensure efficient, secure, and scalable pipelines. These guidelines focus on Go-specific workflows like linting, testing, building static binaries, Docker image creation, and deployment, integrating with tools like Gin, Viper, Cobra, pgx, Prometheus, and Docker. Use GitLab 17+ and Go 1.21+ for optimal compatibility. The `.gitlab-ci.yml` file, placed at the project root, defines stages, jobs, and scripts for automated CI/CD pipelines, enabling early bug detection, code compliance, and streamlined deployments.

## Core Principles

- **Simplicity and Readability**: Start with a minimal pipeline; use comments (`#`) for clarity. Avoid deep nesting (limit `extends` to 3 levels) and prefer `rules` over complex `only/except` (deprecated).
- **Efficiency**: Leverage caching, artifacts, and parallel jobs to reduce build times. Pin versions (e.g., images, refs) for reproducibility; use multi-stage Docker builds for Go binaries.
- **Security**: Mask sensitive variables (e.g., API keys); run as non-root in jobs; scan for vulnerabilities (e.g., Trivy for Docker). Use protected/masked variables for branches/tags.
- **Modularity**: Split configs with `include` (local/remote/project/template); use `extends`, `!reference`, and YAML anchors for reuse. Limit to 150 includes (30s resolution limit).
- **Error Handling**: Use `allow_failure: true` for non-blocking jobs; `retry` for flakiness; `rules` for conditional execution.
- **Go-Specific**: Cache `go.mod`/vendor; build static binaries (`CGO_ENABLED=0`); use `golangci-lint` for linting.

## Project Structure

Organize `.gitlab-ci.yml` in the root; use includes for modularity (e.g., `include: - local: ci/templates/lint.yml`). For complex projects, create a shared repo (e.g., `ci-cd-assets`) for templates.

- **Recommended Layout**:
  ```
  myproject/
  ├── .gitlab-ci.yml         # Main pipeline config
  ├── ci/                    # Local includes
  │   ├── templates/
  │   │   ├── lint.yml       # Linting jobs
  │   │   ├── test.yml       # Testing jobs
  │   │   └── build.yml      # Build/Docker jobs
  │   └── scripts/           # Reusable scripts (e.g., deploy.sh)
  ├── .dockerignore          # Exclude from Docker builds
  ├── Makefile               # Local build/test targets
  ├── go.mod                 # Go modules
  └── README.md              # Pipeline overview
  ```
- **Shared Templates**: Use a central project for reusable YAML (e.g., `include: - project: 'group/ci-templates' ref: v1.0 file: /go-pipeline.yml'`). Pin `ref` (tag/branch/commit) to avoid breaking changes.

## Configuration

- **Stages**: Define sequentially (e.g., `build`, `test`, `lint`, `deploy`). Jobs in the same stage run in parallel.
  ```yaml
  stages:
    - lint
    - test
    - build
    - deploy
  ```
- **Variables**: Global (`variables:`) or per-job; use `${VAR}` for interpolation. Mask secrets (e.g., `MINIO_SECRET`).
  ```yaml
  variables:
    GO_VERSION: "1.21"
    DOCKER_IMAGE: "$CI_REGISTRY_IMAGE:$CI_COMMIT_SHA"
  ```
- **Rules**: Conditionally include jobs (e.g., only on main branch).
  ```yaml
  rules:
    - if: $CI_COMMIT_BRANCH == "main"
    - changes:
        - "**/*.go"
      when: on_success
    - when: manual  # Manual trigger
  ```
- **Caching**: Cache Go modules/vendor for faster builds.
  ```yaml
  cache:
    key: "${CI_COMMIT_REF_SLUG}"
    paths:
      - vendor/
      - ~/.cache/go-build/
    policy: pull-push
  ```
- **Artifacts**: Pass binaries/reports between jobs; expire after 1 week.
  ```yaml
  artifacts:
    paths:
      - bin/
      - coverage/
    expire_in: 1 week
    reports:
      junit: coverage/junit.xml  # For test reports
  ```
- **Includes**: Reuse templates.
  ```yaml
  include:
    - local: 'ci/templates/test.yml'
    - template: Go.gitlab-ci.yml  # Official Go template
  ```

## Examples

### Basic Go Pipeline (Lint, Test, Build)
A simple `.gitlab-ci.yml` for a Go web app with Gin, pgx, and Prometheus.
```yaml
# .gitlab-ci.yml
stages:
  - lint
  - test
  - build

variables:
  GO_VERSION: "1.21"

# Lint with golangci-lint
lint:
  stage: lint
  image: golang:${GO_VERSION}-alpine
  script:
    - go mod download
    - golangci-lint run --timeout=5m
  cache:
    paths:
      - vendor/
  rules:
    - changes:
        - "**/*.go"
        - "go.mod"
        - "go.sum"

# Unit tests
test:
  stage: test
  image: golang:${GO_VERSION}-alpine
  services:
    - name: postgres:16-alpine
      alias: db
  variables:
    DB_URL: "postgres://postgres:postgres@db:5432/testdb?sslmode=disable"
  script:
    - go mod download
    - go test -v -race -coverprofile=coverage.out ./...
    - go tool cover -html=coverage.out -o coverage.html
  artifacts:
    paths:
      - coverage.html
    reports:
      coverage_report:
        coverage_format: cobertura
        path: coverage.out
  cache:
    paths:
      - vendor/

# Build static binary
build:
  stage: build
  image: golang:${GO_VERSION}-alpine
  script:
    - go mod download
    - CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -s' -o bin/myapp ./cmd/api/main.go
  artifacts:
    paths:
      - bin/
  cache:
    paths:
      - vendor/
  rules:
    - if: $CI_COMMIT_BRANCH == "main"

# Docker build and push (only on main)
docker-build:
  stage: deploy
  image: docker:27
  services:
    - docker:dind
  variables:
    DOCKER_TLS_CERTDIR: "/certs"
  script:
    - docker login -u $CI_REGISTRY_USER -p $CI_REGISTRY_PASSWORD $CI_REGISTRY
    - docker build -t $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA .
    - docker push $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA
  rules:
    - if: $CI_COMMIT_BRANCH == "main"
```
- **Run**: Pipeline triggers on push/MR; artifacts include binaries/coverage.
- **Tips**: Use official `Go.gitlab-ci.yml` template as base (`include: - template: Go.gitlab-ci.yml`).

### Advanced Pipeline with Includes and Matrix
For a multi-service Go project with lint/test/build/deploy stages, using includes.
```yaml
# .gitlab-ci.yml (main)
include:
  - local: 'ci/templates/lint.yml'
  - local: 'ci/templates/test.yml'
  - local: 'ci/templates/build.yml'
  - local: 'ci/templates/deploy.yml'

# Override for production
deploy-prod:
  extends: .deploy-job
  variables:
    ENVIRONMENT: production
  rules:
    - if: $CI_COMMIT_BRANCH == "main"
      when: manual
```

**ci/templates/test.yml** (example include):
```yaml
.test-job:
  stage: test
  image: golang:1.21-alpine
  script:
    - go mod download
    - go test -v -race ./... -coverprofile=coverage.out
  artifacts:
    reports:
      coverage_report:
        coverage_format: cobertura
        path: coverage.out
  parallel:
    matrix:
      - GO_TEST_TAG: [unit, integration]
```

- **Benefits**: Reusable, versioned templates (pin with `ref: v1.0`); matrix for parallel tests.

## Integration with Go Tools

- **Viper/Cobra**: Set vars like `DB_URL` in `.gitlab-ci.yml`; load in Go code.
- **Gin/JWT/Zap**: Test endpoints with `go test`; artifacts for coverage.
- **pgx/Mongo/Redis/MinIO**: Use services in jobs (e.g., `services: - postgres:16`); test connectivity.
- **Prometheus**: Expose metrics in build job; scrape in deploy.
- **Retry/Circuit Breaker/Rate Limit**: Test resilience in integration stage.
- **OpenTelemetry**: Add tracing in test jobs; export to GitLab's built-in tracing.

## Testing and CI

- **Lint Pipeline**: Use `golangci-lint` config in `.golangci.yml`.
- **Local Testing**: `gitlab-ci-local --file .gitlab-ci.yml lint` or `dry-run`.
- **CI/CD**: Trigger on push/MR; manual deploys on main. Use `workflow:rules` to skip drafts.

## Additional Resources

- Official YAML Reference: https://docs.gitlab.com/ee/ci/yaml/
- Optimization Guide: https://docs.gitlab.com/ee/ci/yaml/yaml_optimization/
- Go Template: https://docs.gitlab.com/ee/ci/examples/#go
- Community: Reddit r/gitlab Best Practices

Adhere to these practices for robust GitLab CI/CD pipelines in Go projects. Update as GitLab features evolve.