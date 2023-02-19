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

func TestParser_ParseParameter(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  *ast.Parameter
		err   *token.PosError
	}{
		{
			name:  "TypeName only",
			input: "string",
			want: &ast.Parameter{
				TypeName: ast.ElementaryTypeName{
					{
						TokenType: token.String,
						Text:      "string",
						Pos:       token.Pos{Column: 1, Line: 1},
					},
				},
			},
			err: nil,
		},
		{
			name:  "Not ParameterList",
			input: "pragma",
			err: &token.PosError{
				Pos: token.Pos{
					Column: 1,
					Line:   1,
				},
				Msg: "not found type-name.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParseParameter()

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

func TestParser_ParseParameterList(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  ast.ParameterList
		err   *token.PosError
	}{
		{
			name:  "One Parameter",
			input: "string",
			want: ast.ParameterList{
				{
					TypeName: ast.ElementaryTypeName{
						{
							TokenType: token.String,
							Text:      "string",
							Pos:       token.Pos{Column: 1, Line: 1},
						},
					},
				},
			},
			err: nil,
		},
		{
			name:  "Some Parameter",
			input: "string, bool",
			want: ast.ParameterList{
				{
					TypeName: ast.ElementaryTypeName{
						{
							TokenType: token.String,
							Text:      "string",
							Pos:       token.Pos{Column: 1, Line: 1},
						},
					},
				},
				{
					TypeName: ast.ElementaryTypeName{
						{
							TokenType: token.Bool,
							Text:      "bool",
							Pos:       token.Pos{Column: 9, Line: 1},
						},
					},
				},
			},
			err: nil,
		},
		{
			name:  "Not ParameterList",
			input: "pragma",
			err: &token.PosError{
				Pos: token.Pos{
					Column: 1,
					Line:   1,
				},
				Msg: "not found type-name.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParseParameterList()

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
