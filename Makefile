.PHONY: dev down test lint generate

dev:
	docker compose up --build

down:
	docker compose down -v

test:
	cd backend && go test ./...
	cd frontend && npm test

lint:
	cd backend && golangci-lint run ./...
	cd frontend && npm run lint

generate:
	cd backend && go generate ./...
	cd frontend && npm run codegen
