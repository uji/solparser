.PHONY: help submodules

help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

submodules: ## Init submodules
	@git submodule init
	@git submodule update
	@# antlr-grammers-v4
	@git -C antlr-grammars-v4 config core.sparsecheckout true
	@echo '/solidity/' > ./.git/modules/antlr-grammars-v4/info/sparse-checkout
	@git -C antlr-grammars-v4 read-tree -mu HEAD
	@# antlr4
	@git -C antlr4 config core.sparsecheckout true
	@echo '/docker/' > ./.git/modules/antlr4/info/sparse-checkout
	@git -C antlr4 read-tree -mu HEAD

antlr-image: ## Build antlr/antlr4 docker image
	@docker build -t antlr/antlr4 ./antlr4/docker/

parser: ## Generate parser file by antlr
	@mkdir ./parser/
	@cp ./antlr-grammars-v4/solidity/Solidity.g4 ./parser/
	@docker run --rm -v `pwd`:/work antlr/antlr4 -Dlanguage=Go ./parser/Solidity.g4
