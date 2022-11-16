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
		return nil, token.NewPosError(tkn.Pos, "not found pragma")
	}

	pragmaName, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if pragmaName.TokenType != token.Unknown {
		return nil, token.NewPosError(pragmaName.Pos, "not found pragma name")
	}

	expression, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if expression.TokenType != token.BitXor {
		return nil, token.NewPosError(expression.Pos, "not found BitXor expression")
	}

	version, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if version.TokenType != token.Unknown {
		return nil, token.NewPosError(version.Pos, "not found version")
	}

	tkn, err = p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if tkn.TokenType != token.Semicolon {
		return nil, token.NewPosError(tkn.Pos, "not found Semicolon")
	}

	return &ast.PragmaDirective{
		PragmaName: pragmaName.Text,
		PragmaValue: ast.PragmaValue{
			Version:    version.Text,
			Expression: expression.Text,
		},
	}, nil
}
