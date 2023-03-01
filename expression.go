package solparser

import (
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func (p *Parser) ParseExpression() (ast.Expression, error) {
	tkn, err := p.lexer.Peek()
	if err != nil {
		return nil, err
	}

	if tkn.Type == token.NonEmptyStringLiteral {
		return p.ParseLiteral()
	}

	return nil, token.NewPosError(tkn.Position, "not found expression.")
}
