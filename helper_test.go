package solparser_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/uji/solparser"
	"github.com/uji/solparser/token"
)

func pos(c, l int) token.Pos {
	return token.Pos{Column: c, Line: l}
}

func posPtr(c, l int) *token.Pos {
	return &token.Pos{Column: c, Line: l}
}

func tkn(tp token.TokenType, text string, pos token.Pos) token.Token {
	return token.Token{
		Type:     tp,
		Value:    text,
		Position: pos,
	}
}

func tknPtr(tp token.TokenType, text string, pos token.Pos) *token.Token {
	return &token.Token{
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

func assert(t *testing.T, got, want interface{}, gotErr, wantErr error) {
	t.Helper()
	var sErr *token.PosError
	if errors.As(gotErr, &sErr) {
		if diff := cmp.Diff(wantErr, sErr); diff != "" {
			t.Errorf("%s", diff)
		}
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("%s", diff)
	}
}

type TestData[T any] []struct {
	input string
	want  T
	err   error
}

func (ts TestData[T]) Test(t *testing.T, target func(p *solparser.Parser) (T, error)) {
	for _, tt := range ts {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)
			got, err := target(p)
			assert(t, got, tt.want, err, tt.err)
		})
	}
}
