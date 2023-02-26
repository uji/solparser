package solparser_test

import (
	"errors"
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
