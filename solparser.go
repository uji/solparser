package solparser

import (
	"io"

	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/lexer"
)

type Parser struct {
	input io.Reader
	lexer *lexer.Lexer
}

func New(input io.Reader) *Parser {
	return &Parser{
		input: input,
		lexer: lexer.New(input),
	}
}

func (p *Parser) Parse(input io.Reader) (*ast.SourceUnit, error) {
	if !p.lexer.Peek() {
		return nil, nil
	}

	if err := p.lexer.PeekError(); err != nil {
		return nil, err
	}

	var pragmaDirective *ast.PragmaDirective
	var contractDefinition *ast.ContractDefinition
	var functionDefinition *ast.FunctionDefinition

	switch p.lexer.PeekToken().TokenType {
	case lexer.Pragma:
		prgm, err := p.ParsePragmaDirective()
		if err != nil {
			return nil, err
		}
		pragmaDirective = prgm
	case lexer.Abstract, lexer.Contract:
		cntrct, err := p.ParseContractDefinition()
		if err != nil {
			return nil, err
		}
		contractDefinition = cntrct
	case lexer.Function:
		fnc, err := p.ParseFunctionDefinition()
		if err != nil {
			return nil, err
		}
		functionDefinition = fnc
	}

	return &ast.SourceUnit{
		PragmaDirective:    pragmaDirective,
		ContractDefinition: contractDefinition,
		FunctionDefinition: functionDefinition,
	}, nil
}
