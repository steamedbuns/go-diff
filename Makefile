# go-diff Makefile
#
# Run `make help` to list targets.

BINARY      := go-diff
CMD_PKG     := ./cmd/$(BINARY)
BIN_DIR     := bin
COVER_FILE  := coverage.out

# Version info injected at link time. Declare `var version = "dev"` in
# package main to make use of it; -X is a no-op if the variable doesn't exist.
VERSION     ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
LDFLAGS     := -s -w -X main.version=$(VERSION)

GO          ?= go
GOFLAGS     ?=
RUN			?= .

.DEFAULT_GOAL := help

.PHONY: help
help: ## Show this help
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z_-]+:.*##/ {printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: build
build: ## Build the binary into ./bin
	$(GO) build $(GOFLAGS) -trimpath -ldflags '$(LDFLAGS)' -o $(BIN_DIR)/$(BINARY) $(CMD_PKG)

.PHONY: install
install: ## Install the binary into $GOBIN (or $GOPATH/bin)
	$(GO) install $(GOFLAGS) -trimpath -ldflags '$(LDFLAGS)' $(CMD_PKG)

.PHONY: run
run: build ## Build and run; pass arguments with ARGS="a.txt b.txt"
	./$(BIN_DIR)/$(BINARY) $(ARGS)

.PHONY: test
test: ## Run unit tests with the race detector
	$(GO) test $(GOFLAGS) -race -count=1 -run "$(RUN)" ./...

.PHONY: cover
cover: ## Run tests with coverage and print a per-function summary
	$(GO) test $(GOFLAGS) -race -covermode=atomic -coverprofile=$(COVER_FILE) ./...
	$(GO) tool cover -func=$(COVER_FILE)

.PHONY: cover-html
cover-html: cover ## Open the coverage report in a browser
	$(GO) tool cover -html=$(COVER_FILE)

.PHONY: bench
bench: ## Run benchmarks
	$(GO) test $(GOFLAGS) -run=^$$ -bench=. -benchmem ./...

.PHONY: fmt
fmt: ## Format all Go source
	$(GO) fmt ./...

.PHONY: vet
vet: ## Run go vet
	$(GO) vet ./...

.PHONY: lint
lint: ## Run golangci-lint (if installed)
	@command -v golangci-lint >/dev/null 2>&1 || { echo "golangci-lint not installed: https://golangci-lint.run/welcome/install/"; exit 1; }
	golangci-lint run ./...

.PHONY: tidy
tidy: ## Tidy and verify go.mod / go.sum
	$(GO) mod tidy
	$(GO) mod verify

.PHONY: check
check: fmt vet test ## Format, vet and test (run before committing)

.PHONY: clean
clean: ## Remove build and coverage artifacts
	rm -rf $(BIN_DIR) $(COVER_FILE)
	$(GO) clean -testcache
