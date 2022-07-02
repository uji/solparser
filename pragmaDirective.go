package solparser

import (
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func (p *Parser) ParsePragmaDirective() (*ast.PragmaDirective, error) {
	p.lexer.Scan()
	if tkn := p.lexer.Token(); tkn.TokenType != token.Pragma {
		return nil, newError(tkn.Pos, "not found pragma")
	}

	p.lexer.Scan()
	pragmaName := p.lexer.Token()
	if pragmaName.TokenType != token.Unknown {
		return nil, newError(pragmaName.Pos, "not found pragma name")
	}

	p.lexer.Scan()
	expression := p.lexer.Token()
	if expression.TokenType != token.Hat {
		return nil, newError(expression.Pos, "not found Hat expression")
	}

	p.lexer.Scan()
	version := p.lexer.Token()
	if version.TokenType != token.Unknown {
		return nil, newError(version.Pos, "not found version")
	}

	p.lexer.Scan()
	if tkn := p.lexer.Token(); tkn.TokenType != token.Semicolon {
		return nil, newError(tkn.Pos, "not found Semicolon")
	}

	return &ast.PragmaDirective{
		PragmaName: pragmaName.Text,
		PragmaValue: ast.PragmaValue{
			Version:    version.Text,
			Expression: expression.Text,
		},
	}, nil
}
