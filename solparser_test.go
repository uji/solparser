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

func TestParserParseBooleanLiteral(t *testing.T) {
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
					TokenType: token.True,
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
					TokenType: token.False,
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
