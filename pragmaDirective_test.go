package solparser_test

import (
	"strings"
	"testing"

	"github.com/uji/solparser"
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func TestParser_ParsePragmaDirective(t *testing.T) {
	tests := []struct {
		input string
		want  *ast.PragmaDirective
		err   *token.PosError
	}{
		{
			input: "pragma solidity ^0.8.13;",
			want: &ast.PragmaDirective{
				Pragma: pos(1, 1),
				PragmaTokens: []*token.Token{
					{TokenType: token.Identifier, Text: "solidity", Pos: pos(8, 1)},
					{TokenType: token.BitXor, Text: "^", Pos: pos(17, 1)},
					{TokenType: token.Identifier, Text: "0.8.13", Pos: pos(18, 1)},
				},
				Semicolon: pos(24, 1),
			},
		},
		{input: "solidity ^0.8.13;", err: perr(pos(1, 1), "not found pragma.")},
		{input: "pragma ;", err: perr(pos(8, 1), "not found pragma tokens.")},
		{input: "pragma solidity ^0.8.13", err: perr(pos(24, 1), "not found Semicolon.")},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(*testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParsePragmaDirective()
			assert(t, got, tt.want, err, tt.err)
		})
	}
}
