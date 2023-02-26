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

func TestParser_ParseLiteral(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  ast.Literal
		err   *token.PosError
	}{
		{
			name:  "normal",
			input: `"Hello World!!";`,
			want: &ast.StringLiteral{
				Type:     token.StringLiteral,
				Value:    `"Hello World!!"`,
				Position: pos(1, 1),
			},
		},
		{
			name:  `Including \n`,
			input: "\"Hello \nWorld!!\";",
			want: &ast.StringLiteral{
				Type:     token.StringLiteral,
				Value:    "\"Hello \nWorld!!\"",
				Position: pos(1, 1),
			},
		},
		// {
		// 	name:  `Next token is \n`,
		// 	input: "\"Hello World!!\"\n",
		// 	want: &ast.StringLiteral{
		// 		Value: `"Hello World!!"`,
		// 		From: token.Pos{
		// 			Column: 1,
		// 			Line:   1,
		// 		},
		// 		To: token.Pos{
		// 			Column: 15,
		// 			Line:   1,
		// 		},
		// 	},
		// 	err: nil,
		// },
		{
			name:  "Not Literal",
			input: "pragma",
			err: &token.PosError{
				Pos: token.Pos{
					Column: 1,
					Line:   1,
				},
				Msg: "not found string literal quote",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParseLiteral()

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
