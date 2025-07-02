.PHONY: run tidy linter dev

## Run the server
run:
	go run cmd/server/main.go

## Ensure Go modules are up to date
tidy:
	go mod tidy

## Run the linter
linter:
	golangci-lint run ./...

## Tidy modules, run linter, and start the server
dev: tidy linter run
