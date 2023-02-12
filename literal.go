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

	return &ast.StringLiteral{
		Str:      tkn.Text,
		Position: tkn.Pos,
		EndPos: token.Pos{
			Column: tkn.Pos.Column + len(tkn.Text) - 1,
			Line:   tkn.Pos.Line, // TODO: support new line code
		},
	}, nil
}
