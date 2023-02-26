package solparser

import (
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func (p *Parser) ParseTypeName() (ast.TypeName, error) {
	tkn, err := p.lexer.Peek()
	if err != nil {
		return nil, err
	}

	switch tkn.Type {
	case token.Address, token.String, token.Bytes, token.Fixed, token.Bool:
		return p.ParseElementaryTypeName()
	}

	return nil, token.NewPosError(tkn.Position, "not found type-name.")
}

func (p *Parser) ParseElementaryTypeName() (ast.TypeName, error) {
	tkn, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}

	if tkn.Type == token.Address {
		tkn2, err := p.lexer.Peek()
		if err != nil {
			return nil, err
		}

		if tkn2.Type == token.Payable {
			p.lexer.Scan()
			return ast.ElementaryTypeName{&tkn, &tkn2}, nil
		}
	}

	switch tkn.Type {
	case token.Address, token.String, token.Bytes, token.Fixed, token.Bool:
		return ast.ElementaryTypeName{&tkn}, nil
	}

	return nil, token.NewPosError(tkn.Position, "not found elementary type name keyword.")
}
