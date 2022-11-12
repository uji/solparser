package lexer

import (
	"io"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/uji/solparser/scanner"
	"github.com/uji/solparser/token"
)

func TestLexerScan(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		peeked    bool
		peekToken token.Token
		peekErr   error
		wantErr   error
		wantToken token.Token
	}{
		{
			name:  "normal",
			input: "pragma",
			wantToken: token.Token{
				TokenType: token.Pragma,
				Text:      "pragma",
				Pos: token.Pos{
					Column: 1,
					Line:   1,
				},
			},
		},
		{
			name:      "when scan is done",
			input:     "",
			wantToken: token.Token{},
			wantErr:   io.EOF,
		},
		{
			name:   "when peeked",
			input:  "",
			peeked: true,
			peekToken: token.Token{
				TokenType: token.BitXor,
				Text:      "^",
				Pos: token.Pos{
					Column: 4,
					Line:   6,
				},
			},
			wantToken: token.Token{
				TokenType: token.BitXor,
				Text:      "^",
				Pos: token.Pos{
					Column: 4,
					Line:   6,
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			s := scanner.New(strings.NewReader(tt.input))

			l := Lexer{
				scanner:   s,
				peeked:    tt.peeked,
				peekToken: tt.peekToken,
				peekErr:   tt.peekErr,
			}
			tkn, err := l.Scan()
			if err != tt.wantErr {
				t.Errorf("error is wrong, want: %s, got: %s", tt.wantErr, err)
			}
			if diff := cmp.Diff(tt.wantToken, tkn); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func TestLexerPeek(t *testing.T) {
	tests := []struct {
		name  string
		input string
		token token.Token
		err   error
	}{
		{
			name:  "normal",
			input: "pragma",
			token: token.Token{
				TokenType: token.Pragma,
				Text:      "pragma",
				Pos: token.Pos{
					Column: 1,
					Line:   1,
				},
			},
		},
		{
			name:  "when scan is done",
			input: "",
			token: token.Token{},
			err:   io.EOF,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			s := scanner.New(strings.NewReader(tt.input))

			l := Lexer{
				scanner: s,
			}
			tkn, err := l.Peek()
			if err != tt.err {
				t.Errorf("want: %s, got: %s", tt.err, err)
			}
			if diff := cmp.Diff(tt.token, tkn); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
