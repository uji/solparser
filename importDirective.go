package solparser

import (
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func (p *Parser) ParsePath() (ast.Path, error) {
	lit, err := p.lexer.Peek()
	if err != nil {
		return ast.Path{}, err
	}
	if lit.Type != token.NonEmptyStringLiteral {
		return ast.Path{}, token.NewPosError(lit.Position, "not found non-empty-string-literal.")
	}
	if _, err := p.lexer.Scan(); err != nil {
		return ast.Path{}, err
	}
	return ast.Path(lit), nil
}

func (p *Parser) ParseImportDirectivePathElement() (*ast.ImportDirectivePathElement, error) {
	path, err := p.ParsePath()
	if err != nil {
		return nil, err
	}

	as, err := p.lexer.Peek()
	if err != nil {
		return nil, err
	}
	if as.Type != token.As {
		return &ast.ImportDirectivePathElement{Path: path}, nil
	}
	if _, err := p.lexer.Scan(); err != nil {
		return nil, err
	}

	id, err := p.ParseIdentifier()
	if err != nil {
		return nil, err
	}

	return &ast.ImportDirectivePathElement{
		Path:       path,
		As:         &as.Position,
		Identifier: &id,
	}, nil
}

func (p *Parser) ParseImportDirectiveSymbolAliasesElement() (*ast.ImportDirectiveSymbolAliasesElement, error) {
	aliases, err := p.ParseSymbolAliases()
	if err != nil {
		return nil, err
	}

	from, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if from.Type != token.From {
		return nil, token.NewPosError(from.Position, "not found from keyword.")
	}

	path, err := p.ParsePath()
	if err != nil {
		return nil, err
	}

	return &ast.ImportDirectiveSymbolAliasesElement{
		SymbolAliases: aliases,
		From:          from.Position,
		Path:          path,
	}, nil
}

func (p *Parser) ParseImportDirectiveMulElement() (*ast.ImportDirectiveMulElement, error) {
	mul, err := p.lexer.Peek()
	if err != nil {
		return nil, err
	}
	if mul.Type != token.Mul {
		return nil, token.NewPosError(mul.Position, "not found mul keyword.")
	}
	if _, err := p.lexer.Scan(); err != nil {
		return nil, err
	}

	as, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if as.Type != token.As {
		return nil, token.NewPosError(as.Position, "not found as keyword.")
	}

	id, err := p.ParseIdentifier()
	if err != nil {
		return nil, err
	}

	from, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if from.Type != token.From {
		return nil, token.NewPosError(from.Position, "not found from keyword.")
	}

	path, err := p.ParsePath()
	if err != nil {
		return nil, err
	}

	return &ast.ImportDirectiveMulElement{
		Mul:        mul.Position,
		As:         as.Position,
		Identifier: id,
		From:       from.Position,
		Path:       path,
	}, nil
}

func (p *Parser) ParseImportDirective() (*ast.ImportDirective, error) {
	impt, err := p.lexer.Peek()
	if err != nil {
		return nil, err
	}
	if impt.Type != token.Import {
		return nil, token.NewPosError(impt.Position, "not found import keyword.")
	}
	if _, err := p.lexer.Scan(); err != nil {
		return nil, err
	}

	tkn, err := p.lexer.Peek()
	if err != nil {
		return nil, err
	}

	var el ast.ImportDirectiveElement
	switch tkn.Type {
	case token.NonEmptyStringLiteral:
		el, err = p.ParseImportDirectivePathElement()
	case token.LBrace:
		el, err = p.ParseImportDirectiveSymbolAliasesElement()
	case token.Mul:
		el, err = p.ParseImportDirectiveMulElement()
	default:
		return nil, token.NewPosError(tkn.Position, "not found import-directive element.")
	}
	if err != nil {
		return nil, err
	}

	semi, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if semi.Type != token.Semicolon {
		return nil, token.NewPosError(semi.Position, "not found semicolon.")
	}

	return &ast.ImportDirective{
		Import:    impt.Position,
		Element:   el,
		Semicolon: semi.Position,
	}, nil
}
