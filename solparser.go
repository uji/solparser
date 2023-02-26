package solparser

import (
	"io"

	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/lexer"
	"github.com/uji/solparser/token"
)

type Parser struct {
	input io.Reader
	lexer *lexer.Lexer
}

func New(input io.Reader) *Parser {
	return &Parser{
		input: input,
		lexer: lexer.New(input),
	}
}

func (p *Parser) Parse() (*ast.SourceUnit, error) {
	var pragmaDirective *ast.PragmaDirective
	var contractDefinition *ast.ContractDefinition
	var functionDefinition *ast.FunctionDefinition
	for {
		tkn, err := p.lexer.Peek()
		if err != nil {
			return nil, err
		}

		switch tkn.Type {
		case token.Pragma:
			prgm, err := p.ParsePragmaDirective()
			if err != nil {
				return nil, err
			}
			pragmaDirective = prgm
		case token.Abstract, token.Contract:
			cntrct, err := p.ParseContractDefinition()
			if err != nil {
				return nil, err
			}
			contractDefinition = cntrct
		case token.Function:
			fnc, err := p.ParseFunctionDefinition()
			if err != nil {
				return nil, err
			}
			functionDefinition = fnc
		case token.EOS:
			return &ast.SourceUnit{
				PragmaDirective:    pragmaDirective,
				ContractDefinition: contractDefinition,
				FunctionDefinition: functionDefinition,
			}, nil
		default:
			return nil, token.NewPosError(tkn.Position, "invalid")
		}
	}
}

func (p *Parser) ParseBooleanLiteral() (*ast.BooleanLiteral, error) {
	tkn, err := p.lexer.Scan()
	if err != nil {
		return nil, err
	}

	if tkn.Type != token.TrueLiteral && tkn.Type != token.FalseLiteral {
		return nil, token.NewPosError(tkn.Position, "not found keyword true or false")
	}

	return &ast.BooleanLiteral{
		Token: tkn,
	}, nil
}
