patch:
	@go fmt ./...
	@go vet ./...

run:
	@go run main.go

tooling:
	@go install github.com/jackc/tern@latest
	@go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

build: patch
	@GOOS=linux GOARCH=amd64 go build -o bin/core-linux-amd64 main.go
	@GOOS=linux GOARCH=arm64 go build -o bin/core-linux-arm64 main.go
