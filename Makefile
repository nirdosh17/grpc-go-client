.DEFAULT_GOAL=help

build: ## downloads dependencies and create Go binary for the client
	go mod download
	go build -o cmd/main client/*.go

run: build ## runs Go binary
	./cmd/main

help:
	@grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
