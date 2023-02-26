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
	if prgm.Type != token.Pragma {
		return nil, token.NewPosError(prgm.Position, "not found pragma.")
	}

	tkns := make([]*token.Token, 0, 1)
	for {
		tkn, err := p.lexer.Scan()
		if err != nil {
			return nil, err
		}
		if tkn.Type == token.EOS {
			return nil, token.NewPosError(tkn.Position, "not found Semicolon.")
		}
		if tkn.Type == token.Semicolon {
			if len(tkns) == 0 {
				return nil, token.NewPosError(tkn.Position, "not found pragma tokens.")
			}
			return &ast.PragmaDirective{
				Pragma:       prgm.Position,
				PragmaTokens: tkns,
				Semicolon:    tkn.Position,
			}, nil
		}
		tkns = append(tkns, &tkn)
	}
}
