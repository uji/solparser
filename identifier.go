package solparser

import (
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func (p *Parser) ParseIdentifier() (ast.Identifier, error) {
	tkn, err := p.lexer.Scan()
	if err != nil {
		return token.Token{}, err
	}

	switch tkn.TokenType {
	case token.Identifier, token.From, token.Error, token.Revert, token.Global:
		return tkn, nil
	}

	return token.Token{}, token.NewPosError(tkn.Pos, "keyword is not available as identifier.")
}
