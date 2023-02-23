package solparser

import (
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func (p *Parser) ParseVisibility() (ast.Visibility, error) {
	tkn, err := p.lexer.Scan()
	if err != nil {
		return token.Token{}, err
	}

	switch tkn.TokenType {
	case token.Internal, token.External, token.Public, token.Private:
		return tkn, nil
	}

	return token.Token{}, token.NewPosError(tkn.Pos, "not found visibility keyword.")
}

func (p *Parser) ParseStateMutability() (ast.StateMutability, error) {
	tkn, err := p.lexer.Scan()
	if err != nil {
		return token.Token{}, err
	}

	switch tkn.TokenType {
	case token.Pure, token.View, token.Payable:
		return tkn, nil
	}

	return token.Token{}, token.NewPosError(tkn.Pos, "not found state-mutability keyword.")
}

func (p *Parser) ParseFunctionDefinitionReturns() (*ast.FunctionDefinitionReturns, error) {
	tkn, err := p.lexer.Peek()
	if err != nil {
		return nil, err
	}
	if tkn.TokenType != token.Returns {
		return nil, nil
	}
	if _, err := p.lexer.Scan(); err != nil {
		return nil, err
	}

	lparen, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if lparen.TokenType != token.LParen {
		return nil, token.NewPosError(lparen.Pos, "not found arguments LParen.")
	}

	pl, err := p.ParseParameterList()
	if err != nil {
		return nil, err
	}

	rparen, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if rparen.TokenType != token.RParen {
		return nil, token.NewPosError(rparen.Pos, "not found arguments RParen.")
	}

	return &ast.FunctionDefinitionReturns{
		From:          tkn.Pos,
		LParen:        lparen.Pos,
		ParameterList: pl,
		RParen:        rparen.Pos,
	}, nil
}

func (p *Parser) ParseFunctionDefinition() (*ast.FunctionDefinition, error) {
	from, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}

	if from.TokenType != token.Function {
		return nil, token.NewPosError(from.Pos, "not found function keyword.")
	}

	dsc, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}

	switch dsc.TokenType {
	case token.Identifier, token.From, token.Error, token.Revert, token.Global, token.Fallback, token.Receive:
	default:
		return nil, token.NewPosError(dsc.Pos, "not found function description.")
	}

	lparen, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if lparen.TokenType != token.LParen {
		return nil, token.NewPosError(lparen.Pos, "not found arguments LParen.")
	}

	rparen, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if rparen.TokenType != token.RParen {
		return nil, token.NewPosError(rparen.Pos, "not found arguments RParen.")
	}

	modifierList := &ast.ModifierList{
		Visibility:      nil,
		StateMutability: nil,
	}

	for {
		tkn, err := p.lexer.Peek()
		if err != nil {
			return nil, err
		}

		switch tkn.TokenType {
		case token.Internal, token.External, token.Public, token.Private:
			vs, err := p.ParseVisibility()
			if err != nil {
				return nil, err
			}
			modifierList.Visibility = &vs
			continue
		case token.Pure, token.View, token.Payable:
			sm, err := p.ParseStateMutability()
			if err != nil {
				return nil, err
			}
			modifierList.StateMutability = &sm
			continue
		}
		break
	}

	r, err := p.ParseFunctionDefinitionReturns()
	if err != nil {
		return nil, err
	}

	b, err := p.ParseBlock()
	if err != nil {
		return nil, err
	}

	return &ast.FunctionDefinition{
		From:               from.Pos,
		FunctionDescriptor: dsc,
		LParen:             lparen.Pos,
		RParen:             rparen.Pos,
		ModifierList:       modifierList,
		Returns:            r,
		Block:              b,
	}, nil
}
