package lexer_test

import (
	"testing"

	"github.com/uji/solparser/lexer"
)

func TestPosition_String(t *testing.T) {
	cases := []struct {
		pos  lexer.Position
		want string
	}{
		{lexer.Position{Column: 3, Line: 10}, "10:3"},
		{lexer.Position{Column: 0, Line: 10}, "-"},
		{lexer.Position{Column: 1, Line: 0}, "-"},
	}

	for n, c := range cases {
		if c.pos.String() != c.want {
			t.Errorf("#%d: got: %s, want: %s", n, c.pos.String(), c.want)
		}
	}
}
