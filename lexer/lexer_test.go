package lexer

import (
	"bufio"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
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
		{
			" abc\tdef\nghi\rjkl\fmno\vpqr\u0085stu\u00a0\n",
			[]string{
				"abc",
				"def",
				"ghi",
				"jkl",
				"mno",
				"pqr",
				"stu",
			},
		},
		{"^0.8.13", []string{"^", "0.8.13"}},
		{"0.8.13;", []string{"0.8.13", ";"}},
		{"pragma solidity ^0.8.13;", []string{"pragma", "solidity", "^", "0.8.13", ";"}},
	}

	for n, c := range cases {
		buf := strings.NewReader(c.input)
		s := bufio.NewScanner(buf)
		s.Split(ScanTokens)

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
