package solparser_test

import (
	"strings"
	"testing"

	"github.com/uji/solparser"
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func TestParser_ParseSymbolAlias(t *testing.T) {
	tests := []struct {
		input string
		want  *ast.SymbolAlias
		err   *token.PosError
	}{
		{
			input: "symbol as alias1}",
			want: &ast.SymbolAlias{
				Symbol: ast.Identifier(tkn(token.Identifier, "symbol", pos(1, 1))),
				As:     posPtr(8, 1),
				Alias: &ast.Identifier{
					Type:     token.Identifier,
					Value:    "alias1",
					Position: pos(11, 1),
				},
			},
		},
		{
			input: "symbol}",
			want: &ast.SymbolAlias{
				Symbol: ast.Identifier(tkn(token.Identifier, "symbol", pos(1, 1))),
				As:     nil,
				Alias:  nil,
			},
		},
		{
			input: "pragma",
			err:   perr(pos(1, 1), "keyword is not available as identifier."),
		},
		{
			input: "symbol as pragma",
			err:   perr(pos(11, 1), "keyword is not available as identifier."),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParseSymbolAlias()
			assert(t, got, tt.want, err, tt.err)
		})
	}
}

func TestParser_ParseSymbolAliases(t *testing.T) {
	tests := []struct {
		input string
		want  *ast.SymbolAliases
		err   *token.PosError
	}{
		{
			input: "{symbol1}",
			want: &ast.SymbolAliases{
				LBrace: pos(1, 1),
				Aliases: []*ast.SymbolAlias{
					{
						Symbol: ast.Identifier(tkn(token.Identifier, "symbol1", pos(2, 1))),
					},
				},
				Commas: []*token.Pos{},
				RBrace: pos(9, 1),
			},
		},
		{
			input: "{symbol2 as alias1}",
			want: &ast.SymbolAliases{
				LBrace: pos(1, 1),
				Aliases: []*ast.SymbolAlias{
					{
						Symbol: ast.Identifier(tkn(token.Identifier, "symbol2", pos(2, 1))),
						As:     posPtr(10, 1),
						Alias: &ast.Identifier{
							Type:     token.Identifier,
							Value:    "alias1",
							Position: pos(13, 1),
						},
					},
				},
				Commas: []*token.Pos{},
				RBrace: pos(19, 1),
			},
		},
		{
			input: "{symbol1, symbol2 as alias1}",
			want: &ast.SymbolAliases{
				LBrace: pos(1, 1),
				Aliases: []*ast.SymbolAlias{
					{
						Symbol: ast.Identifier(tkn(token.Identifier, "symbol1", pos(2, 1))),
					},
					{
						Symbol: ast.Identifier(tkn(token.Identifier, "symbol2", pos(11, 1))),
						As:     posPtr(19, 1),
						Alias: &ast.Identifier{
							Type:     token.Identifier,
							Value:    "alias1",
							Position: pos(22, 1),
						},
					},
				},
				Commas: []*token.Pos{posPtr(9, 1)},
				RBrace: pos(28, 1),
			},
		},
		{
			input: "symbol1, symbol2 as alian1}",
			err:   perr(pos(1, 1), "not found LBrace."),
		},
		{
			input: "{symbol1, as alian1",
			err:   perr(pos(11, 1), "keyword is not available as identifier."),
		},
		{
			input: "{symbol1, symbol2 as alian1",
			err:   perr(pos(28, 1), "not found RBrace."),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParseSymbolAliases()
			assert(t, got, tt.want, err, tt.err)
		})
	}
}
