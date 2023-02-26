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

	switch tkn.Type {
	case token.Internal, token.External, token.Public, token.Private:
		return tkn, nil
	}

	return token.Token{}, token.NewPosError(tkn.Position, "not found visibility keyword.")
}

func (p *Parser) ParseStateMutability() (ast.StateMutability, error) {
	tkn, err := p.lexer.Scan()
	if err != nil {
		return token.Token{}, err
	}

	switch tkn.Type {
	case token.Pure, token.View, token.Payable:
		return tkn, nil
	}

	return token.Token{}, token.NewPosError(tkn.Position, "not found state-mutability keyword.")
}

func (p *Parser) ParseFunctionDefinitionReturns() (*ast.FunctionDefinitionReturns, error) {
	tkn, err := p.lexer.Peek()
	if err != nil {
		return nil, err
	}
	if tkn.Type != token.Returns {
		return nil, nil
	}
	if _, err := p.lexer.Scan(); err != nil {
		return nil, err
	}

	lparen, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if lparen.Type != token.LParen {
		return nil, token.NewPosError(lparen.Position, "not found arguments LParen.")
	}

	pl, err := p.ParseParameterList()
	if err != nil {
		return nil, err
	}

	rparen, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if rparen.Type != token.RParen {
		return nil, token.NewPosError(rparen.Position, "not found arguments RParen.")
	}

	return &ast.FunctionDefinitionReturns{
		From:          tkn.Position,
		LParen:        lparen.Position,
		ParameterList: pl,
		RParen:        rparen.Position,
	}, nil
}

func (p *Parser) ParseFunctionDefinition() (*ast.FunctionDefinition, error) {
	from, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}

	if from.Type != token.Function {
		return nil, token.NewPosError(from.Position, "not found function keyword.")
	}

	dsc, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}

	switch dsc.Type {
	case token.Identifier, token.From, token.Error, token.Revert, token.Global, token.Fallback, token.Receive:
	default:
		return nil, token.NewPosError(dsc.Position, "not found function description.")
	}

	lparen, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if lparen.Type != token.LParen {
		return nil, token.NewPosError(lparen.Position, "not found arguments LParen.")
	}

	rparen, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if rparen.Type != token.RParen {
		return nil, token.NewPosError(rparen.Position, "not found arguments RParen.")
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

		switch tkn.Type {
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
		From:               from.Position,
		FunctionDescriptor: dsc,
		LParen:             lparen.Position,
		RParen:             rparen.Position,
		ModifierList:       modifierList,
		Returns:            r,
		Block:              b,
	}, nil
}
