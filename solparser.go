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

	contractDefinition, err := p.parseContractDefinition()
	if err != nil {
		return nil, err
	}

	return &ast.SourceUnit{
		PragmaDirective:    pragmaDirective,
		ContractDefinition: contractDefinition,
	}, nil
}

func (p *Parser) ParsePragmaDirective() (*ast.PragmaDirective, error) {
	p.lexer.Scan()
	if p.lexer.Token().TokenType != lexer.Pragma {
		return nil, errors.New("not found pragma")
	}

	p.lexer.Scan()
	pragmaName := p.lexer.Token()
	if pragmaName.TokenType != lexer.Unknown {
		return nil, errors.New("not found unkknown")
	}

	p.lexer.Scan()
	expression := p.lexer.Token()
	if expression.TokenType != lexer.Hat {
		return nil, errors.New("not found expression")
	}

	p.lexer.Scan()
	version := p.lexer.Token()
	if version.TokenType != lexer.Unknown {
		return nil, errors.New("not found version")
	}

	p.lexer.Scan()
	if p.lexer.Token().TokenType != lexer.Semicolon {
		return nil, errors.New("not found semicolon")
	}

	// pragma ~ のパース
	return &ast.PragmaDirective{
		PragmaName: pragmaName.Text,
		PragmaValue: ast.PragmaValue{
			Version:    version.Text,
			Expression: expression.Text,
		},
	}, nil
}

func (p *Parser) parseContractDefinition() (*ast.ContractDefinition, error) {
	// contract ~ のパース
	return nil, nil
}
