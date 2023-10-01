package solparser

import (
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func (p *Parser) ParseIdentifierPath() (ast.IdentifierPath, error) {
	id, err := p.ParseIdentifier()
	if err != nil {
		return ast.IdentifierPath{}, err
	}

	var elements []*ast.IdentifierPathElement
	for {
		prd, err := p.lexer.Peek()
		if err != nil || prd.Type != token.Period {
			break
		}
		p.lexer.Scan()

		elements = append(elements, &ast.IdentifierPathElement{
			Identifier: ast.Identifier(id),
			Period:     &prd.Position,
		})

		i, err := p.ParseIdentifier()
		if err != nil {
			return ast.IdentifierPath{}, err
		}
		id = i
	}

	elements = append(elements, &ast.IdentifierPathElement{
		Identifier: ast.Identifier(id),
	})

	return ast.IdentifierPath{
		Elements: elements,
	}, nil
}
