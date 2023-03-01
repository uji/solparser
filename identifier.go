package solparser

import (
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func (p *Parser) ParseIdentifier() (ast.Identifier, error) {
	tkn, err := p.lexer.Scan()
	if err != nil {
		return ast.Identifier{}, err
	}

	switch tkn.Type {
	case token.Identifier, token.From, token.Error, token.Revert, token.Global:
		return ast.Identifier(tkn), nil
	}

	return ast.Identifier{}, token.NewPosError(tkn.Position, "keyword is not available as identifier.")
}
