run:
	go run cmd/server/main.go

linter:
	@golangci-lint run ./...

dev: linter run
