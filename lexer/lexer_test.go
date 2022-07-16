package lexer

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/uji/solparser/token"
)

type scanResult struct {
	text string
	err  error
}

type mockScanner struct {
	calledCount int
	results     []scanResult
}

func (s *mockScanner) Scan() bool {
	if s.calledCount == len(s.results) {
		return false
	}
	s.calledCount++
	return true
}
func (s *mockScanner) Text() string {
	return s.results[s.calledCount-1].text
}
func (s *mockScanner) Err() error {
	return s.results[s.calledCount-1].err
}

func TestLexerScan(t *testing.T) {
	tests := []struct {
		name        string
		offset      int
		lineOffset  int
		scanResults []scanResult
		result      bool
		token       token.Token
		err         error
	}{
		{
			name:       "normal",
			offset:     4,
			lineOffset: 5,
			scanResults: []scanResult{
				{
					text: "pragma",
					err:  nil,
				},
			},
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
			name:       "when scan space",
			offset:     4,
			lineOffset: 5,
			scanResults: []scanResult{
				{
					text: "  ",
					err:  nil,
				},
				{
					text: "^",
					err:  nil,
				},
			},
			result: true,
			token: token.Token{
				TokenType: token.Hat,
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
			scanResults: []scanResult{
				{
					text: "  ",
					err:  nil,
				},
				{
					text: "\n",
					err:  nil,
				},
				{
					text: "^",
					err:  nil,
				},
			},
			result: true,
			token: token.Token{
				TokenType: token.Hat,
				Text:      "^",
				Pos: token.Pos{
					Column: 1,
					Line:   7,
				},
			},
		},
		{
			name:        "when scan is done",
			offset:      4,
			lineOffset:  5,
			scanResults: []scanResult{},
			result:      false,
			token:       token.Token{},
		},
		{
			name:        "when peeked",
			offset:      4,
			lineOffset:  5,
			scanResults: []scanResult{},
			result:      false,
			token:       token.Token{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &mockScanner{
				results: tt.scanResults,
			}
			l := Lexer{
				offset:     tt.offset,
				lineOffset: tt.lineOffset,
				scanner:    s,
			}
			if rslt := l.Scan(); rslt != tt.result {
				t.Errorf("result is wrong, want: %t, got: %t", tt.result, rslt)
			}
			if err := l.Error(); err != tt.err {
				t.Errorf("error is wrong, want: %s, got: %s", tt.err, err)
			}
			if diff := cmp.Diff(tt.token, l.Token()); diff != "" {
				t.Errorf(diff)
			}
		})
	}

	t.Run("when peeked", func(t *testing.T) {
		s := &mockScanner{
			results: nil,
		}
		peekErr := errors.New("peekErr")
		peekToken := token.Token{
			TokenType: token.Hat,
			Text:      "^",
			Pos: token.Pos{
				Column: 7,
				Line:   6,
			},
		}
		l := Lexer{
			offset:     4,
			lineOffset: 6,
			scanner:    s,
			peeked:     true,
			peekToken:  peekToken,
			peekErr:    peekErr,
		}
		if rslt := l.Scan(); !rslt {
			t.Errorf("result is wrong, want: true, got: true")
		}
		if err := l.Error(); err != peekErr {
			t.Errorf("error is wrong, want: %s, got: %s", peekErr, err)
		}
		if diff := cmp.Diff(peekToken, l.Token()); diff != "" {
			t.Errorf(diff)
		}
	})
}

func TestLexerPeek(t *testing.T) {
	tests := []struct {
		name        string
		offset      int
		lineOffset  int
		scanResults []scanResult
		result      bool
		token       token.Token
		err         error
	}{
		{
			name:       "normal",
			offset:     4,
			lineOffset: 5,
			scanResults: []scanResult{
				{
					text: "pragma",
					err:  nil,
				},
			},
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
			name:        "when scan is done",
			offset:      4,
			lineOffset:  5,
			scanResults: []scanResult{},
			result:      false,
			token:       token.Token{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &mockScanner{
				results: tt.scanResults,
			}
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
