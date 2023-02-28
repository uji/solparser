package solparser

import (
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func (p *Parser) ParseSymbolAlias() (*ast.SymbolAlias, error) {
	smbl, err := p.lexer.Peek()
	if err != nil {
		return nil, err
	}

	if smbl.Type != token.Identifier {
		return nil, token.NewPosError(smbl.Position, "not found identifier.")
	}

	if _, err := p.lexer.Scan(); err != nil {
		return nil, err
	}

	as, err := p.lexer.Peek()
	if err != nil {
		return nil, err
	}
	if as.Type != token.As {
		return &ast.SymbolAlias{Symbol: smbl}, nil
	}
	if _, err := p.lexer.Scan(); err != nil {
		return nil, err
	}

	alias, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}

	if alias.Type != token.Identifier {
		return nil, token.NewPosError(alias.Position, "not found identifier.")
	}

	return &ast.SymbolAlias{
		Symbol: smbl,
		As:     &as.Position,
		Alias:  &alias,
	}, nil
}

func (p *Parser) ParseSymbolAliases() (*ast.SymbolAliases, error) {
	lbrace, err := p.lexer.Peek()
	if err != nil {
		return nil, err
	}

	if lbrace.Type != token.LBrace {
		return nil, token.NewPosError(lbrace.Position, "not found LBrace.")
	}

	if _, err := p.lexer.Scan(); err != nil {
		return nil, err
	}

	alias, err := p.ParseSymbolAlias()
	if err != nil {
		return nil, err
	}

	aliases := []*ast.SymbolAlias{alias}
	commas := make([]*token.Pos, 0)

	for {
		comma, err := p.lexer.Peek()
		if err != nil {
			return nil, err
		}
		if comma.Type != token.Comma {
			break
		}

		if _, err := p.lexer.Scan(); err != nil {
			return nil, err
		}

		alias, err := p.ParseSymbolAlias()
		if err != nil {
			return nil, err
		}

		aliases = append(aliases, alias)
		commas = append(commas, &comma.Position)
	}

	rbrace, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}

	if rbrace.Type != token.RBrace {
		return nil, token.NewPosError(rbrace.Position, "not found RBrace.")
	}

	return &ast.SymbolAliases{
		LBrace:  lbrace.Position,
		Aliases: aliases,
		Commas:  commas,
		RBrace:  rbrace.Position,
	}, nil
}
