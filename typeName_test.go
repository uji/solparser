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

func TestParser_ParseTypeName(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  ast.TypeName
		err   *token.PosError
	}{
		{
			name:  "ElementaryTypeName",
			input: "string)",
			want: ast.ElementaryTypeName{
				{
					Type:     token.String,
					Value:    "string",
					Position: token.Pos{Column: 1, Line: 1},
				},
			},
		},
		{
			name:  "Not TypeName",
			input: "pragma",
			err: &token.PosError{
				Pos: token.Pos{Column: 1, Line: 1},
				Msg: "not found type-name.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParseTypeName()

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

func TestParser_ParseElementaryTypeName(t *testing.T) {
	tests := []struct {
		input string
		want  ast.TypeName
		err   *token.PosError
	}{
		{
			input: "address",
			want: ast.ElementaryTypeName{
				{
					Type:     token.Address,
					Value:    "address",
					Position: token.Pos{Column: 1, Line: 1},
				},
			},
		},
		{
			input: "bool",
			want: ast.ElementaryTypeName{
				{
					Type:     token.Bool,
					Value:    "bool",
					Position: token.Pos{Column: 1, Line: 1},
				},
			},
		},
		{
			input: "string",
			want: ast.ElementaryTypeName{
				{
					Type:     token.String,
					Value:    "string",
					Position: token.Pos{Column: 1, Line: 1},
				},
			},
		},
		{
			input: "address payable",
			want: ast.ElementaryTypeName{
				{
					Type:     token.Address,
					Value:    "address",
					Position: token.Pos{Column: 1, Line: 1},
				},
				{
					Type:     token.Payable,
					Value:    "payable",
					Position: token.Pos{Column: 9, Line: 1},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(*testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParseElementaryTypeName()

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
