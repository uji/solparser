package solparser

import (
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func (p *Parser) ParseIdentifierPath() (ast.IdentifierPath, error) {
	id, err := p.lexer.Peek()
	if err != nil {
		return ast.IdentifierPath{}, err
	}
	if !isIdentifier(id) {
		return ast.IdentifierPath{}, token.NewPosError(id.Position, "not found identifier.")
	}
	p.lexer.Scan()

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

		i, err := p.lexer.Peek()
		if err != nil {
			return ast.IdentifierPath{}, err
		}
		if !isIdentifier(i) {
			return ast.IdentifierPath{}, token.NewPosError(i.Position, "not found identifier.")
		}
		p.lexer.Scan()
		id = i
	}

	elements = append(elements, &ast.IdentifierPathElement{
		Identifier: ast.Identifier(id),
	})

	return ast.IdentifierPath{
		Elements: elements,
	}, nil
}
