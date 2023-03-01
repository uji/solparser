package solparser_test

import (
	"strings"
	"testing"

	"github.com/uji/solparser"
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func TestParser_ParsePath(t *testing.T) {
	tests := []struct {
		input string
		want  ast.Path
		err   *token.PosError
	}{
		{
			input: `"test.sol"`,
			want:  ast.Path(tkn(token.NonEmptyStringLiteral, `"test.sol"`, pos(1, 1))),
		},
		{
			input: "pragma",
			err:   perr(pos(1, 1), "not found non-empty-string-literal."),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParsePath()
			assert(t, got, tt.want, err, tt.err)
		})
	}
}

func TestParser_ParseImportDirectivePathElement(t *testing.T) {
	tests := []struct {
		input string
		want  *ast.ImportDirectivePathElement
		err   *token.PosError
	}{
		{
			input: `"test.sol";`,
			want: &ast.ImportDirectivePathElement{
				Path: ast.Path(tkn(token.NonEmptyStringLiteral, `"test.sol"`, pos(1, 1))),
			},
		},
		{
			input: `"test.sol" as alias1 ;`,
			want: &ast.ImportDirectivePathElement{
				Path: ast.Path(tkn(token.NonEmptyStringLiteral, `"test.sol"`, pos(1, 1))),
				As:   posPtr(12, 1),
				Identifier: &ast.Identifier{
					Type:     token.Identifier,
					Value:    "alias1",
					Position: pos(15, 1),
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParseImportDirectivePathElement()
			assert(t, got, tt.want, err, tt.err)
		})
	}
}

func TestParser_ParseImportDirectiveSymbolAliasesElement(t *testing.T) {
	tests := []struct {
		input string
		want  *ast.ImportDirectiveSymbolAliasesElement
		err   *token.PosError
	}{
		{
			input: `{symbol1} from "test.sol"`,
			want: &ast.ImportDirectiveSymbolAliasesElement{
				SymbolAliases: &ast.SymbolAliases{
					LBrace: pos(1, 1),
					Aliases: []*ast.SymbolAlias{
						{Symbol: ast.Identifier{
							Type:     token.Identifier,
							Value:    "symbol1",
							Position: pos(2, 1),
						}},
					},
					Commas: []*token.Pos{},
					RBrace: pos(9, 1),
				},
				From: pos(11, 1),
				Path: ast.Path(tkn(token.NonEmptyStringLiteral, `"test.sol"`, pos(16, 1))),
			},
		},
		{
			input: `{symbol1} "test.sol"`,
			err:   perr(pos(11, 1), "not found from keyword."),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParseImportDirectiveSymbolAliasesElement()
			assert(t, got, tt.want, err, tt.err)
		})
	}
}

func TestParser_ParseImportDirectiveMulElement(t *testing.T) {
	tests := []struct {
		input string
		want  *ast.ImportDirectiveMulElement
		err   *token.PosError
	}{
		{
			input: `* as alias1 from "test.sol"`,
			want: &ast.ImportDirectiveMulElement{
				Mul: pos(1, 1),
				As:  pos(3, 1),
				Identifier: ast.Identifier{
					Type:     token.Identifier,
					Value:    "alias1",
					Position: pos(6, 1),
				},
				From: pos(13, 1),
				Path: ast.Path(tkn(token.NonEmptyStringLiteral, `"test.sol"`, pos(18, 1))),
			},
		},
		{
			input: `{symbol1} "test.sol"`,
			err:   perr(pos(1, 1), "not found mul keyword."),
		},
		{
			input: `* alias1 from "test.sol"`,
			err:   perr(pos(3, 1), "not found as keyword."),
		},
		{
			input: `* as alias1 "test.sol"`,
			err:   perr(pos(13, 1), "not found from keyword."),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParseImportDirectiveMulElement()
			assert(t, got, tt.want, err, tt.err)
		})
	}
}

func TestParser_ParseImportDirective(t *testing.T) {
	tests := []struct {
		input string
		want  *ast.ImportDirective
		err   *token.PosError
	}{
		{
			input: `import "test.sol";`,
			want: &ast.ImportDirective{
				Import: pos(1, 1),
				Element: &ast.ImportDirectivePathElement{
					Path: ast.Path(tkn(token.NonEmptyStringLiteral, `"test.sol"`, pos(8, 1))),
				},
				Semicolon: pos(18, 1),
			},
		},
		{
			input: `import {symbol1} from "test.sol";`,
			want: &ast.ImportDirective{
				Import: pos(1, 1),
				Element: &ast.ImportDirectiveSymbolAliasesElement{
					SymbolAliases: &ast.SymbolAliases{
						LBrace: pos(8, 1),
						Aliases: []*ast.SymbolAlias{
							{Symbol: ast.Identifier{
								Type:     token.Identifier,
								Value:    "symbol1",
								Position: pos(9, 1),
							}},
						},
						Commas: []*token.Pos{},
						RBrace: pos(16, 1),
					},
					From: pos(18, 1),
					Path: ast.Path(tkn(token.NonEmptyStringLiteral, `"test.sol"`, pos(23, 1))),
				},
				Semicolon: pos(33, 1),
			},
		},
		{
			input: `import * as alias1 from "test.sol";`,
			want: &ast.ImportDirective{
				Import: pos(1, 1),
				Element: &ast.ImportDirectiveMulElement{
					Mul: pos(8, 1),
					As:  pos(10, 1),
					Identifier: ast.Identifier{
						Type:     token.Identifier,
						Value:    "alias1",
						Position: pos(13, 1),
					},
					From: pos(20, 1),
					Path: ast.Path(tkn(token.NonEmptyStringLiteral, `"test.sol"`, pos(25, 1))),
				},
				Semicolon: pos(35, 1),
			},
		},
		{
			input: "symbol as pragma",
			err:   perr(pos(1, 1), "not found import keyword."),
		},
		{
			input: `import alias1 from "test.sol";`,
			err:   perr(pos(8, 1), "not found import-directive element."),
		},
		{
			input: `import * as alias1 from "test.sol"`,
			err:   perr(pos(35, 1), "not found semicolon."),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParseImportDirective()
			assert(t, got, tt.want, err, tt.err)
		})
	}
}
