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

func TestParser_ParseContractDefinition(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  *ast.ContractDefinition
		err   *token.PosError
	}{
		{
			name: "normal case",
			input: `contract HelloWorld {
    function hello() public pure returns (string) {
        return "Hello World!!";
    }
}`,
			want: &ast.ContractDefinition{
				Contract: token.Pos{Column: 1, Line: 1},
				Identifier: ast.Identifier{
					Type:     token.Identifier,
					Value:    "HelloWorld",
					Position: token.Pos{Column: 10, Line: 1},
				},
				LBrace: token.Pos{Column: 21, Line: 1},
				ContractBodyElements: []ast.ContractBodyElement{
					&ast.FunctionDefinition{
						From: token.Pos{Column: 5, Line: 2},
						FunctionDescriptor: token.Token{
							Type:     token.Identifier,
							Value:    "hello",
							Position: token.Pos{Column: 14, Line: 2},
						},
						LParen: token.Pos{Column: 19, Line: 2},
						RParen: token.Pos{Column: 20, Line: 2},
						ModifierList: &ast.ModifierList{
							Visibility: &token.Token{
								Type:     token.Public,
								Value:    "public",
								Position: token.Pos{Column: 22, Line: 2},
							},
							StateMutability: &token.Token{
								Type:     token.Pure,
								Value:    "pure",
								Position: token.Pos{Column: 29, Line: 2},
							},
						},
						Returns: &ast.FunctionDefinitionReturns{
							From:   token.Pos{Column: 34, Line: 2},
							LParen: token.Pos{Column: 42, Line: 2},
							ParameterList: []*ast.Parameter{
								{
									TypeName: ast.ElementaryTypeName{
										{
											Type:     token.String,
											Value:    "string",
											Position: token.Pos{Column: 43, Line: 2},
										},
									},
								},
							},
							RParen: token.Pos{Column: 49, Line: 2},
						},
						Block: &ast.Block{
							LBracePos: token.Pos{Column: 51, Line: 2},
							RBracePos: token.Pos{Column: 5, Line: 4},
							Nodes: []ast.Node{
								&ast.ReturnStatement{
									From:    token.Pos{Column: 9, Line: 3},
									SemiPos: token.Pos{Column: 31, Line: 3},
									Expression: &ast.StringLiteral{
										Type:     token.NonEmptyStringLiteral,
										Position: token.Pos{Column: 16, Line: 3},
										Value:    "\"Hello World!!\"",
									},
								},
							},
						},
					},
				},
				RBrace: token.Pos{Column: 1, Line: 5},
			},
		},
		{
			name: "with abstract",
			input: `abstract contract HelloWorld {
    function hello() public pure returns (string) {
        return "Hello World!!";
    }
}`,
			want: &ast.ContractDefinition{
				Abstract: &token.Pos{Column: 1, Line: 1},
				Contract: token.Pos{Column: 10, Line: 1},
				Identifier: ast.Identifier{
					Type:     token.Identifier,
					Value:    "HelloWorld",
					Position: token.Pos{Column: 19, Line: 1},
				},
				LBrace: token.Pos{Column: 30, Line: 1},
				ContractBodyElements: []ast.ContractBodyElement{
					&ast.FunctionDefinition{
						From: token.Pos{Column: 5, Line: 2},
						FunctionDescriptor: token.Token{
							Type:     token.Identifier,
							Value:    "hello",
							Position: token.Pos{Column: 14, Line: 2},
						},
						LParen: token.Pos{Column: 19, Line: 2},
						RParen: token.Pos{Column: 20, Line: 2},
						ModifierList: &ast.ModifierList{
							Visibility: &token.Token{
								Type:     token.Public,
								Value:    "public",
								Position: token.Pos{Column: 22, Line: 2},
							},
							StateMutability: &token.Token{
								Type:     token.Pure,
								Value:    "pure",
								Position: token.Pos{Column: 29, Line: 2},
							},
						},
						Returns: &ast.FunctionDefinitionReturns{
							From:   token.Pos{Column: 34, Line: 2},
							LParen: token.Pos{Column: 42, Line: 2},
							ParameterList: []*ast.Parameter{
								{
									TypeName: ast.ElementaryTypeName{
										{
											Type:     token.String,
											Value:    "string",
											Position: token.Pos{Column: 43, Line: 2},
										},
									},
								},
							},
							RParen: token.Pos{Column: 49, Line: 2},
						},
						Block: &ast.Block{
							LBracePos: token.Pos{Column: 51, Line: 2},
							RBracePos: token.Pos{Column: 5, Line: 4},
							Nodes: []ast.Node{
								&ast.ReturnStatement{
									From:    token.Pos{Column: 9, Line: 3},
									SemiPos: token.Pos{Column: 31, Line: 3},
									Expression: &ast.StringLiteral{
										Type:     token.NonEmptyStringLiteral,
										Position: token.Pos{Column: 16, Line: 3},
										Value:    "\"Hello World!!\"",
									},
								},
							},
						},
					},
				},
				RBrace: token.Pos{Column: 1, Line: 5},
			},
		},
		{
			name: "not found contract keyword",
			input: `HelloWorld {
    function hello() public pure returns (string) {
        return "Hello World!!";
    }
}`,
			err: &token.PosError{
				Pos: token.Pos{Column: 1, Line: 1},
				Msg: "not found contract keyword.",
			},
		},
		{
			name: "not found identifier",
			input: `contract {
    function hello() public pure returns (string) {
        return "Hello World!!";
    }
}`,
			err: &token.PosError{
				Pos: token.Pos{Column: 10, Line: 1},
				Msg: "keyword is not available as identifier.",
			},
		},
		{
			name:  "not found LBrace",
			input: "contract HelloWorld function",
			err: &token.PosError{
				Pos: token.Pos{Column: 21, Line: 1},
				Msg: "not found left brace.",
			},
		},
		{
			name:  "not found FunctionDefinition",
			input: "contract HelloWorld { hello() }",
			err: &token.PosError{
				Pos: token.Pos{Column: 23, Line: 1},
				Msg: "not found function keyword.",
			},
		},
		{
			name: "not found RBrace",
			input: `contract HelloWorld {
    function hello() public pure returns (string) {
        return "Hello World!!";
    }`,
			err: &token.PosError{
				Pos: token.Pos{Column: 6, Line: 4},
				Msg: "not found right brace.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParseContractDefinition()

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
