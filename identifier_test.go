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

func TestParser_ParseIdentifier(t *testing.T) {
	tests := []struct {
		input string
		want  ast.Identifier
		err   *token.PosError
	}{
		{
			input: "identifier",
			want: token.Token{
				TokenType: token.Identifier,
				Text:      "identifier",
				Pos:       token.Pos{Column: 1, Line: 1},
			},
		},
		{
			input: "from",
			want: token.Token{
				TokenType: token.From,
				Text:      "from",
				Pos:       token.Pos{Column: 1, Line: 1},
			},
		},
		{
			input: "error",
			want: token.Token{
				TokenType: token.Error,
				Text:      "error",
				Pos:       token.Pos{Column: 1, Line: 1},
			},
		},
		{
			input: "global",
			want: token.Token{
				TokenType: token.Global,
				Text:      "global",
				Pos:       token.Pos{Column: 1, Line: 1},
			},
		},
		{
			input: "pragma",
			err: &token.PosError{
				Pos: token.Pos{Column: 1, Line: 1},
				Msg: "keyword is not available as identifier.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(*testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParseIdentifier()

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
