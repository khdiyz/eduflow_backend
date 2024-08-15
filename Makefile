tidy:
	@go mod tidy && go mod vendor

run:
	@swag init -g cmd/main.go > /dev/null && go run cmd/main.go

up:
	@go run cmd/migration/main.go up

down:
	@go run cmd/migration/main.go down

redo:
	@go run cmd/migration/main.go redo

status:
	@go run cmd/migration/main.go status