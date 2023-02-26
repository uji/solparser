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

func TestParser_ParseStatement(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  ast.Statement
		err   *token.PosError
	}{
		{
			name:  "ReturnStatement",
			input: `return "Hello World!!";`,
			want: &ast.ReturnStatement{
				From:    token.Pos{Column: 1, Line: 1},
				SemiPos: token.Pos{Column: 23, Line: 1},
				Expression: &ast.StringLiteral{
					Type:     token.StringLiteral,
					Position: token.Pos{Column: 8, Line: 1},
					Value:    `"Hello World!!"`,
				},
			},
			err: nil,
		},
		{
			name:  "Not Statement",
			input: "pragma",
			err: &token.PosError{
				Pos: token.Pos{Column: 1, Line: 1},
				Msg: "not found statement.",
			},
		},
		{
			name:  "Found broken return statement",
			input: "return pragma",
			err: &token.PosError{
				Pos: token.Pos{Column: 8, Line: 1},
				Msg: "not found expression.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParseStatement()

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

func TestParser_ParseReturnStatement(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  ast.Statement
		err   *token.PosError
	}{
		{
			name:  "normal",
			input: `return "Hello World!!";`,
			want: &ast.ReturnStatement{
				From:    token.Pos{Column: 1, Line: 1},
				SemiPos: token.Pos{Column: 23, Line: 1},
				Expression: &ast.StringLiteral{
					Type:     token.StringLiteral,
					Position: token.Pos{Column: 8, Line: 1},
					Value:    `"Hello World!!"`,
				},
			},
			err: nil,
		},
		{
			name:  "Not found return keyword",
			input: "pragma",
			err: &token.PosError{
				Pos: token.Pos{Column: 1, Line: 1},
				Msg: "not found return keyword.",
			},
		},
		{
			name:  "Not found expression",
			input: "return pragma",
			err: &token.PosError{
				Pos: token.Pos{Column: 8, Line: 1},
				Msg: "not found expression.",
			},
		},
		{
			name:  "Not found semicolon.",
			input: `return "test" pragma`,
			err: &token.PosError{
				Pos: token.Pos{Column: 15, Line: 1},
				Msg: "not found semicolon.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParseReturnStatement()

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
