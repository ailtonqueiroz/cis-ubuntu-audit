.PHONY: build test audit

build:
	@echo "Building CLI..."
	@cd cmd && go build -o ../bin/cis-audit

test:
	@echo "Running tests..."
	@go test ./...

audit:
	@echo "Running CIS audit..."
	@./examples/auto-fix.sh
