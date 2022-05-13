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
		{" ", []string{" "}},
		{"\n", []string{"\n"}},
		{"a", []string{"a"}},
		{" a ", []string{" ", "a", " "}},
		{"abc  def ", []string{"abc", "  ", "def", " "}},
		{"a\tb\nc\r\td\f", []string{"a", "\t", "b", "\n", "c", "\r\t", "d", "\f"}},
		{"e\vf\u0085g\u00a0\n", []string{"e", "\v", "f", "\u0085", "g", "\u00a0", "\n"}},
		{"^0.8.13", []string{"^", "0.8.13"}},
		{"0.8.13;", []string{"0.8.13", ";"}},
		{"pragma solidity ^0.8.13;", []string{"pragma", " ", "solidity", " ", "^", "0.8.13", ";"}},
		{"contract HelloWorld {", []string{"contract", " ", "HelloWorld", " ", "{"}},
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
