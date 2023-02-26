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

	// peek to find literal end position.
	peek, err := p.lexer.Peek()
	if err != nil {
		return nil, err
	}

	return &ast.StringLiteral{
		Value: tkn.Value,
		From:  tkn.Position,
		To: token.Pos{
			Column: peek.Position.Column - 1,
			Line:   peek.Position.Line,
		},
	}, nil
}
