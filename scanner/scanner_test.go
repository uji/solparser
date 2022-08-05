package scanner

import (
	"bufio"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/uji/solparser/lexer"
)

func TestSplit(t *testing.T) {
	tests := []struct {
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
		{`\'test\'`, []string{`\`, "'", "test", `\`, "'"}},
	}

	for n, c := range tests {
		buf := strings.NewReader(c.input)
		s := bufio.NewScanner(buf)
		s.Split(lexer.Scan)

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

func TestScannerScan(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		peeked   bool
		peekText string
		peekErr  error
		want     bool
		wantErr  error
		wantText string
	}{
		{
			name:     "normal",
			input:    "pragma solidity",
			want:     true,
			wantErr:  nil,
			wantText: "pragma",
		},
		{
			name:     "when scan is done",
			input:    "",
			want:     false,
			wantErr:  nil,
			wantText: "",
		},
		{
			name:     "when peeked",
			input:    "",
			peeked:   true,
			peekText: "peekedText",
			want:     true,
			wantErr:  nil,
			wantText: "peekedText",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.input))
			scanner.Split(Split)

			s := Scanner{
				scanner:  scanner,
				peeked:   tt.peeked,
				peekText: tt.peekText,
				peekErr:  tt.peekErr,
			}

			if rslt := s.Scan(); rslt != tt.want {
				t.Errorf("want: %t, got: %t", tt.want, rslt)
			}

			if err := s.Err(); err != tt.wantErr {
				t.Errorf("want: %s, got: %s", tt.wantErr, err)
			}

			if text := s.Text(); text != tt.wantText {
				t.Errorf("want: %s, got: %s", tt.wantText, text)
			}
		})
	}
}
