.DEFAULT_GOAL=help

build: ## downloads dependencies and create Go binary for the client
	go mod download
	go build -o cmd/main client/*.go

run: build ## run client
	./cmd/main

run-with-tls: fetch-ca-cert ## run client with TLS enabled
	TLS=true make run

fetch-ca-cert: ## fetches CA cert from server repository in Github
	curl -s https://raw.githubusercontent.com/nirdosh17/grpc-go-service/main/ssl/ca.crt --create-dirs -o ssl/ca.crt

help:
	@grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
