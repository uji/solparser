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
			input: "pragma solidity ^0.8.13;",
			want: []lexer.Token{
				{lexer.Pragma, "pragma", lexer.Position{Start: 0, Size: 6, Line: 0}},
				{lexer.Unknown, "solidity", lexer.Position{Start: 7, Size: 8, Line: 0}},
				{lexer.Hat, "^", lexer.Position{Start: 16, Size: 1, Line: 0}},
				{lexer.Unknown, "0.8.13", lexer.Position{Start: 17, Size: 6, Line: 0}},
				{lexer.Semicolon, ";", lexer.Position{Start: 23, Size: 1, Line: 0}},
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
