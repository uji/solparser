package solparser

import (
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func (p *Parser) ParseInheritanceSpecifire() (*ast.InheritanceSpecifier, error) {
	ip, err := p.ParseIdentifierPath()
	if err != nil {
		return nil, err
	}

	lparen, err := p.lexer.Peek()
	if err != nil {
		return nil, err
	}
	if lparen.Type != token.LParen {
		return &ast.InheritanceSpecifier{
			IdentifierPath: ip,
		}, nil
	}

	cal, err := p.ParseCallArgumentList()
	if err != nil {
		return nil, err
	}

	return &ast.InheritanceSpecifier{
		IdentifierPath:   ip,
		CallArgumentList: cal,
	}, nil
}
