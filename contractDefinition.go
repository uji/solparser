package solparser

import (
	"errors"

	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func (p *Parser) ParseContractDefinition() (*ast.ContractDefinition, error) {
	var abstract bool
	tkn, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if tkn.TokenType == token.Abstract {
		abstract = true
		p.lexer.Scan()
	}
	if tkn.TokenType != token.Contract {
		return nil, errors.New("not found contract definition")
	}

	tkn, err = p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if tkn.TokenType != token.Identifier {
		return nil, errors.New("not found left brace")
	}

	tkn, err = p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if tkn.TokenType != token.LBrace {
		return nil, errors.New("not found left brace")
	}

	tkn, err = p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if tkn.TokenType != token.RBrace {
		return nil, errors.New("not found right brace")
	}

	return &ast.ContractDefinition{
		Abstract: abstract,
	}, nil
}
