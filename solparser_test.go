package solparser_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/uji/solparser"
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  *ast.SourceUnit
		err   *token.PosError
	}{
		{
			name: "normal",
			input: `pragma solidity ^0.8.13;

contract HelloWorld {
    function hello() public pure returns (string) {
        return "Hello World!!";
    }
}`,
			want: &ast.SourceUnit{
				PragmaDirective: &ast.PragmaDirective{
					Pragma: token.Pos{Column: 1, Line: 1},
					PragmaTokens: []*token.Token{
						{
							TokenType: token.Identifier,
							Text:      "solidity",
							Pos:       token.Pos{Column: 8, Line: 1},
						},
						{
							TokenType: token.BitXor,
							Text:      "^",
							Pos:       token.Pos{Column: 17, Line: 1},
						},
						{
							TokenType: token.Identifier,
							Text:      "0.8.13",
							Pos:       token.Pos{Column: 18, Line: 1},
						},
					},
					Semicolon: token.Pos{Column: 24, Line: 1},
				},
				ContractDefinition: &ast.ContractDefinition{
					Contract: token.Pos{Column: 1, Line: 3},
					Identifier: token.Token{
						TokenType: token.Identifier,
						Text:      "HelloWorld",
						Pos:       token.Pos{Column: 10, Line: 3},
					},
					LBrace: token.Pos{Column: 21, Line: 3},
					ContractBodyElements: []ast.ContractBodyElement{
						&ast.FunctionDefinition{
							From: token.Pos{Column: 5, Line: 4},
							FunctionDescriptor: token.Token{
								TokenType: token.Identifier,
								Text:      "hello",
								Pos:       token.Pos{Column: 14, Line: 4},
							},
							LParen: token.Pos{Column: 19, Line: 4},
							RParen: token.Pos{Column: 20, Line: 4},
							ModifierList: &ast.ModifierList{
								Visibility: &token.Token{
									TokenType: token.Public,
									Text:      "public",
									Pos:       token.Pos{Column: 22, Line: 4},
								},
								StateMutability: &token.Token{
									TokenType: token.Pure,
									Text:      "pure",
									Pos:       token.Pos{Column: 29, Line: 4},
								},
							},
							Returns: &ast.FunctionDefinitionReturns{
								From:   token.Pos{Column: 34, Line: 4},
								LParen: token.Pos{Column: 42, Line: 4},
								ParameterList: []*ast.Parameter{
									{
										TypeName: ast.ElementaryTypeName{
											{
												TokenType: token.String,
												Text:      "string",
												Pos:       token.Pos{Column: 43, Line: 4},
											},
										},
									},
								},
								RParen: token.Pos{Column: 49, Line: 4},
							},
							Block: &ast.Block{
								LBracePos: token.Pos{Column: 51, Line: 4},
								RBracePos: token.Pos{Column: 5, Line: 6},
								Nodes: []ast.Node{
									&ast.ReturnStatement{
										From:    token.Pos{Column: 9, Line: 5},
										SemiPos: token.Pos{Column: 31, Line: 5},
										Expression: &ast.StringLiteral{
											From:  token.Pos{Column: 16, Line: 5},
											To:    token.Pos{Column: 30, Line: 5},
											Value: "\"Hello World!!\"",
										},
									},
								},
							},
						},
					},
					RBrace: token.Pos{Column: 1, Line: 7},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.Parse()

			var sErr *token.PosError
			if errors.As(err, &sErr) {
				if diff := cmp.Diff(tt.err, sErr); diff != "" {
					t.Errorf("%s", diff)
				}
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("%s", diff)
			}
		})
	}
}

func TestParser_ParseBooleanLiteral(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  *ast.BooleanLiteral
		err   *token.PosError
	}{
		{
			name:  "true case",
			input: "true",
			want: &ast.BooleanLiteral{
				Token: token.Token{
					TokenType: token.TrueLiteral,
					Text:      "true",
					Pos: token.Pos{
						Column: 1,
						Line:   1,
					},
				},
			},
			err: nil,
		},
		{
			name:  "false case",
			input: "false",
			want: &ast.BooleanLiteral{
				Token: token.Token{
					TokenType: token.FalseLiteral,
					Text:      "false",
					Pos: token.Pos{
						Column: 1,
						Line:   1,
					},
				},
			},
			err: nil,
		},
		{
			name:  "not true or false",
			input: "solidity",
			want:  nil,
			err: &token.PosError{
				Pos: token.Pos{
					Column: 1,
					Line:   1,
				},
				Msg: "not found keyword true or false",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParseBooleanLiteral()

			var sErr *token.PosError
			if errors.As(err, &sErr) {
				if diff := cmp.Diff(tt.err, sErr); diff != "" {
					t.Errorf("%s", diff)
				}
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("%s", diff)
			}
		})
	}
}
