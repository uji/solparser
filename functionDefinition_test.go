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

func TestParser_ParseVisibility(t *testing.T) {
	tests := []struct {
		input string
		want  ast.Visibility
		err   *token.PosError
	}{
		{
			input: "internal",
			want: ast.Visibility{
				Type:     token.Internal,
				Value:    "internal",
				Position: token.Pos{Column: 1, Line: 1},
			},
		},
		{
			input: "pragma",
			err: &token.PosError{
				Pos: token.Pos{Column: 1, Line: 1},
				Msg: "not found visibility keyword.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(*testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParseVisibility()

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

func TestParser_ParseStateMutability(t *testing.T) {
	tests := []struct {
		input string
		want  ast.StateMutability
		err   *token.PosError
	}{
		{
			input: "pure",
			want: ast.StateMutability{
				Type:     token.Pure,
				Value:    "pure",
				Position: token.Pos{Column: 1, Line: 1},
			},
		},
		{
			input: "pragma",
			err: &token.PosError{
				Pos: token.Pos{Column: 1, Line: 1},
				Msg: "not found state-mutability keyword.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(*testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParseStateMutability()

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

func TestParser_ParseFunctionDefinitionReturns(t *testing.T) {
	tests := []struct {
		input string
		want  *ast.FunctionDefinitionReturns
		err   *token.PosError
	}{
		{
			input: "returns (string)",
			want: &ast.FunctionDefinitionReturns{
				From:   token.Pos{Column: 1, Line: 1},
				LParen: token.Pos{Column: 9, Line: 1},
				ParameterList: []*ast.Parameter{
					{
						TypeName: ast.ElementaryTypeName{
							{
								Type:     token.String,
								Value:    "string",
								Position: token.Pos{Column: 10, Line: 1},
							},
						},
					},
				},
				RParen: token.Pos{Column: 16, Line: 1},
			},
		},
		{
			input: ";",
			want:  nil,
		},
		{
			input: "returns string)",
			err: &token.PosError{
				Pos: token.Pos{Column: 9, Line: 1},
				Msg: "not found arguments LParen.",
			},
		},
		{
			input: "returns (string {",
			err: &token.PosError{
				Pos: token.Pos{Column: 17, Line: 1},
				Msg: "not found arguments RParen.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(*testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParseFunctionDefinitionReturns()

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

func TestParser_ParseFunctionDefinition(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  *ast.FunctionDefinition
		err   *token.PosError
	}{
		{
			name: "normal case",
			input: `function hello() public pure returns (string) {
        return "Hello World!!";
    }`,
			want: &ast.FunctionDefinition{
				From: token.Pos{Column: 1, Line: 1},
				FunctionDescriptor: token.Token{
					Type:     token.Identifier,
					Value:    "hello",
					Position: token.Pos{Column: 10, Line: 1},
				},
				LParen: token.Pos{Column: 15, Line: 1},
				RParen: token.Pos{Column: 16, Line: 1},
				ModifierList: &ast.ModifierList{
					Visibility: &token.Token{
						Type:     token.Public,
						Value:    "public",
						Position: token.Pos{Column: 18, Line: 1},
					},
					StateMutability: &token.Token{
						Type:     token.Pure,
						Value:    "pure",
						Position: token.Pos{Column: 25, Line: 1},
					},
				},
				Returns: &ast.FunctionDefinitionReturns{
					From:   token.Pos{Column: 30, Line: 1},
					LParen: token.Pos{Column: 38, Line: 1},
					ParameterList: []*ast.Parameter{
						{
							TypeName: ast.ElementaryTypeName{
								{
									Type:     token.String,
									Value:    "string",
									Position: token.Pos{Column: 39, Line: 1},
								},
							},
						},
					},
					RParen: token.Pos{Column: 45, Line: 1},
				},
				Block: &ast.Block{
					LBracePos: token.Pos{Column: 47, Line: 1},
					RBracePos: token.Pos{Column: 5, Line: 3},
					Nodes: []ast.Node{
						&ast.ReturnStatement{
							From:    token.Pos{Column: 9, Line: 2},
							SemiPos: token.Pos{Column: 31, Line: 2},
							Expression: &ast.StringLiteral{
								Type:     token.StringLiteral,
								Position: token.Pos{Column: 16, Line: 2},
								Value:    "\"Hello World!!\"",
							},
						},
					},
				},
			},
		},
		{
			name: "function description is invalid",
			input: `function pragma() public pure returns (string) {
        return "Hello World!!";
    }`,
			err: &token.PosError{
				Pos: token.Pos{Column: 10, Line: 1},
				Msg: "not found function description.",
			},
		},
		{
			name: "not found lparen",
			input: `function hello) public pure returns (string) {
        return "Hello World!!";
    }`,
			err: &token.PosError{
				Pos: token.Pos{Column: 15, Line: 1},
				Msg: "not found arguments LParen.",
			},
		},
		{
			name: "not found rparen",
			input: `function hello( public pure returns (string) {
        return "Hello World!!";
    }`,
			err: &token.PosError{
				Pos: token.Pos{Column: 17, Line: 1},
				Msg: "not found arguments RParen.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParseFunctionDefinition()

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
