package solparser

import (
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func (p *Parser) ParseBlock() (*ast.Block, error) {
	lblace, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}

	if lblace.TokenType != token.LBrace {
		return nil, token.NewPosError(lblace.Pos, "not found LBrace.")
	}

	stmt, err := p.ParseStatement()
	if err != nil {
		return nil, err
	}

	rblace, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}

	if rblace.TokenType != token.RBrace {
		return nil, token.NewPosError(rblace.Pos, "not found RBrace.")
	}

	return &ast.Block{
		LBracePos: lblace.Pos,
		RBracePos: rblace.Pos,
		Nodes:     []ast.Node{stmt},
	}, nil
}
