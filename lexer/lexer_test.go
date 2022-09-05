package lexer

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	scanner "github.com/uji/solparser/scanner2"
	"github.com/uji/solparser/token"
)

func TestLexerScan(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		peeked     bool
		peekToken  token.Token
		peekErr    error
		wantResult bool
		wantErr    error
		wantToken  token.Token
	}{
		// TODO: fix test after implementing new scanner.
		// {
		// 	name:       "normal",
		// 	input:      "pragma",
		// 	wantResult: true,
		// 	wantToken: token.Token{
		// 		TokenType: token.Pragma,
		// 		Text:      "pragma",
		// 		Pos: token.Pos{
		// 			Column: 5,
		// 			Line:   6,
		// 		},
		// 	},
		// },
		// {
		// 	name:       "when scan is done",
		// 	input:      "",
		// 	wantResult: false,
		// 	wantToken:  token.Token{},
		// },
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
			wantResult: true,
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
			if rslt := l.Scan(); rslt != tt.wantResult {
				t.Errorf("result is wrong, want: %t, got: %t", tt.wantResult, rslt)
			}
			if err := l.Error(); err != tt.wantErr {
				t.Errorf("error is wrong, want: %s, got: %s", tt.wantErr, err)
			}
			if diff := cmp.Diff(tt.wantToken, l.Token()); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func TestLexerPeek(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result bool
		token  token.Token
		err    error
	}{
		{
			name:   "normal",
			input:  "pragma",
			result: true,
			token: token.Token{
				TokenType: token.Pragma,
				Text:      "pragma",
				Pos: token.Pos{
					Column: 5,
					Line:   6,
				},
			},
		},
		{
			name:   "when scan is done",
			input:  "",
			result: false,
			token:  token.Token{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			s := scanner.New(strings.NewReader(tt.input))

			l := Lexer{
				scanner: s,
			}
			if rslt := l.Peek(); tt.result != rslt {
				t.Errorf("result is wrong, want: %t, got: %t", tt.result, rslt)
			}
			if err := l.PeekError(); err != tt.err {
				t.Errorf("want: %s, got: %s", tt.err, err)
			}
			if diff := cmp.Diff(tt.token, l.PeekToken()); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
