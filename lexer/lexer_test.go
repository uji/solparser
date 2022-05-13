package lexer_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/uji/solparser/lexer"
)

// Test that the token splitter.
func TestScanTokens(t *testing.T) {
	cases := []struct {
		input string
		want  []string
	}{
		{"", []string{}},
		{" ", []string{}},
		{"\n", []string{}},
		{"a", []string{"a"}},
		{" a ", []string{"a"}},
		{"abc def", []string{"abc", "def"}},
		{" abc def ", []string{"abc", "def"}},
		{" a\tb\nc\rd\fe\vf\u0085g\u00a0\n", []string{"a", "b", "c", "d", "e", "f", "g"}},
		{"^0.8.13", []string{"^", "0.8.13"}},
		{"0.8.13;", []string{"0.8.13", ";"}},
		{"pragma solidity ^0.8.13;", []string{"pragma", "solidity", "^", "0.8.13", ";"}},
		{"contract HelloWorld {", []string{"contract", "HelloWorld", "{"}},
	}

	for n, c := range cases {
		buf := strings.NewReader(c.input)
		s := bufio.NewScanner(buf)
		s.Split(lexer.ScanTokens)

		got := make([]string, 0, len(c.want))
		for i := 0; i < len(c.want); i++ {
			if !s.Scan() {
				break
			}
			got = append(got, s.Text())
		}
		if s.Scan() {
			t.Errorf("#%d: scan ran too long, got %q", n, s.Text())
		}
		if diff := cmp.Diff(c.want, got); diff != "" {
			t.Errorf("#%d: %s", n, diff)
		}
		err := s.Err()
		if err != nil {
			t.Errorf("#%d: %v", n, err)
		}
	}
}

func TestLexer_Scan(t *testing.T) {
	cases := []struct {
		input string
		want  []lexer.Token
	}{
		{
			input: "pragma solidity ^0.8.13;",
			want: []lexer.Token{
				{lexer.Pragma, "pragma", lexer.Position{Start: 0, Size: 6, Line: 0}},
				{lexer.Unknown, "solidity", lexer.Position{Start: 6, Size: 8, Line: 0}},
				{lexer.Hat, "^", lexer.Position{Start: 14, Size: 1, Line: 0}},
				{lexer.Unknown, "0.8.13", lexer.Position{Start: 15, Size: 6, Line: 0}},
				{lexer.Semicolon, ";", lexer.Position{Start: 21, Size: 1, Line: 0}},
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
