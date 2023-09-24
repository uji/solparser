package solparser

import (
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func (p *Parser) ParseCallArgumentList() (*ast.CallArgumentList, error) {
	lparen, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}

	tkn, err := p.lexer.Peek()
	if err != nil {
		return nil, err
	}
	if tkn.Type == token.RParen {
		p.lexer.Scan()
		return &ast.CallArgumentList{
			LParen: lparen.Position,
			RParen: tkn.Position,
		}, nil
	}

	var expretions []*ast.CallArgumentListExpretion
	var namedExpretions *ast.CallArgumentListNamedExpretions
	if tkn.Type == token.LBrace {
		es, err := p.ParseCallArgumentListExpretions()
		if err != nil {
			return nil, err
		}
		expretions = es
	} else {
		nes, err := p.ParseCallArgumentListNamedExpretions()
		if err != nil {
			return nil, err
		}
		namedExpretions = nes
	}

	rparen, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if rparen.Type != token.RParen {
		return nil, token.NewPosError(rparen.Position, "not found RParen.")
	}

	return &ast.CallArgumentList{
		LParen:           lparen.Position,
		Expressions:      expretions,
		NamedExpretionsa: namedExpretions,
		RParen:           rparen.Position,
	}, nil
}

func (p *Parser) ParseCallArgumentListExpretions() ([]*ast.CallArgumentListExpretion, error) {
	ex, err := p.ParseExpression()
	if err != nil {
		return nil, err
	}

	var exs []*ast.CallArgumentListExpretion
	for {
		cmm, err := p.lexer.Peek()
		if err != nil || cmm.Type != token.Comma {
			break
		}
		p.lexer.Scan()

		exs = append(exs, &ast.CallArgumentListExpretion{
			Expression: ex,
			Comma:      &cmm.Position,
		})

		e, err := p.ParseExpression()
		ex = e
	}

	exs = append(exs, &ast.CallArgumentListExpretion{
		Expression: ex,
	})

	return exs, nil
}

func (p *Parser) ParseCallArgumentListNamedExpretions() (*ast.CallArgumentListNamedExpretions, error) {
	lbrace, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}

	exs := make([]*ast.CallArgumentListNamedExpretion, 0, 1)

	for {
		id, err := p.lexer.Peek()
		if err != nil {
			return nil, err
		}

		if !isIdentifier(id) {
			return nil, token.NewPosError(id.Position, "not found identifier.")
		}
		p.lexer.Scan()

		cln, err := p.lexer.Scan()
		if err != nil {
			return nil, err
		}
		if cln.Type != token.Colon {
			return nil, token.NewPosError(cln.Position, "not found Colon.")
		}

		ex, err := p.ParseExpression()
		if err != nil {
			return nil, err
		}

		cmm, err := p.lexer.Peek()
		if err != nil {
			return nil, err
		}
		if cmm.Type != token.Comma {
			exs = append(exs, &ast.CallArgumentListNamedExpretion{
				Identifier: ast.Identifier(id),
				Colon:      cln.Position,
				Expression: ex,
			})
			break
		}
		p.lexer.Scan()
		exs = append(exs, &ast.CallArgumentListNamedExpretion{
			Identifier: ast.Identifier(id),
			Colon:      cln.Position,
			Expression: ex,
			Comma:      &cmm.Position,
		})
	}

	rbrace, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}
	if rbrace.Type != token.RBrace {
		return nil, token.NewPosError(rbrace.Position, "not found RBrace.")
	}

	return &ast.CallArgumentListNamedExpretions{
		RBrace:          rbrace.Position,
		NamedExpretions: exs,
		LBrace:          lbrace.Position,
	}, nil
}
