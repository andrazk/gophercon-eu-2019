APP        := tenerife
GIT_COMMIT := $(shell git rev-parse --short HEAD | sed -E 's/[^a-zA-Z0-9]+/-/g')
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
BUILD_TIME := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS    := -ldflags " \
	-X $(APP)/internal/diagnostics.revision=$(GIT_COMMIT) \
	-X $(APP)/internal/diagnostics.buildTime=$(BUILD_TIME) \
	-X $(APP)/internal/diagnostics.branch=$(GIT_BRANCH) \
"

.DEFAULT_GOAL := help
.PHONY: help
help: ## Show help
	@echo "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:"
	@grep -E '^[a-zA-Z_/%\-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2}'

build: ## Build an executable
	CGO_ENABLED=0 go build \
		$(LDFLAGS) \
		-mod=vendor \
		-o ./bin/$(APP) \
		$(APP)/cmd/$(APP)

run: ## Run the executable
	./bin/$(APP)

test: ## Run tests
	go test -v -race -cover ./...

docker/build: ## Build deployable container
	docker build -t $(APP) .

docker/run: ## Run deployable container
	docker run \
		--rm -ti \
		-p 8080:8080 \
		-p 9090:9090 \
		$(APP)
