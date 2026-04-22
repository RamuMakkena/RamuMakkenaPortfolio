# MakkenaRamu.com — Portfolio

Full-stack portfolio site built as a Senior SDE practice playground.

| Layer | Stack |
|---|---|
| Backend | Go + gqlgen + sqlc + Postgres |
| Frontend | React + Vite + TypeScript + Tailwind |
| Infra | Terraform → AWS (ECS Fargate + RDS + CloudFront) |
| Observability | OpenTelemetry + Jaeger + Prometheus + slog |
| CI/CD | GitHub Actions |

## Quick start
```bash
make dev       # starts everything via docker-compose
make test      # run all tests
make lint      # run all linters
```
EOF