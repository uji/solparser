package token_test

import (
	"testing"

	"github.com/uji/solparser/token"
)

func TestPosition_String(t *testing.T) {
	cases := []struct {
		pos  token.Position
		want string
	}{
		{token.Position{Column: 3, Line: 10}, "10:3"},
		{token.Position{Column: 0, Line: 10}, "-"},
		{token.Position{Column: 1, Line: 0}, "-"},
	}

	for n, c := range cases {
		if got := c.pos.String(); got != c.want {
			t.Errorf("#%d: got: %s, want: %s", n, got, c.want)
		}
	}
}
