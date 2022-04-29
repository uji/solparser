antlr4go:
	docker build -t antlr/antlr4 --platform linux/amd64 ./antlr4/docker/
	docker run --rm -v `pwd`:/work ./antlr4/docker/Dockerfile -Dlanguage=Go ./solidity-antlr/Solidity.g4
