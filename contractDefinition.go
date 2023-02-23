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
	if cntr.TokenType == token.Abstract {
		pos := cntr.Pos
		abstractPos = &pos
		cntr, err = p.lexer.Scan()
		if err != nil {
			return nil, err
		}
	}
	if cntr.TokenType != token.Contract {
		return nil, token.NewPosError(cntr.Pos, "not found contract keyword.")
	}

	i, err := p.ParseIdentifier()
	if err != nil {
		return nil, err
	}

	lbrace, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if lbrace.TokenType != token.LBrace {
		return nil, token.NewPosError(lbrace.Pos, "not found left brace.")
	}

	fn, err := p.ParseFunctionDefinition()
	if err != nil {
		return nil, err
	}

	rbrace, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if rbrace.TokenType != token.RBrace {
		return nil, token.NewPosError(rbrace.Pos, "not found right brace.")
	}

	return &ast.ContractDefinition{
		Abstract:             abstractPos,
		Contract:             cntr.Pos,
		Identifier:           i,
		LBrace:               lbrace.Pos,
		ContractBodyElements: []ast.ContractBodyElement{fn},
		RBrace:               rbrace.Pos,
	}, nil
}
