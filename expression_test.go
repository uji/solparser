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

func TestParser_ParseExpression(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  ast.Expression
		err   *token.PosError
	}{
		{
			name:  "Literal",
			input: `"Hello World!!";`,
			want: &ast.StringLiteral{
				Str: `"Hello World!!"`,
				Position: token.Pos{
					Column: 1,
					Line:   1,
				},
				EndPos: token.Pos{
					Column: 15,
					Line:   1,
				},
			},
			err: nil,
		},
		{
			name:  "Not expression",
			input: "pragma",
			err: &token.PosError{
				Pos: token.Pos{
					Column: 1,
					Line:   1,
				},
				Msg: "not found expression.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParseExpression()

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
