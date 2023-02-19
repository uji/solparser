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

	if tkn.TokenType == token.Return {
		return p.ParseReturnStatement()
	}

	return nil, token.NewPosError(tkn.Pos, "not found statement.")
}

func (p *Parser) ParseReturnStatement() (ast.Statement, error) {
	rtn, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}

	if rtn.TokenType != token.Return {
		return nil, token.NewPosError(rtn.Pos, "not found return keyword.")
	}

	exp, err := p.ParseExpression()
	if err != nil {
		return nil, err
	}

	semi, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}

	if semi.TokenType != token.Semicolon {
		return nil, token.NewPosError(semi.Pos, "not found semicolon.")
	}

	return &ast.ReturnStatement{
		From:       rtn.Pos,
		SemiPos:    semi.Pos,
		Expression: exp,
	}, nil
}
