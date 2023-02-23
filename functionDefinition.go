package solparser

import (
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func (p *Parser) ParseVisibility() (ast.Visibility, error) {
	tkn, err := p.lexer.Scan()
	if err != nil {
		return token.Token{}, err
	}

	switch tkn.TokenType {
	case token.Internal, token.External, token.Public, token.Private:
		return tkn, nil
	}

	return token.Token{}, token.NewPosError(tkn.Pos, "not found visibility keyword.")
}

func (p *Parser) ParseStateMutability() (ast.StateMutability, error) {
	tkn, err := p.lexer.Scan()
	if err != nil {
		return token.Token{}, err
	}

	switch tkn.TokenType {
	case token.Pure, token.View, token.Payable:
		return tkn, nil
	}

	return token.Token{}, token.NewPosError(tkn.Pos, "not found state-mutability keyword.")
}
