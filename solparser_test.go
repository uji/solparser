package solparser_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/uji/solparser"
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/lexer"
)

func TestParserParsePragmaDirective(t *testing.T) {
	cases := []struct {
		input string
		want  *ast.PragmaDirective
		err   *solparser.Error
	}{
		{
			input: "pragma solidity ^0.8.13;",
			want: &ast.PragmaDirective{
				PragmaName: "solidity",
				PragmaValue: ast.PragmaValue{
					Version:    "0.8.13",
					Expression: "^",
				},
			},
			err: nil,
		},
		{
			input: "solidity ^0.8.13;",
			want:  nil,
			err: &solparser.Error{
				Token: lexer.Token{
					TokenType: lexer.Unknown,
					Text:      "solidity",
					Pos: lexer.Position{
						Column: 1,
						Line:   1,
					},
				},
				Msg: "not found pragma",
			},
		},
	}

	for n, c := range cases {
		r := strings.NewReader(c.input)
		p := solparser.New(r)

		got, err := p.ParsePragmaDirective()

		var sErr *solparser.Error
		if errors.As(err, &sErr) {
			if diff := cmp.Diff(c.err, sErr); diff != "" {
				t.Errorf("#%d %s", n, diff)
			}
		}

		if diff := cmp.Diff(c.want, got); diff != "" {
			t.Errorf("#%d %s", n, diff)
		}
	}
}
