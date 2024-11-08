tidy:
	@go mod tidy

run:
	@go run cmd/app/main.go

up:
	@go run cmd/migration/main.go up

down:
	@go run cmd/migration/main.go down

redo:
	@go run cmd/migration/main.go redo

status:
	@go run cmd/migration/main.go status

swag:
	@swag init -g cmd/app/main.go