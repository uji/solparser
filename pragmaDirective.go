package solparser

import (
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func (p *Parser) ParsePragmaDirective() (*ast.PragmaDirective, error) {
	tkn, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if tkn.TokenType != token.Pragma {
		return nil, newError(tkn.Pos, "not found pragma")
	}

	pragmaName, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if pragmaName.TokenType != token.Unknown {
		return nil, newError(pragmaName.Pos, "not found pragma name")
	}

	expression, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if expression.TokenType != token.BitXor {
		return nil, newError(expression.Pos, "not found BitXor expression")
	}

	version, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if version.TokenType != token.Unknown {
		return nil, newError(version.Pos, "not found version")
	}

	tkn, err = p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if tkn.TokenType != token.Semicolon {
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
