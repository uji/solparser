package solparser_test

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/uji/solparser"
	"github.com/uji/solparser/ast"
)

func TestParserParsePragmaDirective(t *testing.T) {
	cases := []struct {
		input string
		want  *ast.PragmaDirective
		err   error
	}{
		{
			input: "pragma solidity ^0.8.13;",
			want: &ast.PragmaDirective{
				PragmaName: "solidity",
				PragmaValue: ast.PragmaValue{
					Version:    "0.8.13",
					Expression: "^",
				},
			},
			err: nil,
		},
	}

	for n, c := range cases {
		r := strings.NewReader(c.input)
		p := solparser.New(r)
		got, err := p.ParsePragmaDirective()
		if err != c.err {
			t.Errorf("#%d unexpected err want: %s, got: %s", n, c.err, err)
		}
		if err != nil {
			break
		}
		if diff := cmp.Diff(c.want, got); diff != "" {
			t.Errorf("#%d want: %s, got: %s", n, c.want, got)
		}
	}
}
