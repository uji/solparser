package solparser_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/uji/solparser"
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func TestParserParsePragmaDirective(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  *ast.PragmaDirective
		err   *token.PosError
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
			err: &token.PosError{
				Pos: token.Pos{
					Column: 1,
					Line:   1,
				},
				Msg: "not found pragma",
			},
		},
		{
			name:  "there is no pragma name",
			input: "pragma ^0.8.13;",
			want:  nil,
			err: &token.PosError{
				Pos: token.Pos{
					Column: 8,
					Line:   1,
				},
				Msg: "not found pragma name",
			},
		},
		{
			name:  "there is no BitXor expression",
			input: "pragma solidity 0.8.13;",
			want:  nil,
			err: &token.PosError{
				Pos: token.Pos{
					Column: 17,
					Line:   1,
				},
				Msg: "not found BitXor expression",
			},
		},
		{
			name:  "there is no Semicolon",
			input: "pragma solidity ^0.8.13",
			want:  nil,
			err: &token.PosError{
				Pos: token.Pos{
					Column: 18,
					Line:   1,
				},
				Msg: "not found Semicolon",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParsePragmaDirective()

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
