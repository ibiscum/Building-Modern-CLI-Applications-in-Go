.PHONY: all lint

all: lint

lint:
	go list -f '{{.Dir}}/...' -m | xargs golangci-lint run