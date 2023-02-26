package solparser

import (
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func (p *Parser) ParseStatement() (ast.Statement, error) {
	tkn, err := p.lexer.Peek()
	if err != nil {
		return nil, err
	}

	if tkn.Type == token.Return {
		return p.ParseReturnStatement()
	}

	return nil, token.NewPosError(tkn.Position, "not found statement.")
}

func (p *Parser) ParseReturnStatement() (ast.Statement, error) {
	rtn, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}

	if rtn.Type != token.Return {
		return nil, token.NewPosError(rtn.Position, "not found return keyword.")
	}

	exp, err := p.ParseExpression()
	if err != nil {
		return nil, err
	}

	semi, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}

	if semi.Type != token.Semicolon {
		return nil, token.NewPosError(semi.Position, "not found semicolon.")
	}

	return &ast.ReturnStatement{
		From:       rtn.Position,
		SemiPos:    semi.Position,
		Expression: exp,
	}, nil
}
