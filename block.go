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

	if lblace.Type != token.LBrace {
		return nil, token.NewPosError(lblace.Position, "not found LBrace.")
	}

	stmt, err := p.ParseStatement()
	if err != nil {
		return nil, err
	}

	rblace, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}

	if rblace.Type != token.RBrace {
		return nil, token.NewPosError(rblace.Position, "not found RBrace.")
	}

	return &ast.Block{
		LBracePos: lblace.Position,
		RBracePos: rblace.Position,
		Nodes:     []ast.Node{stmt},
	}, nil
}
