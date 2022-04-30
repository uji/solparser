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
	}
}

func (p *Parser) Parse(input io.Reader) (*ast.SourceUnit, error) {
	pragmaDirective, err := p.parsePragmaDirective()
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

func (p *Parser) parsePragmaDirective() (*ast.PragmaDirective, error) {
	// pragma ~ のパース
	return nil, nil
}

func (p *Parser) parseContractDefinition() (*ast.ContractDefinition, error) {
	// contract ~ のパース
	return nil, nil
}
