package lexer

import (
	"errors"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/uji/solparser/scanner"
	"github.com/uji/solparser/token"
)

func TestLexer_Scan(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantToken token.Token
	}{
		{
			name:  "normal",
			input: "pragma",
			wantToken: token.Token{
				Type:  token.Pragma,
				Value: "pragma",
				Position: token.Pos{
					Column: 1,
					Line:   1,
				},
			},
		},
		{
			name:  "There is a space at the beginning",
			input: " pragma",
			wantToken: token.Token{
				Type:  token.Pragma,
				Value: "pragma",
				Position: token.Pos{
					Column: 2,
					Line:   1,
				},
			},
		},
		{
			name:  "There is a StringLiteral",
			input: `"pragma"`,
			wantToken: token.Token{
				Type:  token.StringLiteral,
				Value: `"pragma"`,
				Position: token.Pos{
					Column: 1,
					Line:   1,
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			s := scanner.New(strings.NewReader(tt.input))
			l := Lexer{
				scanner: s,
			}

			tkn, err := l.Scan()
			if err != nil {
				t.Errorf("error is not nil, got: %s", err)
			}
			if diff := cmp.Diff(tt.wantToken, tkn); diff != "" {
				t.Errorf(diff)
			}
		})
	}

	t.Run("when scan is done", func(t *testing.T) {
		s := scanner.New(strings.NewReader(""))
		l := Lexer{
			scanner: s,
		}

		got, err := l.Scan()
		if err != nil {
			t.Errorf("got error: %s", err)
		}
		want := token.Token{
			Type:  token.EOS,
			Value: token.EOSString,
			Position: token.Pos{
				Column: 1,
				Line:   1,
			},
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("when peeked", func(t *testing.T) {
		peekToken := token.Token{
			Type:  token.BitXor,
			Value: "^",
			Position: token.Pos{
				Column: 4,
				Line:   6,
			},
		}
		s := scanner.New(strings.NewReader(""))
		l := Lexer{
			scanner:   s,
			peeked:    true,
			peekToken: peekToken,
			peekErr:   nil,
		}

		tkn, err := l.Scan()
		if err != nil {
			t.Errorf("error is not nil, got: %s", err)
		}
		if diff := cmp.Diff(peekToken, tkn); diff != "" {
			t.Errorf(diff)
		}
	})
}

func TestLexer_Peek(t *testing.T) {
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
				Type:  token.Pragma,
				Value: "pragma",
				Position: token.Pos{
					Column: 1,
					Line:   1,
				},
			},
		},
		{
			name:  "when scan is done",
			input: "",
			token: token.Token{
				Type:  token.EOS,
				Value: token.EOSString,
				Position: token.Pos{
					Column: 1,
					Line:   1,
				},
			},
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

func TestLexer_ScanStringLiteral(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantText string
	}{
		{
			name:     "double-quoted-printable",
			input:    `"Hello world!!";`,
			wantText: `"Hello world!!"`,
		},
		{
			name:     "single-quoted-printable",
			input:    `\'Hello world!!\';`,
			wantText: `\'Hello world!!\'`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			s := scanner.New(strings.NewReader(tt.input))
			l := Lexer{
				scanner: s,
			}

			tkn, err := l.ScanStringLiteral()
			if err != nil {
				t.Errorf("error is not nil, got: %s", err)
			}
			if diff := cmp.Diff(tt.wantText, tkn.Value); diff != "" {
				t.Errorf(diff)
			}
		})
	}

	t.Run(`not found "`, func(t *testing.T) {
		s := scanner.New(strings.NewReader("a"))
		l := Lexer{
			scanner: s,
		}
		exptErr := &token.PosError{
			Pos: token.Pos{
				Column: 1,
				Line:   1,
			},
			Msg: `not found " or \'`,
		}

		_, err := l.ScanStringLiteral()
		var pErr *token.PosError
		if !errors.As(err, &pErr) {
			t.Errorf("error is unexpected, got: %s", err)
		}
		if diff := cmp.Diff(pErr, exptErr); diff != "" {
			t.Errorf(diff)
		}
	})
}
