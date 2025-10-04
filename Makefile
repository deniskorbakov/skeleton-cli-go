# use linter for formatted code
lint:
	docker run -t --rm -v $$(pwd):/app -w /app golangci/golangci-lint:v2.5.0 golangci-lint run
build:
	go mod vendor
	go build -ldflags "-w -s" -o cli cmd/cli/main.go