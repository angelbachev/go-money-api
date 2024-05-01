##@ Go Money API
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

build: ## Build project
	@go build -o bin/go-money ./cmd/go-money/main.go

run: build ## Run web server
	@./bin/go-money

docker: ## Start docker containers
	docker-compose up -d

test:
	@go test -v ./...


migrate-db: ## Execute db migrations
	@migrate -database "mysql://root:secret@tcp(localhost:3346)/go_money?multiStatements=true&x-tls-insecure-skip-verify=false" -path "db/migrations" up

revert-db: ## Revert db migrations
	@migrate -database "mysql://root:secret@tcp(localhost:3346)/go_money?multiStatements=true&x-tls-insecure-skip-verify=false" -path "db/migrations" down


build-seed: ## Build seed
	@go build -o bin/seed ./cmd/seed/main.go