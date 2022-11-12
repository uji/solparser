.PHONY: help coverage
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

coverage: ## generate test coverage profile to cover.html
	go test -cover -coverprofile=cover.out ./...
	go tool cover -html=cover.out -o cover.html
