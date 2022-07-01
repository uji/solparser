package solparser

import (
	"errors"

	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/lexer"
)

func (p *Parser) ParseContractDefinition() (*ast.ContractDefinition, error) {
	var abstract bool
	p.lexer.Scan()
	tkn := p.lexer.Token()
	if tkn.TokenType == lexer.Abstract {
		abstract = true
		p.lexer.Scan()
	}
	if tkn.TokenType != lexer.Contract {
		return nil, errors.New("not found contract definition")
	}

	p.lexer.Scan()
	if p.lexer.Token().TokenType != lexer.Unknown {
		return nil, errors.New("not found left brace")
	}

	if p.lexer.Token().TokenType != lexer.BraceL {
		return nil, errors.New("not found left brace")
	}

	p.lexer.Scan()
	if p.lexer.Token().TokenType != lexer.BraceR {
		return nil, errors.New("not found right brace")
	}

	return &ast.ContractDefinition{
		Abstract: abstract,
	}, nil
}
