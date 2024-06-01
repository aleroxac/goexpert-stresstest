.PHONY: help
help: ## Show this menu
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)


.PHONY: build
build: ## Build the container image
	@docker build -t aleroxac/goexpert-stresstest:v1 .

.PHONY: run
run: ## Run the code locally
	@go run cmd/app/main.go --url=https://google.com --requests=10 --concurrency=2


.PHONY: runc
runc: ## Run the app via docker
	@docker run aleroxac/goexpert-stresstest:v1 --url=https://google.com --requests=10 --concurrency=2
