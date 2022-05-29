package lexer_test

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/uji/solparser/lexer"
)

func TestLexer_Scan(t *testing.T) {
	cases := []struct {
		input string
		want  []lexer.Token
	}{
		{
			input: "pragma solidity ^0.8.13;\n\ncontract HelloWorld { ",
			want: []lexer.Token{
				{lexer.Pragma, "pragma", lexer.Position{Column: 1, Line: 1}},
				{lexer.Unknown, "solidity", lexer.Position{Column: 8, Line: 1}},
				{lexer.Hat, "^", lexer.Position{Column: 17, Line: 1}},
				{lexer.Unknown, "0.8.13", lexer.Position{Column: 18, Line: 1}},
				{lexer.Semicolon, ";", lexer.Position{Column: 24, Line: 1}},
				{lexer.Contract, "contract", lexer.Position{Column: 1, Line: 3}},
				{lexer.Unknown, "HelloWorld", lexer.Position{Column: 10, Line: 3}},
				{lexer.BraceL, "{", lexer.Position{Column: 21, Line: 3}},
			},
		},
	}

	for n, c := range cases {
		buf := strings.NewReader(c.input)
		l := lexer.New(buf)
		got := make([]lexer.Token, 0, len(c.want))
		for i := 0; i < len(c.want); i++ {
			l.Scan()
			got = append(got, l.Token())
		}
		if l.Scan() {
			t.Errorf("#%d: scan ran too long, got %q", n, got)
		}
		if diff := cmp.Diff(c.want, got); diff != "" {
			t.Errorf("#%d: %s", n, diff)
		}
		if err := l.Error(); err != nil {
			t.Errorf("#%d: %v", n, err)
		}
	}
}
