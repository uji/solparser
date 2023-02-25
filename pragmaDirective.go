package solparser

import (
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func (p *Parser) ParsePragmaDirective() (*ast.PragmaDirective, error) {
	prgm, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if prgm.TokenType != token.Pragma {
		return nil, token.NewPosError(prgm.Pos, "not found pragma.")
	}

	tkns := make([]*token.Token, 0, 1)
	for {
		tkn, err := p.lexer.Scan()
		if err != nil {
			return nil, err
		}
		if tkn.TokenType == token.EOS {
			return nil, token.NewPosError(tkn.Pos, "not found Semicolon.")
		}
		if tkn.TokenType == token.Semicolon {
			if len(tkns) == 0 {
				return nil, token.NewPosError(tkn.Pos, "not found pragma tokens.")
			}
			return &ast.PragmaDirective{
				Pragma:       prgm.Pos,
				PragmaTokens: tkns,
				Semicolon:    tkn.Pos,
			}, nil
		}
		tkns = append(tkns, &tkn)
	}
}
