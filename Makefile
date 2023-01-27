run: build
	@./bin/gobank

build:
	@go build -o bin/gobank

test:
	@go test -v ./...