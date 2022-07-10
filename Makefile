.PHONY: help

help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

antlr-image: ## Build antlr/antlr4 docker image
	@docker build -t antlr/antlr4 ./antlr4/docker/

parser: ## Generate parser file by antlr
	@mkdir ./parser/
	@cp ./antlr-grammars-v4/solidity/Solidity.g4 ./parser/
	@docker run --rm -v `pwd`:/work antlr/antlr4 -Dlanguage=Go ./parser/Solidity.g4
