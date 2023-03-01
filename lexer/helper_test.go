package lexer

import (
	"errors"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/uji/solparser/token"
)

func pos(c, l int) token.Pos {
	return token.Pos{Column: c, Line: l}
}

func tkn(tp token.TokenType, text string, pos token.Pos) token.Token {
	return token.Token{
		Type:     tp,
		Value:    text,
		Position: pos,
	}
}

func perr(pos token.Pos, msg string) *token.PosError {
	return &token.PosError{
		Pos: pos,
		Msg: msg,
	}
}

type TestData[T any] []struct {
	input string
	want  T
	err   error
}

func (ts TestData[T]) Test(t *testing.T, target func(l *Lexer) (T, error)) {
	for _, tt := range ts {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			l := New(strings.NewReader(tt.input))
			got, err := target(l)

			var sErr *token.PosError
			if errors.As(err, &sErr) {
				if diff := cmp.Diff(tt.err, sErr); diff != "" {
					t.Errorf("%s", diff)
				}
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("%s", diff)
			}
		})
	}
}
