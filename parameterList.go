package solparser

import (
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func (p *Parser) ParseParameter() (*ast.Parameter, error) {
	tn, err := p.ParseTypeName()
	if err != nil {
		return nil, err
	}

	return &ast.Parameter{
		TypeName: tn,
	}, nil
}

func (p *Parser) ParseParameterList() (ast.ParameterList, error) {
	prms := make(ast.ParameterList, 0, 1)
	for {
		prm, err := p.ParseParameter()
		if err != nil {
			return nil, err
		}

		prms = append(prms, prm)

		comma, err := p.lexer.Peek()
		if err != nil {
			return nil, err
		}

		if comma.TokenType != token.Comma {
			return prms, nil
		}

		p.lexer.Scan()
	}
}
