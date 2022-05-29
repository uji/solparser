package solparser_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/uji/solparser"
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/lexer"
)

func TestParserParsePragmaDirective(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  *ast.PragmaDirective
		err   *solparser.Error
	}{
		{
			name:  "normal case",
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
		{
			name:  "there is no pragma keyword",
			input: "solidity ^0.8.13;",
			want:  nil,
			err: &solparser.Error{
				Pos: lexer.Position{
					Column: 1,
					Line:   1,
				},
				Msg: "not found pragma",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParsePragmaDirective()

			var sErr *solparser.Error
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
