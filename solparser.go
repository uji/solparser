package solparser

import (
	"errors"
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
	pragmaDirective, err := p.ParsePragmaDirective()
	if err != nil {
		return nil, err
	}

	contractDefinition, err := p.ParseContractDefinition()
	if err != nil {
		return nil, err
	}

	return &ast.SourceUnit{
		PragmaDirective:    pragmaDirective,
		ContractDefinition: contractDefinition,
	}, nil
}

func (p *Parser) ParseModirierList() (*ast.ModifierList, error) {
	return nil, nil
}

func (p *Parser) ParseReturnParameters() (*ast.ReturnParameters, error) {
	return nil, nil
}

func (p *Parser) ParseFunctionDefinition() (*ast.FunctionDefinition, error) {
	return nil, nil
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
