package lexer

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/uji/solparser/scanner"
	"github.com/uji/solparser/token"
)

func testLexerScan(t *testing.T) {
	tests := []struct {
		name       string
		offset     int
		lineOffset int
		input      string
		peeked     bool
		peekToken  token.Token
		peekErr    error
		wantResult bool
		wantErr    error
		wantToken  token.Token
	}{
		{
			name:       "normal",
			offset:     4,
			lineOffset: 5,
			input:      "pragma",
			wantResult: true,
			wantToken: token.Token{
				TokenType: token.Pragma,
				Text:      "pragma",
				Pos: token.Pos{
					Column: 5,
					Line:   6,
				},
			},
		},
		{
			name:       "when scan space",
			offset:     4,
			lineOffset: 5,
			input:      "  ^",
			wantResult: true,
			wantToken: token.Token{
				TokenType: token.BitXor,
				Text:      "^",
				Pos: token.Pos{
					Column: 7,
					Line:   6,
				},
			},
		},
		{
			name:       "when scan \\n",
			offset:     4,
			lineOffset: 5,
			input:      "  \n^",
			wantResult: true,
			wantToken: token.Token{
				TokenType: token.BitXor,
				Text:      "^",
				Pos: token.Pos{
					Column: 1,
					Line:   7,
				},
			},
		},
		{
			name:       "when scan is done",
			offset:     4,
			lineOffset: 5,
			input:      "",
			wantResult: false,
			wantToken:  token.Token{},
		},
		{
			name:       "when peeked",
			offset:     4,
			lineOffset: 5,
			input:      "",
			peeked:     true,
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
				Text:      "",
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
				offset:     tt.offset,
				lineOffset: tt.lineOffset,
				scanner:    s,
				peeked:     tt.peeked,
				peekToken:  tt.peekToken,
				peekErr:    tt.peekErr,
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
		name       string
		offset     int
		lineOffset int
		input      string
		result     bool
		token      token.Token
		err        error
	}{
		{
			name:       "normal",
			offset:     4,
			lineOffset: 5,
			input:      "pragma",
			result:     true,
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
			name:       "when scan is done",
			offset:     4,
			lineOffset: 5,
			input:      "",
			result:     false,
			token:      token.Token{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			s := scanner.New(strings.NewReader(tt.input))

			l := Lexer{
				offset:     tt.offset,
				lineOffset: tt.lineOffset,
				scanner:    s,
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
