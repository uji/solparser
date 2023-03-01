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

func TestParser_ParseBlock(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  *ast.Block
		err   *token.PosError
	}{
		{
			name:  "One Statement",
			input: "{ \nreturn \"Hello World!!\";\n }",
			want: &ast.Block{
				LBracePos: token.Pos{Column: 1, Line: 1},
				RBracePos: token.Pos{Column: 2, Line: 3},
				Nodes: []ast.Node{
					&ast.ReturnStatement{
						From:    token.Pos{Column: 1, Line: 2},
						SemiPos: token.Pos{Column: 23, Line: 2},
						Expression: &ast.StringLiteral{
							Type:     token.NonEmptyStringLiteral,
							Position: token.Pos{Column: 8, Line: 2},
							Value:    "\"Hello World!!\"",
						},
					},
				},
			},
			err: nil,
		},
		{
			name:  "Not found LBrace",
			input: "pragma",
			err: &token.PosError{
				Pos: token.Pos{Column: 1, Line: 1},
				Msg: "not found LBrace.",
			},
		},
		{
			name:  "Not found RBrace",
			input: "{ \nreturn \"Hello World!!\";",
			err: &token.PosError{
				Pos: token.Pos{Column: 24, Line: 2},
				Msg: "not found RBrace.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParseBlock()

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
