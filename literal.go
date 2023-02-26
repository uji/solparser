package solparser

import (
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func (p *Parser) ParseLiteral() (ast.Literal, error) {
	tkn, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}

	if tkn.Type != token.StringLiteral {
		return nil, token.NewPosError(tkn.Position, "not found string literal quote")
	}

	lit := ast.StringLiteral(tkn)
	return &lit, nil
}
