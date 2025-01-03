.PHONY: default run swagger

default: run-api

run-api:
	@go run cmd/api/main.go
swagger:
	@swag init -d cmd/api,internal/handler
	