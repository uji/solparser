package solparser

import (
	"errors"

	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/lexer"
)

func (p *Parser) ParseContractDefinition() (*ast.ContractDefinition, error) {
	p.lexer.Scan()
	keyward := p.lexer.Token()
	if keyward.TokenType != lexer.Contract {
		return nil, errors.New("not found contract definition")
	}

	p.lexer.Scan()
	if p.lexer.Token().TokenType != lexer.BraceL {
		return nil, errors.New("not found left brace")
	}

	part, err := p.ParseContractPart()
	if err != nil {
		return nil, err
	}

	p.lexer.Scan()
	if p.lexer.Token().TokenType != lexer.BraceR {
		return nil, errors.New("not found right brace")
	}

	return &ast.ContractDefinition{
		ContractPart: part,
	}, nil
}

func (p *Parser) ParseContractPart() (*ast.ContractPart, error) {
	funcDef, err := p.ParseFunctionDefinition()
	if err != nil {
		return nil, err
	}

	return &ast.ContractPart{
		FunctionDefinition: funcDef,
	}, nil
}
