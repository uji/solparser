package solparser

import (
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func (p *Parser) ParseContractDefinition() (*ast.ContractDefinition, error) {
	cntr, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	var abstractPos *token.Pos
	if cntr.Type == token.Abstract {
		pos := cntr.Position
		abstractPos = &pos
		cntr, err = p.lexer.Scan()
		if err != nil {
			return nil, err
		}
	}
	if cntr.Type != token.Contract {
		return nil, token.NewPosError(cntr.Position, "not found contract keyword.")
	}

	i, err := p.ParseIdentifier()
	if err != nil {
		return nil, err
	}

	lbrace, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if lbrace.Type != token.LBrace {
		return nil, token.NewPosError(lbrace.Position, "not found left brace.")
	}

	fn, err := p.ParseFunctionDefinition()
	if err != nil {
		return nil, err
	}

	rbrace, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if rbrace.Type != token.RBrace {
		return nil, token.NewPosError(rbrace.Position, "not found right brace.")
	}

	return &ast.ContractDefinition{
		Abstract:             abstractPos,
		Contract:             cntr.Position,
		Identifier:           i,
		LBrace:               lbrace.Position,
		ContractBodyElements: []ast.ContractBodyElement{fn},
		RBrace:               rbrace.Position,
	}, nil
}
