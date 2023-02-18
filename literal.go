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

	if tkn.TokenType != token.StringLiteral {
		return nil, token.NewPosError(tkn.Pos, "not found string literal quote")
	}

	// peek to find literal end position.
	peek, err := p.lexer.Peek()
	if err != nil {
		return nil, err
	}

	return &ast.StringLiteral{
		Str:      tkn.Text,
		Position: tkn.Pos,
		EndPos: token.Pos{
			Column: peek.Pos.Column - 1,
			Line:   peek.Pos.Line,
		},
	}, nil
}
